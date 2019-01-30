package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

//Universe type 2 dimentional
type Universe [][]bool

func (u Universe) String() string {

	var b byte

	serialize := make([]byte, 0, (width+1)*height)

	for _, r := range u {

		for _, cell := range r {

			b = ' '
			if cell {
				b = '*'
			}
			serialize = append(serialize, b)
		}
		serialize = append(serialize, '\n')
	}
	return string(serialize)
}

//Set value to slice
func (u Universe) Set(row int, column int, value bool) {
	u[row][column] = value
}

//Seed type 2 dimentional
func (u Universe) Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; (width*height/100)*25 > i; i++ {

		row := rand.Intn(height)
		column := rand.Intn(width)

		if !u[row][column] {
			u.Set(row, column, true)
		} else {
			i--
		}
	}
}

//Show display universe
func (u Universe) Show() {
	fmt.Print(u.String())
}

// NewUniverse create universe
func NewUniverse() Universe {

	u := make(Universe, height)

	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func main() {
	v := NewUniverse()
	v.Seed()
	v.Show()
}
