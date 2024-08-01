package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"hotel_booking/microservice/controller"
)

func SetupRoutes(logger *logrus.Logger) *mux.Router {
	router := mux.NewRouter()

	// Booking routes
	router.HandleFunc("/book", controller.CreateBookingHandler(logger)).Methods("POST")
	router.HandleFunc("/book", controller.GetBookingHandler(logger)).Methods("GET")
	router.HandleFunc("/book", controller.UpdateBookingHandler(logger)).Methods("PUT")
	router.HandleFunc("/book", controller.DeleteBookingHandler(logger)).Methods("DELETE")

	// Room routes
	router.HandleFunc("/room", controller.CreateRoomHandler(logger)).Methods("POST")
	router.HandleFunc("/room", controller.GetRoomHandler(logger)).Methods("GET")
	router.HandleFunc("/room", controller.UpdateRoomHandler(logger)).Methods("PUT")
	router.HandleFunc("/room", controller.DeleteRoomHandler(logger)).Methods("DELETE")

	// Customer routes
	router.HandleFunc("/customer", controller.CreateCustomerHandler(logger)).Methods("POST")
	router.HandleFunc("/customer", controller.GetCustomerHandler(logger)).Methods("GET")
	router.HandleFunc("/customer", controller.UpdateCustomerHandler(logger)).Methods("PUT")
	router.HandleFunc("/customer", controller.DeleteCustomerHandler(logger)).Methods("DELETE")

	return router
}
