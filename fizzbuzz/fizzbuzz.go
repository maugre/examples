package main

import "fmt"

const (
	FIZZ = 3
	BUZZ = 5
)

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		default:
			fmt.Println(i)
		case i%FIZZ == 0 && i%BUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		}
	}
}
