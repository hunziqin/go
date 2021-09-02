package main

import "fmt"

type Humen struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Humen
	speciality string
}

func main() {
	hunzi := Student{Humen{"hunzi", 19, 149}, "computer science"}
	fmt.Println("His name is ", hunzi.name)
	fmt.Println("His age is ", hunzi.age)
	fmt.Println("His weight is ", hunzi.weight)
	fmt.Println("His speciality is ", hunzi.speciality)
	hunzi.speciality = "AI"
	fmt.Println("hunzi change his speciality")
	fmt.Println("His speciality is ", hunzi.speciality)
	fmt.Println("hunzi become old")
	hunzi.age = 20
	fmt.Println("His age is", hunzi.age)
	fmt.Println("hunzi is not an athlet anymore")
	hunzi.weight += 1
	fmt.Println("His weight is", hunzi.weight)
}
