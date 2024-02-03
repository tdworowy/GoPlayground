package main

import "fmt"

func main() {
	var a = 0
	var b = 1
	for i := 0; i < 1000; i++ {
		a, b = b, a+b
	}
	fmt.Println(b)
}
