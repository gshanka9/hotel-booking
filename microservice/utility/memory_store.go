package utility

import (
	"hotel_booking/microservice/model"
	"sync"
)

type MemoryStore struct {
	Bookings  map[string]*model.Booking
	Rooms     map[string]*model.Room
	Customers map[string]*model.Customer
	Mu        sync.Mutex
}

var Store *MemoryStore

func init() {
	Store = &MemoryStore{
		Bookings:  make(map[string]*model.Booking),
		Rooms:     make(map[string]*model.Room),
		Customers: make(map[string]*model.Customer),
	}
}
