package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi,hunzi"
	var low string
	var up string

	fmt.Printf("The original string is: %s\n", str)
	low = strings.ToLower(str)
	fmt.Printf("The low string is: %s\n", low)
	up = strings.ToUpper(str)
	fmt.Printf("The up string is: %s\n", up)
}
