package main

import "fmt"

//map -> hash,object

func main() {
	//creating map

	// m := make(map[string]string)

	//assign value
	// m["name"] = "Ayan Hasnain"
	// m["age"] = "30"

	//get value

	// fmt.Println(m["name"], m["age"])

	//if key does not exist then value is nil
	// fmt.Println(m["name2"])

	// m := map[string]int{
	// 	"apple":  10,
	// 	"banana": 20,
	// 	"orange": 30,
	// }

	// fmt.Println(m["apple"]) //10
	// fmt.Println(len(m)) //3

	//delete
	// delete(m, "apple")//apple will be deleted

	// fmt.Println(m)//map[banana:20 orange:30]

	// clear(m)
	// fmt.Println(m) //map[]

	//creating map best way
	// m := map[string]int{
	// 	"apple":  10,
	// 	"banana": 20,
	// 	"orange": 30,
	// }

	// fmt.Println(m) //map[apple:10 banana:20 orange:30]

	//accessing map
	m := map[string]int{
		"phone": 3,
		"price": 20,
	}
	v, ok := m["phone"]
	if ok {
		fmt.Println(v) //3
	} else {
		fmt.Println("not found")
	}

}
