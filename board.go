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

func (self *board) makeMove(pos int) error {
	theMove := self.whoseTurn()
	// fmt.Printf("move to?: %v\n", pos)
	// fmt.Printf("wats the move %c\n", self.whoseTurn())

	var row int
	if pos >= 0 && pos < 5 || isBottom(pos) {
		if pos >= 0 && pos < 5 {
			row = 0

		} else if isBottom(pos) {
			pos = pos - 5
			row = 4

		}
		fmt.Printf("where are we moving? (%v,%v)\n", row, pos)

		oldPeice := self.grid[row][pos]
		if oldPeice != theMove && oldPeice != '#' {
			return errors.New("that is an illegal Move")
		}
		for ; row < 5; row++ {
			if isBottom(pos) {
				row = 4 - row
			}
			restOfRow := self.grid[row][pos+1 : len(self.grid[0])]
			oldPeice = self.grid[row][pos]
			startOfRow := self.grid[row][0:pos]
			var newRow []rune
			newRow = append(newRow, startOfRow...)
			if (pos >= 0 && pos < 5 && row == 4) || isBottom(pos) && row == 0 {
				newRow = append(newRow, theMove)
			} else {
				var newPeice rune
				if isBottom(pos) {
					newPeice = self.grid[row-1][pos]
				} else {
					newPeice = self.grid[row+1][pos]
				}
				newRow = append(newRow, newPeice)
			}
			newRow = append(newRow, restOfRow...)
			self.grid[row] = newRow
		}

	}

	if pos >= 10 && pos < 15 {
		fmt.Printf("do we get here?\n")
		pos = pos - 10

		// fmt.Printf("%v\n", self.grid[thing])
		var newRow []rune
		oldPeice := self.grid[pos][0]
		if oldPeice != theMove && oldPeice != '#' {
			// fmt.Printf("what is this thing: %c\n", oldPeice)
			return errors.New("that is an illegal Move")
		}
		newRow = append(newRow, self.grid[pos][1:len(self.grid)]...)
		newRow = append(newRow, theMove) //self.grid[pos][0])
		self.grid[pos] = newRow

	}

	if pos >= 15 && pos < 20 {
		pos = pos - 15

		var newRow []rune
		oldPeice := self.grid[pos][len(self.grid)-1]
		if oldPeice != theMove && oldPeice != '#' {
			// fmt.Printf("what is this thing: %c\n", oldPeice)
			return errors.New("that is an illegal Move")
		}
		newRow = append(newRow, theMove) //self.grid[pos][0])
		newRow = append(newRow, self.grid[pos][0:len(self.grid)-1]...)
		self.grid[pos] = newRow

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
