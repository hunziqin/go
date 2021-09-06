package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "666"
	var an int
	var new string

	fmt.Printf("%d 位系统\n", strconv.IntSize)
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(str)
	fmt.Printf("The intrger is: %d\n", an)
	an = an + 5
	new = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", new)
}
