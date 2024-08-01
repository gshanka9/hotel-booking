package model

type Room struct {
	ID      string  `json:"id"`
	HotelID string  `json:"hotel_id"`
	Number  string  `json:"number"`
	Type    string  `json:"type"`
	Price   float64 `json:"price"`
}
