package v1

import (
	"time"

	"github.com/metafates/application-design-test/internal/entity"
)

type _Order struct {
	User struct {
		Email string `json:"email"`
	} `json:"user"`

	Booking _Booking `json:"booking"`
}

func (o _Order) Map() (entity.Order, error) {
	booking, err := o.Booking.Map()
	if err != nil {
		return entity.Order{}, err
	}

	return entity.Order{
		ID: entity.NewOrderID(),
		User: entity.User{
			Email: o.User.Email,
		},
		Booking: booking,
	}, nil
}

type _Booking struct {
	TimeRange struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	} `json:"time_range"`
	Room struct {
		Type string `json:"type"`
	} `json:"room"`
}

func _NewBooking(booking entity.Booking) _Booking {
	var b _Booking

	b.TimeRange.From = booking.TimeRange.From
	b.TimeRange.To = booking.TimeRange.To

	b.Room.Type = booking.Room.Type.String()

	return b
}

func (b *_Booking) Map() (entity.Booking, error) {
	roomType, err := entity.RoomTypeString(b.Room.Type)
	if err != nil {
		return entity.Booking{}, err
	}

	return entity.Booking{
		TimeRange: entity.TimeRange{
			From: b.TimeRange.From,
			To:   b.TimeRange.To,
		},
		Room: entity.Room{
			Type: roomType,
		},
	}, nil
}
