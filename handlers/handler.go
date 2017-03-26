package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"passenger/data"
	"strconv"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	title := r.FormValue("title")
	firstName := r.FormValue("fname")
	lastName := r.FormValue("lname")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	createdBy := r.FormValue("created_by")
	country, _ := strconv.Atoi(r.FormValue("country_id"))

	//validate email
	err := RegisterUser(title, firstName, lastName, email, password, phone, address, createdBy, country)
	if err != nil {
		log.Println("Error creating user")
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created"))
}

func SearchVehicles(w http.ResponseWriter, r *http.Request) {
	startLocation := r.FormValue("start-location")
	pickupDate := r.FormValue("pickup-date")
	dropoffDate := r.FormValue("dropoff-date")
	vehicleType, _ := strconv.Atoi(r.FormValue("vtype"))

	vehicleData, err := data.GetVehicles(startLocation, pickupDate, dropoffDate, vehicleType)
	if err != nil {
		return
	}

	jsn, _ := json.Marshal(vehicleData)
	jsnStr := string(jsn)
	w.Write([]byte(jsnStr))

	//send vehicleData as json array
}

func Login(w http.ResponseWriter, r *http.Request) {
	// API call to Booking service
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// API call to Booking service
}

func Verify(w http.ResponseWriter, r *http.Request) {
	// API call to Booking service
}

func AddBooking(w http.ResponseWriter, r *http.Request) {
	//	passengerId, _ := strconv.Atoi(r.FormValue("id"))
	//	driverId, _ := strconv.Atoi(r.FormValue("driver_id"))
	//	vehicleId, _ := strconv.Atoi(r.FormValue("vehicle_id"))
	//	startLoc := r.FormValue("start_loc")
	//	startDate := r.FormValue("start_date")
	//	endDate := r.FormValue("end_date") //
	//	noOfPassengers, _ := strconv.Atoi(r.FormValue("passengers"))
	//	estCost, _ := strconv.ParseFloat(r.FormValue("est_cost"), 64)
	// API call to Booking service
}

func UnassignVehicle(w http.ResponseWriter, r *http.Request) {

}

func EditBooking(w http.ResponseWriter, r *http.Request) {
	//	passengerId := r.FormValue("passengerId")
	//	startDate := r.FormValue("startDate")
	//	endDate := r.FormValue("endDate")
	//	pickupAddress := r.FormValue("pickupAddress")
	//	numPassengers := r.FormValue("numPassengers")
	//	pickupLocation := r.FormValue("pickupLocation")
	//	dropLocation := r.FormValue("dropLocation")
	//	vehicleType := r.FormValue("vehicleType")
	//	driverId := r.FormValue("driverId")

	// API call to Booking Service
}

func ViewBooking(w http.ResponseWriter, r *http.Request) {
	//bookingId := r.FormValue("bookingId")
	// API call to Booking service

}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	//bookingId := r.FormValue("bookingId")
	// API call to Booking service
}
