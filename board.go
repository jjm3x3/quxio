package main

import (
	"errors"
	"fmt"
)

type board struct {
	grid [][]rune
	turn bool
}

func (self *board) whoseTurn() rune {
	if self.turn {
		return 'X'
	} else {
		return 'O'
	}
}

func newBoard(layout [][]rune) *board {
	var newBoard board
	if layout != nil {
		newBoard = board{grid: layout, turn: true}
	} else {
		newBoard = board{grid: [][]rune{
			[]rune{'#', '#', '#', '#', '#'},
			[]rune{'#', '#', '#', '#', '#'},
			[]rune{'#', '#', '#', '#', '#'},
			[]rune{'#', '#', '#', '#', '#'},
			[]rune{'#', '#', '#', '#', '#'}}, turn: true}
	}
	return &newBoard
}

func (self *board) checkForWin() rune {
	leads := make([]rune, 12)

	for i := 0; i < len(self.grid); i++ {
		for j := 0; j < len(self.grid); j++ {
			if i == 0 && j == 0 {
				leads[0] = self.grid[i][j]
				leads[5] = self.grid[i][j]
				leads[10] = self.grid[i][j]
			} else if i == 0 {
				if j == 4 {
					leads[11] = self.grid[i][j]
				}
				leads[j+5] = self.grid[i][j]
				if leads[0] != self.grid[i][j] {
					leads[0] = '#'
				}
			} else if j == 0 {
				leads[i] = self.grid[i][j]
				if leads[j+5] != self.grid[i][j] {
					leads[j+5] = '#'
				}
			} else {
				if leads[i] != self.grid[i][j] {
					leads[i] = '#'
				}
				if leads[j+5] != self.grid[i][j] {
					leads[j+5] = '#'
				}
				if i == j {
					if leads[10] != self.grid[i][j] {
						leads[10] = '#'
					}

				}
			}
			if i+j == 4 {
				if leads[11] != self.grid[i][j] {
					leads[11] = '#'
				}

			}
			// fmt.Printf("what is the val of %c\n", leads)
		}
	}

	result := '#'

	for i := range leads {
		// fmt.Printf("What does this look like: %c\n", leads[i])
		if leads[i] != '#' {
			if self.whoseTurn() == leads[i] {
				return leads[i]
			} else {
				result = leads[i]
			}
		}

	}

	return result
}

func (self *board) getPosition(i, j int) rune {
	return self.grid[i][j]
}

func (self *board) printBoard() {
	fmt.Printf("         \n")
	fmt.Printf("   01234 \n")
	fmt.Printf("         \n")
	for j := range self.grid {
		fmt.Printf("%d ", 10+j)
		for i := range self.grid[0] {
			fmt.Printf("%c", self.getPosition(j, i))
		}
		fmt.Printf(" %d\n", 15+j)
	}
	fmt.Printf("         \n")
	fmt.Printf("   56789 \n")
	fmt.Printf("         \n")
}

func isBottom(move int) bool {
	return move >= 5 && move < 10
}

func isTop(move int) bool {
	return move >= 0 && move < 5
}

func (self *board) makeMove(pos int) error {
	theMove := self.whoseTurn()
	// fmt.Printf("move to?: %v\n", pos)
	// fmt.Printf("wats the move %c\n", self.whoseTurn())

	var row int
	if isTop(pos) || isBottom(pos) {
		col := pos
		row = 0
		if isBottom(pos) {
			col = pos - 5
			row = 4
		}
		// fmt.Printf("where are we moving? (%v,%v)\n", row, col)

		oldPeice := self.grid[row][col]
		if oldPeice != theMove && oldPeice != '#' {
			return errors.New("that is an illegal Move")
		}
		for row = 0; row < 5; row++ {
			rowCur := row
			if isBottom(pos) {
				rowCur = 4 - row
				// fmt.Printf("oporate on which row: %v\n", row)
			}
			restOfRow := self.grid[rowCur][col+1 : len(self.grid[0])]
			oldPeice = self.grid[rowCur][col]
			startOfRow := self.grid[rowCur][0:col]
			var newRow []rune
			newRow = append(newRow, startOfRow...)
			if (isTop(pos) && rowCur == 4) || isBottom(pos) && rowCur == 0 {
				newRow = append(newRow, theMove)
			} else {
				var newPeice rune
				if isBottom(pos) {
					newPeice = self.grid[rowCur-1][col]
				} else {
					newPeice = self.grid[rowCur+1][col]
				}
				newRow = append(newRow, newPeice)
			}
			newRow = append(newRow, restOfRow...)
			self.grid[rowCur] = newRow
		}

	}

	if pos >= 10 && pos < 15 {
		row = row - 10

		// fmt.Printf("%v\n", self.grid[thing])
		var newRow []rune
		oldPeice := self.grid[row][0]
		if oldPeice != theMove && oldPeice != '#' {
			// fmt.Printf("what is this thing: %c\n", oldPeice)
			return errors.New("that is an illegal Move")
		}
		newRow = append(newRow, self.grid[row][1:len(self.grid)]...)
		newRow = append(newRow, theMove) //self.grid[row][0])
		self.grid[row] = newRow
	}

	if pos >= 15 && pos < 20 {
		row = row - 15

		var newRow []rune
		oldPeice := self.grid[row][len(self.grid)-1]
		if oldPeice != theMove && oldPeice != '#' {
			// fmt.Printf("what is this thing: %c\n", oldPeice)
			return errors.New("that is an illegal Move")
		}
		newRow = append(newRow, theMove) //self.grid[row][0])
		newRow = append(newRow, self.grid[row][0:len(self.grid)-1]...)
		self.grid[row] = newRow

	}

	// switch the turn
	self.turn = !self.turn
	return nil
}

func (self *board) equals(other *board) bool {
	for i := 0; i < len(self.grid); i++ {
		for j := 0; j < len(self.grid[0]); j++ {
			if self.getPosition(i, j) != other.getPosition(i, j) {
				return false
			}
		}
	}
	return true
}
