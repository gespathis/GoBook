package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print(u.String())
}

//Alive check if a planet is alive
func (u Universe) Alive(x, y int) bool {
	y = (y + height) % height
	x = (x + width) % width
	return u[y][x]
}

// NewUniverse create universe
func NewUniverse() Universe {

	u := make(Universe, height)

	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Neighbors coount live neiboors
func (u Universe) Neighbors(x, y int) int {

	n := 0

	for i := -1; i <= 1; i++ {
		for k := -1; k <= 1; k++ {
			if u.Alive(x+k, y+i) && !(k == 0 && i == 0) {
				n++
			}
		}
	}
	return n
}

//Next the state of planet at next generation
func (u Universe) Next(x, y int) bool {

	n := u.Neighbors(x, y)
	return n == 2 || n == 3 && u.Alive(x, y)
}

//Step create the next generation
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(y, x, a.Next(x, y))
		}
	}
}

func main() {
	a := NewUniverse()
	b := NewUniverse()
	a.Seed()

	//slice testing
	//c := []string{}
	//c = append(c, "123")
	//fmt.Print(c)
	//test := make([]string, 0, 10)
	//test = append(test, "addone")
	//fmt.Print(cap(test))
	//println(time.Second, time.Millisecond,time.Microsecond, time.Nanosecond)

	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()

		time.Sleep(time.Second / 30)

		a, b = b, a
	}
	//fmt.Println(v.Neighbors(1, 1))
}
