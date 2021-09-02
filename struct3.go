package main

import "fmt"

type Skills []string

type Hunmen struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Hunmen
	Skills
	int
	speciality string
}

func main() {
	hunzi := Student{Hunmen: Hunmen{"hunzi", 19, 149}, speciality: "biology"}
	fmt.Println("His name is", hunzi.name)
	fmt.Println("His age is", hunzi.age)
	fmt.Println("His weight is", hunzi.weight)
	fmt.Println("His speciality is", hunzi.speciality)
	hunzi.Skills = []string{"anatomy"}
	fmt.Println("His skills are", hunzi.Skills)
	fmt.Println("His acquired tow new ones ")
	hunzi.Skills = append(hunzi.Skills, "physics", "goland")
	fmt.Println("His skills now are ", hunzi.Skills)
	hunzi.int = 19
	fmt.Println("His preferred number is", hunzi.int)
}
