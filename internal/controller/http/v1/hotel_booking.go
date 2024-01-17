package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metafates/application-design-test/internal/entity"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/metafates/application-design-test/internal/usecase"
)

const jsonMimeType = "application/json"

func registerHotelBookingRoutes(router *mux.Router, service usecase.HotelBookingService) {
	routes := &hotelBookingRoutes{
		service: service,
	}

	router.Path("/orders").
		Methods(http.MethodPost).
		Headers("Content-Type", jsonMimeType).
		HandlerFunc(routes.postOrder)

	router.Path("/bookings").
		Methods(http.MethodGet).
		Headers("Accept", jsonMimeType).
		HandlerFunc(routes.getBookings)
}

type hotelBookingRoutes struct {
	service usecase.HotelBookingService
}

func (h *hotelBookingRoutes) postOrder(w http.ResponseWriter, r *http.Request) {
	var requestBody _Order

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, fmt.Sprintf("json: %s", err.Error()), http.StatusBadRequest)
		return
	}

	order, err := requestBody.Map()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.MakeOrder(r.Context(), order); err != nil {
		var bookingTimeUnavailableError usecase.BookingUnavailableError

		switch {
		case errors.As(err, &bookingTimeUnavailableError):
			http.Error(w, err.Error(), http.StatusConflict)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *hotelBookingRoutes) getBookings(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.Orders(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := r.URL.Query()

	if email := strings.TrimSpace(query.Get("email")); email != "" {
		filtered := make([]entity.Order, 0, cap(orders))

		for _, order := range orders {
			if strings.EqualFold(order.User.Email, email) {
				filtered = append(filtered, order)
			}
		}

		orders = filtered
	}

	responseBody := make([]_Booking, 0, cap(orders))
	for _, order := range orders {
		responseBody = append(responseBody, _NewBooking(order.Booking))
	}

	w.Header().Set("Content-Type", jsonMimeType)

	if err = json.NewEncoder(w).Encode(responseBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
