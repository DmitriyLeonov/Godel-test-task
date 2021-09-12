package main

import "fmt"

type Passenger struct {
	name       string
	surname    string
	seatNumber string
}

type Plane struct {
	manufacturer  string
	model         string
	seats         int
	maxSpeed      int
	passengerInfo Passenger
}

func (p Plane) Info() {
	fmt.Println("manufacturer:", p.manufacturer, "\nmodel:", p.model,
		"\nseats:", p.seats, "\nmaximum speed:", p.maxSpeed)
}

func planeInfo(p Plane) {
	fmt.Println("manufacturer:", p.manufacturer, "\nmodel:", p.model,
		"\nseats:", p.seats, "\nmaximum speed:", p.maxSpeed)
}

func main() {
	p := Plane{
		manufacturer: "Boeing",
		model:        "737-800",
		seats:        160,
		maxSpeed:     946,
		passengerInfo: Passenger{
			name:       "Dmitry",
			surname:    "Leonov",
			seatNumber: "7A",
		},
	}
	p.Info()
	planeInfo(p)
}
