package model

type Booking struct {
	ID      string `json:"id"`
	HotelID string `json:"hotel_id"`
	UserID  string `json:"user_id"`
	RoomID  string `json:"room_id"`
	Date    string `json:"date"`
}
