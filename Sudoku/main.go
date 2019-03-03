package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows, columns = 9, 9
	empty         = 0
)

//ErrBounds errors
var (
	ErrBounds    = errors.New("out of Bounds")
	ErrDigit     = errors.New("invalid Digit")
	ErrInRow     = errors.New("Invalid Row")
	ErrInCol     = errors.New("Invalid Column")
	ErrInRegion  = errors.New("Invalid Region")
	ErrFixedCell = errors.New("is Fixed length")
)

//Cell structure
type Cell struct {
	value   uint8
	isFixed bool
}

// SudokuGame board
type SudokuGame struct {
	Grid [rows][columns]Cell
}

func (s *SudokuGame) isFixed(row, column int8) bool {

	return s.Grid[row][column].isFixed
}

func (s *SudokuGame) inColumn(column int8, setValue uint8) bool {

	for i := 0; i < rows; i++ {
		if s.Grid[i][column].value == setValue {
			return true
		}
	}
	return false
}

func (s *SudokuGame) inRow(row int8, setValue uint8) bool {

	for i := 0; i < columns; i++ {
		if s.Grid[row][i].value == setValue {
			return true
		}
	}
	return false
}

func (s *SudokuGame) clear(row int8, column int8) error {

	switch {
	case !inBounds(row, column):
		return ErrBounds
	case s.isFixed(row, column):
		return ErrFixedCell
	}

	s.Grid[row][column].value = empty

	return nil
}

func (s *SudokuGame) set(row int8, column int8, setValue uint8) error {

	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(setValue):
		return ErrBounds
	case s.isFixed(row, column):
		return ErrFixedCell
	case s.inColumn(column, setValue):
		return ErrInCol
	case s.inRow(row, setValue):
		return ErrInRow
	case s.inRegion(row, column, setValue):
		return ErrInRegion
	}

	s.Grid[row][column].value = setValue
	return nil
}

func (s *SudokuGame) inRegion(row, column int8, setValue uint8) bool {

	startRow, startColumn := row/3*3, column/3*3

	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if s.Grid[r][c].value == setValue {
				return true
			}
		}
	}
	return false
}

func inBounds(row, column int8) bool {

	if row >= rows || row < 0 || column >= columns || column < 0 {
		return false
	}

	return true
}

func validDigit(setValue uint8) bool {

	if setValue >= 1 && setValue <= 9 {
		return true
	}
	return false
}

//NewSudoKu create  board
func NewSudoKu(board [rows][columns]uint8) *SudokuGame {

	var initialGrid [rows][columns]Cell

	for i := 0; i < rows; i++ {
		for k := 0; k < columns; k++ {
			cellValue := board[i][k]
			if cellValue != empty {
				initialGrid[i][k] = Cell{value: cellValue, isFixed: true}
			}
		}
	}

	return &SudokuGame{Grid: initialGrid}
}

func main() {

	s := NewSudoKu([rows][columns]uint8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.set(1, 2, 3)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range s.Grid {
		fmt.Println(row)
	}
}
