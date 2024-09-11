package main

import "time"

func main() {

	// Switch
	// No need to write break statement
	// No need to write default
	day := 0
	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid Day")
	}

	// Switch with multiple cases
	var today = time.Now().Weekday()
	switch today {
	case time.Saturday, time.Sunday:
		println("Weekend")
	default:
		println("Weekday")
	}

	// Type Switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("I am an integer")
		case string:
			println("I am a string")
		case float32:
			println("I am a float")
		default:
			println("I am something else")
		}
	}
	whoAmI(1)
	whoAmI("Anish")
	whoAmI(1.1)
	whoAmI(true)

}