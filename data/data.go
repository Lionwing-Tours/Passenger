package data

import (
	"errors"
	"log"
	"passenger/db"
	"time"
)

var (
	ErrUserDoesNotExist = errors.New("User Does Not Exist")
	ErrInvalidPassword  = errors.New("Invalid Password")
)

const (
	getVehiclesQuery = `SELECT vehicle.vehicle_id, 
								vehicle.make, 
								vehicle.model, 
								vehicle.year, 
								vehicle.no_passengers, 
								vehicle.no_luggages,
								vehicle_travel_rates.per_day, 
								vehicle_travel_rates.mileage_limit, 
								vehicle_travel_rates.extra_km_charge,
								vehicle.rating, 
								driver.rating, 
								vehicle.image_string
							FROM vehicle 
								INNER JOIN vehicle_travel_rates
									ON vehicle.vehicle_id=vehicle_travel_rates.vehicle_id
								INNER JOIN driver
									ON driver.driver_id=vehicle_travel_rates.driver_id 
								INNER JOIN driver_availability
									ON driver_availability.driver_id=driver.driver_id
										WHERE vehicle.location=? AND
										driver.driver_id=vehicle.driver_id AND
										 ? < driver_availability.not_available_from
											AND ? < driver_availability.not_available_from
										OR ? > driver_availability.not_available_to 
											AND ? > driver_availability.not_available_to`

	insertUserQuery = `INSERT INTO passengers 
						(salutation, name, lastname, email, 
						sec_password, phone, address, 
						created_date, updated_date, last_login,
						user_status, created_by, country_id) 
						VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
)

type VehicleInfo struct {
	Id                 int
	Make               string
	Model              string
	Year               string
	NumberOfPassengers int
	NumberOfLuggages   int
	PerDay             float64
	MileageLimit       int
	ExtraKmCharge      float64
	VehicleRating      int
	DriverRating       int
	ImageString        string
}

func GetVehicles(startLocation string, pickupDate, dropoffDate string, vehicleType int) (vehiclesInfo []VehicleInfo, err error) {
	pickDate, _ := time.Parse("2006-01-02 00:00:00", pickupDate)
	dropDate, _ := time.Parse("2006-01-02 00:00:00", dropoffDate)

	rows, err := db.MySqlDB.Query(getVehiclesQuery, startLocation, pickDate, dropDate, pickDate, dropDate)
	defer rows.Close()
	if err != nil {
		log.Println("Error fetching vehicle info: ", err)
		return vehiclesInfo, err
	}

	return vehiclesInfo, nil
}
