package main

import "fmt"

//iterating over arrays and slices
func main() {
	// nums := []int{6, 7, 8}

	// for i, num := range nums {
	// 	fmt.Println(i, num)

	// }

	// m := map[string]string{"name": "Ayan Hasnain", "age": "30"}

	// for k, v := range m {
	// 	fmt.Println(k, v) //name Ayan Hasnain //age 30
	// }

	//unicode codepoint rune
	// i -> starting byte of rune
	// 255 -> 1 byte
	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}
