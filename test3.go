package main

import "fmt"

func main() {
	for d := 10; d > 0; d-- {
		if d == 6 {
			break
		}
		fmt.Println(d)
	}
	b := 0
	for c := 0; c < 10; c++ {
		b += c
	}
	fmt.Println(b)
}
