package main

import "fmt"

func main() {
	a := 10
Here:
	fmt.Println(a)
	a++
	goto Here
}
