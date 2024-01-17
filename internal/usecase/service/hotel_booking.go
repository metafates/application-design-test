package service

import (
	"context"
	"log/slog"

	"github.com/metafates/application-design-test/internal/entity"
	"github.com/metafates/application-design-test/internal/usecase"
)

var _ usecase.HotelBookingService = (*HotelBookingServiceUseCase)(nil)

func NewHotelBooking(logger *slog.Logger, repo usecase.HotelBookingRepo) *HotelBookingServiceUseCase {
	return &HotelBookingServiceUseCase{
		logger: logger,
		repo:   repo,
	}
}

type HotelBookingServiceUseCase struct {
	logger *slog.Logger
	repo   usecase.HotelBookingRepo
}

func (h *HotelBookingServiceUseCase) MakeOrder(ctx context.Context, order entity.Order) error {
	if err := h.repo.StoreOrder(ctx, order); err != nil {
		h.logger.Error("failed to create order", "order", order.ID, "err", err)
		return err
	}

	h.logger.Info("new order", "order", order.ID)
	return nil
}

func (h *HotelBookingServiceUseCase) Orders(ctx context.Context) ([]entity.Order, error) {
	orders, err := h.repo.RetrieveOrders(ctx)
	if err != nil {
		h.logger.Error("failed to get orders", "err", err)
		return nil, err
	}

	h.logger.Debug("retrieved orders", "count", len(orders))
	return orders, nil
}
