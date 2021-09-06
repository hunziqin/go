package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi,hunzi "
	var new string

	new = strings.Repeat(str, 6)
	fmt.Printf("The new repeated string is: %s\n", new)
}
