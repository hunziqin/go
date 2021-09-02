package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "《" + h.name + " - " + strconv.Itoa(h.age) + " years - 电话 " + h.phone + "》"
}

func main() {
	Bob := Human{"Bob", 19, "13132884236"}
	fmt.Println("This Human is : ", Bob)
}
