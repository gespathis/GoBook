package main

import (
	"fmt"
)

type surfaceMovement interface {
	moveUp()
	moveDown()
	moveRight()
	moveLeft()
	showCoordinates()
}

type groundAnimal struct {
	x, y int
}

func (t *groundAnimal) moveUp() {
	(*t).y--
}

func (t *groundAnimal) moveDown() {
	t.y++
}

func (t *groundAnimal) moveRight() {
	t.x++
}

func (t *groundAnimal) moveLeft() {
	t.x--
}

func (t *groundAnimal) showCoordinates() {
	fmt.Printf("[x:%2v y:%2v] \n", t.x, t.y)
}

type turtle struct {
	surfaceMovement
}

func main() {

	newturtle := turtle{surfaceMovement: &groundAnimal{}}

	if newturtle.surfaceMovement != nil {
		newturtle.moveUp()
		newturtle.moveUp()
		newturtle.moveUp()
		newturtle.moveLeft()
		newturtle.moveLeft()
		newturtle.moveRight()
		newturtle.moveDown()
		newturtle.moveDown()
	}

	fmt.Printf("turtle move is %5v \n", "slowwwww")
	newturtle.showCoordinates()
}
