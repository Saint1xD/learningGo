package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	const a = 10
	var b = 5
	const a string = "hello"
	fmt.Println(a + b)
	fmt.Println("%s hi", a)
}
