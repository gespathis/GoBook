package main

import (
	"GoBook/MartianAnimalSanctuary/Domain"
	"fmt"
	"math/rand"
	"time"
)

func action(a ark.Animal) {

	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v eats %v.\n", a, a.Eat())
	case 1:
		fmt.Printf("%v is %v.\n", a, a.Move())
	}
}

const sunrise, sunset = 7, 19

func main() {

	rand.Seed(time.Now().UnixNano())
	animals := make([]ark.Animal, 0, 2)
	animals = append(animals, ark.Lion{Name: "Simba"})
	animals = append(animals, ark.Deer{Name: "Bambi"})

	var hour int

	for sol := 1; sol <= 3; {
		fmt.Printf("%02d:00 - ", hour)

		if hour < sunrise || hour >= sunset {
			fmt.Println("Animals are sleeping")
		} else {
			i := rand.Intn(len(animals))
			action(animals[i])
		}
		time.Sleep(time.Second)
		hour++

		if hour >= 24 {
			hour = 0
			sol++
		}
	}

	//fmt.Printf("lets begin with %5v and %5v \n", carnival.Name, vegeterian.Name)
}
