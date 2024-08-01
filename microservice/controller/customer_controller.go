package controller

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"hotel_booking/microservice/model"
	"hotel_booking/microservice/service"
	"hotel_booking/microservice/utility"
	"net/http"
)

func CreateCustomerHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var customer model.Customer
		if errr := json.NewDecoder(r.Body).Decode(&customer); errr != nil {
			utility.LogWithStack(logger, errr)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if err := service.CreateCustomer(logger, &customer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer created successfully"))
	}
}

func GetCustomerHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		customer, err := service.GetCustomer(logger, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

func UpdateCustomerHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var customer model.Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			utility.LogWithStack(logger, err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		id := r.URL.Query().Get("id")
		if err := service.UpdateCustomer(logger, id, &customer); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer updated successfully"))
	}
}

func DeleteCustomerHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if err := service.DeleteCustomer(logger, id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer deleted successfully"))
	}
}
