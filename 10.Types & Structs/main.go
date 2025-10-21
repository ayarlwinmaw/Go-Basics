package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "John",
		LastName:  "Smith",
	}

	log.Println("User's first name is", user.FirstName)
	log.Println("User's last name is", user.LastName)
	log.Println("Phone number", user.PhoneNumber)
	log.Println("Age", user.Age)
	log.Println("Birthdate", user.BirthDate)
}
