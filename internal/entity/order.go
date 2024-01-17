package entity

import (
	"github.com/metafates/application-design-test/internal/pkg/uuid"
)

type OrderID string

func NewOrderID() OrderID {
	return OrderID(uuid.NewPseudoUUID())
}

type Order struct {
	ID      OrderID
	User    User
	Booking Booking
}
