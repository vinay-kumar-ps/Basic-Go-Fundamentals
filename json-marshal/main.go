package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "kunu",
		Last:  "vava",
		Age:   1,
	}
	p2 := person{
		First: "tinku",
		Last:  "baby",
		Age:   3,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	vs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(vs))
}
