package controller

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"hotel_booking/microservice/model"
	"hotel_booking/microservice/service"
	"hotel_booking/microservice/utility"
	"net/http"
)

func CreateBookingHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var booking model.Booking
		if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
			utility.LogWithStack(logger, err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if err := service.CreateBooking(logger, &booking); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Booking created successfully"))
	}
}

func GetBookingHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		booking, err := service.GetBooking(logger, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(booking)
	}
}

func UpdateBookingHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var booking model.Booking
		if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
			utility.LogWithStack(logger, err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		id := r.URL.Query().Get("id")
		if err := service.UpdateBooking(logger, id, &booking); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Booking updated successfully"))
	}
}

func DeleteBookingHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if err := service.DeleteBooking(logger, id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Booking deleted successfully"))
	}
}
