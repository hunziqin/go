package main

import "fmt"

type Hunmen struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Hunmen
	school string
}

type Employee struct {
	Hunmen
	company string
}

func (h *Hunmen) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	hunzi := Student{Hunmen{"hunzi", 19, "13132884236"}, "MIT"}
	yang := Employee{Hunmen{"yang", 18, "18269641207"}, "goland inc"}

	hunzi.SayHi()
	yang.SayHi()
}
