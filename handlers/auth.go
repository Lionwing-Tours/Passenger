package handlers

import (
	"database/sql"
	"errors"
	"log"
	"passenger/db"
	"time"
)

const (
	insertUserQuery = `INSERT INTO passengers 
						(salutation, name, lastname, email, 
						sec_password, phone, address, 
						created_date, updated_date, last_login,
						user_status, created_by, country_id) 
						VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
)

var (
	ErrUserDoesNotExist = errors.New("User Does Not Exist")
	ErrInvalidPassword  = errors.New("Invalid Password")
)

func AuthenticateUser(email, password string) (string, error) {
	var Email string
	var Password string

	row := db.MySqlDB.QueryRow("SELECT email FROM passengers WHERE email = ?", email)
	err := row.Scan(&Email)
	if err != nil {
		if err == sql.ErrNoRows || len(Email) == 0 {
			log.Println("ERROR fetching passenger device details from DB:", err)
			return "", ErrUserDoesNotExist
		}
	}

	row = db.MySqlDB.QueryRow("SELECT password FROM passengers WHERE passengers.email = ?", email)
	err = row.Scan(&password)
	if err != nil {
		log.Println("Undefined Error")
		return "", err
	}
	return Password, nil
}

func RegisterUser(salutation, name, lastname, email, sec_password, phone, address, created_by string, country_id int) error {
	_, err := db.MySqlDB.Exec(insertUserQuery, salutation, name, lastname,
		email, sec_password, phone, address, time.Now(), time.Now(), time.Now(),
		"A", created_by, country_id)

	if err != nil {
		log.Println("Error Inserting booking: ", err)
		return err
	}

	return nil
}
