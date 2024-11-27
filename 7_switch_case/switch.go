package main

import "fmt"

// import (
// 	"fmt"
// 	"time"
// )

func main() {
	//simple switch
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")

	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Workday")
	// }

	//type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoAmI(true)
}
