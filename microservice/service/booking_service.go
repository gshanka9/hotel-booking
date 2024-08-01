package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"hotel_booking/microservice/model"
	"hotel_booking/microservice/utility"
)

func CreateBooking(logger *logrus.Logger, booking *model.Booking) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Bookings[booking.ID]; exists {
		return errors.New("booking already exists")
	}

	utility.Store.Bookings[booking.ID] = booking
	logger.Infof("Booking created in memory: %+v", booking)
	return nil
}

func GetBooking(logger *logrus.Logger, id string) (*model.Booking, error) {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	booking, exists := utility.Store.Bookings[id]
	if !exists {
		return nil, errors.New("booking not found")
	}

	return booking, nil
}

func UpdateBooking(logger *logrus.Logger, id string, booking *model.Booking) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Bookings[id]; !exists {
		return errors.New("booking not found")
	}

	utility.Store.Bookings[id] = booking
	logger.Infof("Booking updated in memory: %+v", booking)
	return nil
}

func DeleteBooking(logger *logrus.Logger, id string) error {
	utility.Store.Mu.Lock()
	defer utility.Store.Mu.Unlock()

	if _, exists := utility.Store.Bookings[id]; !exists {
		return errors.New("booking not found")
	}

	delete(utility.Store.Bookings, id)
	logger.Infof("Booking deleted from memory: %s", id)
	return nil
}
