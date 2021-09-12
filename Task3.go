package main

import "fmt"

type Writer interface{
	Output(i int) string
}

type Informator interface{
	Info()
}

func about(informator Informator){
	informator.Info()
}

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

func (p Passenger) Info(){
	fmt.Println("name:",p.name,"\nsurname:",p.surname,"\nseat number:",p.seatNumber)
}

func planeInfo(p Plane) {
	fmt.Println("manufacturer:", p.manufacturer, "\nmodel:", p.model,
		"\nseats:", p.seats, "\nmaximum speed:", p.maxSpeed)
}

func (p Plane) Output(i int) string{
	
	return fmt.Sprintln(i,"|",p.manufacturer,p.model)
}

func Write(w Writer, i int){
	fmt.Println(w.Output(i))
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
	about(p)
	fmt.Println("--------------------")
	about(p.passengerInfo)

	fmt.Println("--------------------")

	var planes [3]Plane=[3]Plane{
		p,
		{
			manufacturer:  "Airbus",
			model:         "A320",
			seats:         165,
			maxSpeed:      845,
			passengerInfo: Passenger{
				name:       "Dmitry",
				surname:    "Leonov",
				seatNumber: "7A",
			},
		},
		{
			manufacturer:  "Cessna",
			model:         "172",
			seats:         2,
			maxSpeed:      302,
			passengerInfo: Passenger{
				name:       "Dmitry",
				surname:    "Leonov",
				seatNumber: "7A",
			},
		},
	}
	for i:=0;i<len(planes);i++{
		Write(planes[i],i)
	}
}
