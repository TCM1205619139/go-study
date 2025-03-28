package main

import "fmt"

// func main() {
// 	defer func() {
// 		r := recover()
// 		fmt.Println("Recovered in main", r)
// 	}()
// 	a()
// 	fmt.Println("Main end")
// }

// func a() {
// 	defer func() {
// 		r := recover()
// 		fmt.Println("panic recover", r)
// 	}()

// 	panic("a")
// }

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
