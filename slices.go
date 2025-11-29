package main

import (
	"fmt"
	"slices"
)

func main() {
	var numbers []int
	var numbers1 []int = []int{1, 2, 3, 4}
	slice := make([]int, 4)
	fmt.Println(numbers)
	fmt.Println(numbers1)
	fmt.Println(slice)
	var int_array [5]int = [5]int{1, 2, 3, 4, 5}
	var int_slice = int_array[0:]
	fmt.Println(int_slice)
	int_slice = append(int_slice, 6, 7)
	fmt.Println(int_slice)

	sliceCopy := make([]int, len(int_slice))
	copy(sliceCopy, int_slice) //-> deep copy

	// sliceCopy := int_slice -> copy by reference

	fmt.Println("The copied slice is : ", sliceCopy)

	for idx, val := range sliceCopy {
		fmt.Printf("The value sliceCopy[%d]: %d\n", idx, val)
	}

	int_slice[3] = 50
	fmt.Println(int_slice)
	fmt.Println(sliceCopy)

	if slices.Equal(int_slice, sliceCopy) {
		fmt.Println("Equal")
	}

	// twoD := make([][]int, 3, 5)

	// for i := range 3 {
	// 	for j := range 5 {
	// 		twoD[i][j] = 0
	// 	}
	// }
	// fmt.Println(twoD)

	slice2 := sliceCopy[2:4]
	fmt.Println(slice2)

}
