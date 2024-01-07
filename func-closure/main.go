package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height float64
	weight float64
}

// Function type calculator
type compare func(int, int) bool

func main() {
	fmt.Println("Demo : function closure")

	people := []Person{
		{
			name:   "Abc",
			height: 5.8,
			weight: 100,
		},
		{
			name:   "Xyz",
			height: 5.0,
			weight: 60,
		},
		{
			name:   "Pqr",
			height: 4.3,
			weight: 74,
		},
		{
			name:   "Mnr",
			height: 6.1,
			weight: 95,
		},
	}

	by := func(property string) compare {
		return func(i, j int) bool {
			switch property {
			case "height":
				return people[i].height > people[j].height
			case "weight":
				return people[i].weight > people[j].weight
			}
			return people[i].name > people[j].name
		}
	}

	fmt.Printf("Before sorting %v\n", people)
	sort.Slice(people, by("height"))
	fmt.Printf("After sorting by height %v\n", people)
	sort.Slice(people, by("weight"))
	fmt.Printf("After sorting by weight %v\n", people)
}
