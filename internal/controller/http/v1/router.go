package v1

import (
	"github.com/gorilla/mux"
	"github.com/metafates/application-design-test/internal/usecase"
	"net/http"
)

func RegisterRoutes(router *mux.Router, hotelBookingService usecase.HotelBookingService) *mux.Router {
	router.
		Path("/health").
		Methods(http.MethodGet).
		HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	registerHotelBookingRoutes(
		router.PathPrefix("/hotels").Subrouter(),
		hotelBookingService,
	)

	return router
}
