package entity

type Booking struct {
	TimeRange TimeRange `json:"time_range"`
	Room      Room      `json:"room"`
}
