package usecase

import (
	"fmt"
	"github.com/metafates/application-design-test/internal/entity"
)

type BookingUnavailableError struct {
	Booking entity.Booking
}

func (b BookingUnavailableError) Error() string {
	return fmt.Sprintf(
		"booking with %q room at %s is unavailable",
		b.Booking.Room.Type,
		b.Booking.TimeRange,
	)
}
