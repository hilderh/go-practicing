package main

import "fmt"

func main() {
	fmt.Println("Hello world print")
	foo()
	fmt.Println("I want to print some for loop")
	for i := 0; i < 100; i++ {
		fmt.Println("I'm a number in loop")
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("Is pair")
		} else {
			fmt.Println("Is unpair")
		}
	}
}
func foo() {
	fmt.Println("Hello what are you doing???? foo ")
}
