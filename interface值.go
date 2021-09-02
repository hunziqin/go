package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s.Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	hun := Student{Human{"hun", 19, "13132884236"}, "MIT", 0.00}
	zi := Student{Human{"zi", 20, "18269641207"}, "Harvard", 100}
	yang := Employee{Human{"yang", 21, "15977583284"}, "Goland Inc", 1000}
	chen := Employee{Human{"chen", 22, "1525294255"}, "Things Ltd", 5000}

	var i Men
	i = hun
	fmt.Println("This is hun,a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = chen
	fmt.Println("This is chen,an Employee")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let is use a silce of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = zi, yang, hun

	for _, value := range x {
		value.SayHi()
	}
}
