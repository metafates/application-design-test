package usecase

import (
	"context"

	"github.com/metafates/application-design-test/internal/entity"
)

type (
	HotelBookingService interface {
		MakeOrder(ctx context.Context, order entity.Order) error
		Orders(ctx context.Context) ([]entity.Order, error)
	}

	HotelBookingRepo interface {
		StoreOrder(ctx context.Context, order entity.Order) error
		RetrieveOrders(ctx context.Context) ([]entity.Order, error)
	}
)
