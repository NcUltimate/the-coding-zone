package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rows, cols int

type Cell struct {
	row int
	col int
}

type Deflector struct {
	cell      Cell
	deflector string
}

var deflectors map[Cell]string
var target Cell

func main() {
	file, err := os.Open("../../test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	i := 0
	deflectors = make(map[Cell]string)
	input := bufio.NewScanner(file)
	for input.Scan() {
		indata := input.Text()
		i += 1
		if len(strings.TrimSpace(indata)) == 0 {
			break
		}
		if i == 1 {
			rows, _ = strconv.Atoi(strings.Fields(indata)[0])
			cols, _ = strconv.Atoi(strings.Fields(indata)[1])
		} else if i == 2 {
			target_row, _ := strconv.Atoi(strings.Fields(indata)[0])
			target_col, _ := strconv.Atoi(strings.Fields(indata)[1])
			target = Cell{row: target_row, col: target_col}
		} else {
			t_row, _ := strconv.Atoi(strings.Fields(indata)[0])
			t_col, _ := strconv.Atoi(strings.Fields(indata)[1])
			t_cell := Cell{row: t_row, col: t_col}
			deflectors[t_cell] = strings.Fields(indata)[2]
		}
	}

	// start from north
	for i, j := 1, 1; j <= cols; j++ {
		if traverse_from_north_forwards(i, j) {
			fmt.Println("N ", i)
		}
	}

	// start from south
	for i, j := rows, 1; j <= cols; j++ {
		if traverse_from_south_forwards(i, j) {
			fmt.Println("S ", i)
		}
	}

	// start from east
	for i, j := 1, cols; i <= rows; i++ {
		if traverse_from_east_forwards(i, j) {
			fmt.Println("E ", i)
		}
	}

	// start from west
	for i, j := 1, 1; i <= rows; i++ {
		if traverse_from_west_forwards(i, j) {
			fmt.Println("W ", i)
		}
	}

}

func traverse_from_south_forwards(row int, col int) bool {
	for a, b := row, col; a >= 1; a-- {
		def, ok := deflectors[Cell{row: a, col: b}]
		if target.row == a && target.col == b {
			return true
		}

		if ok {
			if def == "/" {
				return traverse_from_west_forwards(a, b+1)
			} else {
				return traverse_from_east_forwards(a, b-1)
			}
		}
	}
	return false
}

func traverse_from_west_forwards(row int, col int) bool {
	for a, b := row, col; b <= cols; b++ {
		def, ok := deflectors[Cell{row: a, col: b}]
		if target.row == a && target.col == b {
			return true
		}
		if ok {
			if def == "/" {
				return traverse_from_south_forwards(a-1, b)
			} else {
				return traverse_from_north_forwards(a+1, b)
			}
		}
	}
	return false
}

func traverse_from_north_forwards(row int, col int) bool {
	for a, b := row, col; a <= rows; a++ {
		def, ok := deflectors[Cell{row: a, col: b}]
		if target.row == a && target.col == b {
			return true
		}
		if ok {
			if def == "/" {
				return traverse_from_east_forwards(a, b-1)
			} else {
				return traverse_from_west_forwards(a, b+1)
			}
		}
	}
	return false
}

func traverse_from_east_forwards(row int, col int) bool {
	for a, b := row, col; b >= 1; b-- {
		def, ok := deflectors[Cell{row: a, col: b}]
		if target.row == a && target.col == b {
			return true
		}
		if ok {
			if def == "/" {
				return traverse_from_north_forwards(a+1, b)
			} else {
				return traverse_from_south_forwards(a-1, b)
			}
		}
	}
	return false
}
