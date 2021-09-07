package main

import "fmt"

func main() {
	var a string = "hunzi"
	var p *string = &a
	*p = "yangchen"
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the pointer *p: %s\n", *p)
	fmt.Printf("Here is the pointer a: %s\n", a)
}
