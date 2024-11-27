package main

import "fmt"

func main() {
	// for -> the only construct in Go for looping

	//While loop
	// i := 1

	// for i <= 3 {
	// 	fmt.Println(i) // Using fmt.Println for consistent printing
	// 	i = i + 1
	// }

	//infinite loop

	// for{
	// 	println("hello")
	// }

	//classic for loop

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
