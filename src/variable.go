package main

import(
	"fmt"
	"unsafe"
)

func main()  {
	var int8 int8 = 127
	var int16 int16 = 32767
	var int32 int32 = 2147483647
	var int64 int64 = 9223372036854775807
	
	fmt.Println(int8)
	fmt.Println(int16)
	fmt.Println(int32)
	fmt.Println(int64)

	fmt.Println("--------------------")

	var uint8 uint8 = 255
	var uint16 uint16 = 65535
	var uint32 uint32 = 4294967295
	var uint64 uint64 = 18446744073709551615
	var num = 25

	fmt.Println(uint8)
	fmt.Println(uint16)
	fmt.Println(uint32)
	fmt.Println(uint64)

	fmt.Printf("%T\n", int8)
	fmt.Println(unsafe.Sizeof(num))
}