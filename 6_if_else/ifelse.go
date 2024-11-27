package main

import "fmt"

func main() {
	// if else statement in golang
	// age := 17

	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// } else {
	// 	fmt.Println("You are not an adult")
	// }

	const role = "admin"

	if role == "admin" {
		fmt.Println("Welcome Boss !")
	} else {
		fmt.Println("You are not admin")
	}

}
