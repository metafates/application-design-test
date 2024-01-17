package entity

//go:generate go run github.com/dmarkham/enumer -type=RoomType -json -trimprefix RoomType
type RoomType int

const (
	RoomTypeEconomy RoomType = iota + 1
	RoomTypeStandard
	RoomTypeLuxury
)

type Room struct {
	Type RoomType
}
