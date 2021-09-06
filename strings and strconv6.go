package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "The hunzi yangshengdui chennenghong"
	s1 := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", s1)
	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	var str1 string = "Hunzi|yang|chen"
	s2 := strings.Split(str1, "|")
	fmt.Printf("Splitted in slice: %v\n", s2)
	for _, val := range s2 {
		fmt.Printf("%s -", val)
	}
	fmt.Println()
	str2 := strings.Join(s2, ";")
	fmt.Printf("s2 joined by ;: %s\n", str2)
}
