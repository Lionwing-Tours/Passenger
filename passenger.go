package main

import (
	"net/http"
	"passenger/handlers"
	"passenger/jwt"
)

func main() {
	http.HandleFunc("/setToken", jwt.SetToken)
	http.HandleFunc("/show", jwt.Show)
	http.HandleFunc("/accounts/register", handlers.Register)
	http.HandleFunc("/accounts/login", handlers.Login)
	http.HandleFunc("/accounts/logout", handlers.Logout)
	http.HandleFunc("/accounts/verify", handlers.Verify)
	http.HandleFunc("/searchVehicles", handlers.SearchVehicles)
	http.HandleFunc("/bookings/add", handlers.AddBooking)
	http.HandleFunc("/bookings/unassign", handlers.UnassignVehicle)
	http.HandleFunc("/bookings/edit", handlers.EditBooking)
	http.HandleFunc("/bookings/view", handlers.ViewBooking)
	http.HandleFunc("/bookings/delete", handlers.DeleteBooking)

	http.ListenAndServe("localhost:3000", nil)

}
