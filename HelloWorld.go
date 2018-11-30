// package main

// import (
// 	"fmt"
// )

// //*****How to run
// // Open the terminal or cmd prompt, CD to helloworld.go dir, Run the below command
// // go run HelloWorld.go

// //alias print method
// var print = fmt.Println

// func main() {
// 	display("hello world!")
// }

// func display(text string){
// 	print(text)
// }

package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)
}

func sum(xi ...int) int {
	fmt.Println("I am from sum")
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
