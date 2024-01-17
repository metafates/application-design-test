package entity

import (
	"fmt"
	"time"
)

type TimeRange struct {
	From time.Time
	To   time.Time
}

func (t TimeRange) String() string {
	return fmt.Sprintf(
		"%s - %s",
		t.From,
		t.To,
	)
}

func (t TimeRange) Overlaps(other TimeRange) bool {
	if t.From.Equal(other.From) {
		return true
	}

	if t.To.Equal(other.To) {
		return true
	}

	if t.From.Before(other.From) && t.To.After(other.From) {
		return true
	}

	if t.From.Before(other.To) && t.To.After(other.To) {
		return true
	}

	return false
}
