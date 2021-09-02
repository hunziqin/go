package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var hunzi person
	hunzi.name, hunzi.age = "huzni", 19
	chen := person{"chen", 18}
	yang := person{age: 17, name: "yang"}
	hc_Older, hc_diff := Older(hunzi, chen)
	hy_Older, hy_diff := Older(hunzi, yang)
	cy_Older, cy_diff := Older(chen, yang)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", hunzi.name, chen.name, hc_Older.name, hc_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", hunzi.name, yang.name, hy_Older.name, hy_diff)
	fmt.Printf("Of %s and %s, %s is Older by %d years\n", chen.name, yang.name, cy_Older.name, cy_diff)
}
