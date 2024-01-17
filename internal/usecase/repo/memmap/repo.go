package memmap

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/metafates/application-design-test/internal/entity"
	"github.com/metafates/application-design-test/internal/usecase"
)

var _ usecase.HotelBookingRepo = (*Repo)(nil)

type Repo struct {
	byID   map[entity.OrderID]entity.Order
	logger *slog.Logger
}

func (r *Repo) StoreOrder(_ context.Context, order entity.Order) error {
	if _, exists := r.byID[order.ID]; exists {
		return fmt.Errorf("order %s exists", order.ID)
	}

	for _, actual := range r.byID {
		isSameRoom := actual.Booking.Room.Type == order.Booking.Room.Type
		isTimeOverlaps := actual.Booking.TimeRange.Overlaps(order.Booking.TimeRange)

		if isSameRoom && isTimeOverlaps {
			return usecase.BookingUnavailableError{Booking: order.Booking}
		}
	}

	r.byID[order.ID] = order
	return nil
}

func (r *Repo) RetrieveOrders(context.Context) ([]entity.Order, error) {
	orders := make([]entity.Order, 0, len(r.byID))
	for _, order := range r.byID {
		orders = append(orders, order)
	}

	return orders, nil
}
