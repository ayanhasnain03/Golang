package main

import (
	"fmt"
)

func main() {
	//slices are like arrays but they are dynamic in size
	//mostly used for arrays in golang
	// + useful methods

	//unintialised slice is nill
	// var nums []int

	// fmt.Println(nums == nil) //true

	// var nums = make([]int, 2, 5)

	// fmt.Println(nums) //[0,0]
	// fmt.Println(nums == nil) //false

	//capacity -> maximum size of the slice
	// fmt.Println(cap(nums)) //5

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)

	// fmt.Println(nums)
	// fmt.Println(cap(nums)) //capacity dynamically increase

	// Create a slice with initial length 2 and capacity 5

	// var nums = make([]int, 0, 5)
	// Append a new element to nums

	// nums = append(nums, 2)

	// Create another slice with the same length as nums
	// var nums2 = make([]int, len(nums))

	// Copy elements from nums to nums2
	// copy(nums2, nums)

	// Print the two slices
	// fmt.Println(nums, nums2)

	//slice operator
	// var nums = []int{1, 2, 3}

	// fmt.Println(nums[0:2]) //[1,2]
	// fmt.Println(nums[:2]) //[1,2]
	// fmt.Println(nums[1:]) //[2,3]

	//slice

	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}

	// fmt.Println(slices.Equal(nums1, nums2)) //true

	//multidimensional slice
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
}
