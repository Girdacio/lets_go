package main

import "fmt"

/**
[] Arrays
 - A fixed-size collection of elements of the same type.
 - Indexed from 0 to n-1, where n is the size of the array.
 - Contiguous block of memory, which allows for efficient access.
 - Example: var arr [5]int = [5]int{1, 2, 3, 4, 5}
*/
func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("the length is %v with capacity %v\n", len(intSlice), cap((intSlice)))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("the length is %v with capacity %v\n", len(intSlice), cap((intSlice)))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	//var intSlice3 []int32 = make([]int32, 3, 8)
}
