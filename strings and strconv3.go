package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello,how is it going, Hunzi?"
	var manyhunzi string = "hunzihunzihunzi"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is:", manyhunzi)
	fmt.Printf("%d\n", strings.Count(manyhunzi, "hunzi"))
}
