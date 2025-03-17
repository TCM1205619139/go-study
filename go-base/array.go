package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 100.00
	result1 := getSquaredRoot1(num)
	result2 := getSquaredRoot2(num)

	fmt.Println("result1=", result1, "result2=", result2)

	value := 81

	result3 := calcValue(value, realFunc)
	fmt.Println("result3=", result3)

	array1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for index := range array1 {
		// array[index]类型是一维数组
		fmt.Println(reflect.TypeOf(array1[index]), "%T", array1[index])
		fmt.Printf("index=%d, value=%v\n", index, array1[index])
	}
}
