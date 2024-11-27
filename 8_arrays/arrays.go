package main

import "fmt"

//numbered sequence of specified length

func main() {
	// var nums [4]int
	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3
	// nums[3] = 4
	//array length
	// fmt.Println(len(nums))
	// fmt.Println(nums[3]) //4
	// fmt.Println(nums)

	//declaring array in one line
	// nums := [4]int{1, 2, 3, 4}
	// fmt.Println(nums)

	//2d arrays
	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums) //[[1 2] [3 4]]

	// - fixed sizes
	//- memory optimization
	//- constants time access
}
