package main

import (
	"bufio"
	"errors"
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

var coins map[Cell]bool
var blocks map[Cell]bool
var player_start Cell

func main() {
	//file, err := os.Open("../../medium.txt")
	file, err := os.Open("../../sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	coins = make(map[Cell]bool)
	blocks = make(map[Cell]bool)
	input := bufio.NewScanner(file)

	i := 0
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
			t_row, _ := strconv.Atoi(strings.Fields(indata)[0])
			t_col, _ := strconv.Atoi(strings.Fields(indata)[1])
			player_start = Cell{t_row, t_col}
		} else {
			t_row, _ := strconv.Atoi(strings.Fields(indata)[0])
			t_col, _ := strconv.Atoi(strings.Fields(indata)[1])
			t_cell := Cell{row: t_row, col: t_col}
			type_cell := strings.Fields(indata)[2]

			if type_cell == "C" {
				coins[t_cell] = true
			} else {
				coins[t_cell] = false
			}
			if type_cell == "B" {
				blocks[t_cell] = true
			} else {
				blocks[t_cell] = false
			}
		}
	}
	fmt.Println(get_max_move(player_start.row, player_start.col))
}

func get_max_move(x, y int) int {
	num_coins := 0
	for !moves_finished(Cell{x, y}) {
		sub_s_move := 0

		sub_h_jump := 0
		sub_l_jump := 0
		sub_coins := 0

		s_move_cells, s_move_end_cell := make([]Cell, 0), Cell{x + 1, y}
		if !blocks[Cell{x + 1, y}] {
			s_move_cells, s_move_end_cell, _ = standard_move(x, y)
		}

		s_jump_cells, s_jump_end_cell, _ := standard_jump(x, y)
		h_jump_cells, h_jump_end_cell, _ := high_jump(x, y)
		l_jump_cells, l_jump_end_cell, _ := long_jump(x, y)

		s_move_num := len(s_move_cells)
		s_jump_num := len(s_jump_cells)
		h_jump_num := len(h_jump_cells)
		l_jump_num := len(l_jump_cells)

		moves_coins := make(map[string]int)
		moves_coins["standard_move"] = s_move_num
		moves_coins["standard_jump"] = s_jump_num
		moves_coins["high_jump"] = h_jump_num
		moves_coins["long_jump"] = l_jump_num

		move := decide_move(moves_coins)
		switch move {
		case "standard_move":
			if s_move_num > 0 {
				clear_coins(s_move_cells)
			}
			x, y = s_move_end_cell.row, s_move_end_cell.col
			sub_coins += s_move_num
		case "standard_jump":
			if s_jump_num > 0 {
				clear_coins(s_jump_cells)
			}
			x, y = s_jump_end_cell.row, s_move_end_cell.col
			sub_coins += s_jump_num
		case "high_jump":
			if h_jump_num > 0 {
				clear_coins(h_jump_cells)
			}
			x, y = h_jump_end_cell.row, s_move_end_cell.col
			sub_coins += h_jump_num
		case "long_jump":
			if l_jump_num > 0 {
				clear_coins(l_jump_cells)
			}
			x, y = l_jump_end_cell.row, s_move_end_cell.col
			sub_coins += l_jump_num
		}
		if s_move_end_cell.row != x && s_move_end_cell.col != y && !blocks[Cell{x + 1, y}] {
			sub_s_move = get_max_move(s_move_end_cell.row, s_move_end_cell.col)
		}
		if h_jump_end_cell.row != x && h_jump_end_cell.col != y {
			sub_h_jump = get_max_move(h_jump_end_cell.row, h_jump_end_cell.col)
		}
		if l_jump_end_cell.row != x && l_jump_end_cell.col != y && !blocks[Cell{l_jump_end_cell.row + 1, l_jump_end_cell.col + 1}] {
			sub_l_jump = get_max_move(l_jump_end_cell.row, l_jump_end_cell.col)
		}

		largest_coins := 0
		if largest_coins < sub_coins {
			largest_coins = sub_coins
		}
		if largest_coins < sub_s_move {
			largest_coins = sub_s_move
		}
		if largest_coins < sub_h_jump {
			largest_coins = sub_h_jump
		}
		if largest_coins < sub_l_jump {
			largest_coins = sub_l_jump
		}
		num_coins += largest_coins
	}
	if num_coins > 0 {
		fmt.Println("----", num_coins)
	}
	return num_coins
}

func moves_finished(cell Cell) bool {
	below_cell := Cell{cell.row - 1, cell.col}
	forward_cell := Cell{cell.row, cell.col + 1}
	upward_cell := Cell{cell.row + 1, cell.col}
	above_upward_cell := Cell{cell.row + 2, cell.col}
	if cell.row+1 >= rows {
		return true
	}
	if coins[upward_cell] || coins[above_upward_cell] {
		return false
	}
	if blocks[below_cell] || blocks[forward_cell] || blocks[upward_cell] {
		return false
	}
	return false
}

func clear_coins(cells []Cell) {
	for _, v := range cells {
		coins[v] = false
	}
}

func decide_move(moves map[string]int) string {
	return_string := "standard_move"
	if moves[return_string] < moves["standard_jump"] {
		return_string = "standard_jump"
	}
	if moves[return_string] < moves["high_jump"] {
		return_string = "high_jump"
	}
	if moves[return_string] < moves["long_jump"] {
		return_string = "long_jump"
	}

	return return_string
}

func falling_move(jump_cell Cell, return_cell []Cell) ([]Cell, Cell, error) {
	for jump_cell.col >= 0 {
		if coins[jump_cell] {
			return_cell = append(return_cell, jump_cell)
		}
		if blocks[jump_cell] {
			return return_cell, Cell{jump_cell.row, jump_cell.col + 1}, errors.New("Encountered block high_jump")
		}
		if jump_cell.col == 0 {
			return return_cell, jump_cell, nil
		}
		jump_cell = Cell{row: jump_cell.row, col: jump_cell.col - 1}
	}
	if jump_cell.col > 0 {
		return return_cell, jump_cell, nil
	}
	return return_cell, Cell{jump_cell.row, 0}, nil
}

func standard_move(x, y int) ([]Cell, Cell, error) {
	fmt.Println("standard_move ", x, y)
	s_move_cell := Cell{x + 1, y}
	return_cell := make([]Cell, 0)
	if x == rows {
		return return_cell, Cell{x, y}, nil
	}
	if coins[s_move_cell] {
		return_cell = append(return_cell, s_move_cell)
	}
	if blocks[s_move_cell] == true {
		return return_cell, Cell{x, y}, errors.New("Encountered block stanndard_move")
	}
	falling_return_cell, jump_cell, falling_err := falling_move(Cell{s_move_cell.row, s_move_cell.col - 1}, make([]Cell, 0))
	if falling_err == nil {
		return_cell = append(return_cell, falling_return_cell...)
		return return_cell, jump_cell, nil
	}

	return return_cell, Cell{x + 1, 0}, nil
}

func standard_jump(x, y int) ([]Cell, Cell, error) {
	fmt.Println("standard_jump ", x, y)
	return_cell, end_cell, err := s_jump(x, y)
	if err == nil {
		jump_cell := Cell{end_cell.row, end_cell.col - 1}
		falling_return_cell, jump_cell, falling_err := falling_move(Cell{end_cell.row, end_cell.col - 1}, make([]Cell, 0))
		if falling_err == nil {
			return_cell = append(return_cell, falling_return_cell...)
			return return_cell, jump_cell, nil
		}
	} else {
		return return_cell, end_cell, err
	}
	return return_cell, Cell{x, y}, nil
}

func s_jump(x, y int) ([]Cell, Cell, error) {
	s_jump_cell := Cell{x, y + 1}
	return_cell := make([]Cell, 0)
	for x <= s_jump_cell.row && y+2 <= s_jump_cell.col {
		if coins[s_jump_cell] {
			return_cell = append(return_cell, s_jump_cell)
		}
		if blocks[s_jump_cell] == true {
			return return_cell, Cell{s_jump_cell.row, s_jump_cell.col - 1}, errors.New("Encountered block standard_jump")
		}
		s_jump_cell = Cell{s_jump_cell.row, s_jump_cell.col + 1}
	}
	return return_cell, s_jump_cell, nil
}

func high_jump(x, y int) ([]Cell, Cell, error) {
	fmt.Println("high_jump ", x, y)
	h_jump_cell := Cell{x + 1, y + 3}
	return_cell, end_cell, err := s_jump(x, y)
	if err == nil {
		if coins[h_jump_cell] {
			return_cell = append(return_cell, h_jump_cell)
		}
		if blocks[h_jump_cell] {
			fmt.Println("Encountered block high_jump")
			h_jump_cell = Cell{row: h_jump_cell.row, col: h_jump_cell.col - 1}
			for h_jump_cell.col >= 0 {
				if coins[h_jump_cell] {
					return_cell = append(return_cell, h_jump_cell)
				}
				if blocks[h_jump_cell] {
					return return_cell, Cell{h_jump_cell.row, h_jump_cell.col + 1}, errors.New("Encountered block high_jump")
				}
				if h_jump_cell.col == 0 {
					return return_cell, h_jump_cell, nil
				}
				h_jump_cell = Cell{h_jump_cell.row, h_jump_cell.col - 1}
			}
			return return_cell, h_jump_cell, errors.New("Encountered block high_jump")
		}
	} else {
		return return_cell, end_cell, err
	}
	h_jump_cell = Cell{row: h_jump_cell.row, col: h_jump_cell.col - 1}
	for h_jump_cell.col >= 0 {
		if coins[h_jump_cell] {
			return_cell = append(return_cell, h_jump_cell)
		}
		if blocks[h_jump_cell] {
			return return_cell, Cell{h_jump_cell.row, h_jump_cell.col + 1}, errors.New("Encountered block high_jump")
		}
		if h_jump_cell.col == 0 {
			return return_cell, h_jump_cell, nil
		}
		h_jump_cell = Cell{h_jump_cell.row, h_jump_cell.col - 1}
	}
	return return_cell, h_jump_cell, nil
}

func long_jump(x, y int) ([]Cell, Cell, error) {
	fmt.Println("long_jump ", x, y)
	l_jump_cell := Cell{x + 1, y + 1}
	return_cell := make([]Cell, 0)
	if coins[l_jump_cell] {
		return_cell = append(return_cell, l_jump_cell)
	}
	if blocks[l_jump_cell] {
		return return_cell, Cell{l_jump_cell.row - 1, l_jump_cell.col - 1}, errors.New("Encountered block long_jump")
	}
	x = x + 1
	y = y + 1
	for l_jump_cell.row >= x+2 && l_jump_cell.col >= y {
		if coins[l_jump_cell] {
			return_cell = append(return_cell, l_jump_cell)
		}
		if blocks[l_jump_cell] {
			return return_cell, Cell{l_jump_cell.row - 1, l_jump_cell.col}, errors.New("Encountered block long_jump")
		}
		l_jump_cell = Cell{l_jump_cell.row + 1, l_jump_cell.col}
	}
	l_jump_cell = Cell{row: l_jump_cell.row, col: l_jump_cell.col - 1}
	for l_jump_cell.col >= 0 {
		if coins[l_jump_cell] {
			return_cell = append(return_cell, l_jump_cell)
		}
		if blocks[l_jump_cell] {
			return return_cell, Cell{l_jump_cell.row, l_jump_cell.col + 1}, errors.New("Encountered block long_jump")
		}
		if l_jump_cell.col == 0 {
			return return_cell, l_jump_cell, nil
		}
		l_jump_cell = Cell{row: l_jump_cell.row, col: l_jump_cell.col - 1}
	}

	return return_cell, l_jump_cell, nil
}
