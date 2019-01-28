package main

import (
	"fmt"
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
	v.Show()
}
