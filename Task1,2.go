package Task12

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
	//task1
	p.Info()
	fmt.Println("--------------------")
	planeInfo(p)

	fmt.Println("--------------------")
	//task2
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
		fmt.Println(i,"|",planes[i].manufacturer,planes[i].model)
	}
}
