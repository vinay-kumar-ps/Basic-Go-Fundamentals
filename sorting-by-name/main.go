package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Q", 64}
	p3 := Person{"M", 56}
	p4 := Person{"Moneypenny", 27}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}
