package controller

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"hotel_booking/microservice/model"
	"hotel_booking/microservice/service"
	"hotel_booking/microservice/utility"
	"net/http"
)

func CreateRoomHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var room model.Room
		if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
			utility.LogWithStack(logger, err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if err := service.CreateRoom(logger, &room); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Room created successfully"))
	}
}

func GetRoomHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		room, err := service.GetRoom(logger, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(room)
	}
}

func UpdateRoomHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var room model.Room
		if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
			utility.LogWithStack(logger, err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		id := r.URL.Query().Get("id")
		if err := service.UpdateRoom(logger, id, &room); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Room updated successfully"))
	}
}

func DeleteRoomHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if err := service.DeleteRoom(logger, id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Room deleted successfully"))
	}
}
