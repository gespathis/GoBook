package main

import (
	"fmt"
)

type item struct {
	name string
}
type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {

	if c != nil && i != nil {
		c.leftHand = i
		fmt.Printf("%v picks up a %v\n", c.name, i.name)
	}
}

func (c *character) give(to *character) {

	if c == nil || to == nil {
		return
	}

	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
		return
	}

	if to.leftHand != nil {
		fmt.Printf("%v is holding something else\n", to.name)
		return
	}

	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gives %v a %v\n", c.name, to.name, to.leftHand.name)

}

func (c character) String() string {

	if c.leftHand == nil {
		return fmt.Sprintf("%v is not carrying anything", c.name)
	}

	return fmt.Sprintf("%v is carrying a %v", c.name, c.leftHand.name)
}

func main() {

	arthur := &character{name: "Arthur"}

	item1 := &item{name: "axe"}

	arthur.pickup(item1)
	knight := &character{name: "lanselot"}

	arthur.give(knight)
	//item1.name = "matrix axe"
	fmt.Println(arthur)
	fmt.Println(knight)
}
