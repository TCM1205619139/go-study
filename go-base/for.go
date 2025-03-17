package main

import (
	"fmt"
	"math"
)

func add(a, b int, c, d string) (int, string) {
	return a + b, c + d
}

func swap(a int, b int) {
	println("[func|swap]a=", a, "b=", b)
	a, b = b, a
	println("[func|swap]a=", a, "b=", b)
}

func swapRef(pa *int, pb *int) {
	println("[func|swapRef]a=", *pa, "b=", *pb)
	var temp = *pa
	*pa = *pb
	*pb = temp
	println("[func|swapRef]a=", *pa, "b=", *pb)
}

func getSquaredRoot1(num float64) float64 {
	return math.Sqrt(num)
}

var getSquaredRoot2 = func(x float64) float64 {
	return math.Sqrt(x)
}

type call_back func(int) int

func calcValue(x int, callback call_back) int {
	return callback(x)
}

func realFunc(x int) int {
	return x * x
}

func main() {
	// var index uint64 = 0
	var array = []int{1, 3, 5, 7, 9}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	// for {
	// 	index += 1
	// 	fmt.Println("Infinite loop %d", index)
	// }

	fmt.Print("---------------------------------\n")

	a, b := 1, 2
	c, d := "c", "d"
	res1, res2 := add(a, b, c, d)
	println("res1=", res1, "res2=", res2) // 3, cd

	println("[func|main]a=", a, "b=", b) // 1, 2
	swap(a, b)
	println("[func|main]a=", a, "b=", b) // 1, 2

	println("[func|main]a=", a, "b=", b) // 1, 2
	swapRef(&a, &b)
	println("[func|main]a=", a, "b=", b) // 2, 1

	// LOOP:
	// 	println("Enter your age:")

	// 	var age uint8

	// 	_, err := fmt.Scan(&age)
	// 	if err != nil {
	// 		println("Invalid input")
	// 		goto LOOP
	// 	}

	// 	println(&age, age)
	// 	if age < 18 {
	// 		println("you are not eligible to vote")
	// 		goto LOOP
	// 	} else {
	// 		println("you are eligible to vote")
	// 	}

	// 	println("Thank you for voting")

}
