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

var coins map[Cell]bool
var blocks map[Cell]bool
var player_start Cell

func main() {
	file, err := os.Open("../../medium.txt")
	//file, err := os.Open("../../sample.txt")
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
			rows, _ = strconv.Atoi(strings.Fields(indata)[1])
			cols, _ = strconv.Atoi(strings.Fields(indata)[0])
		} else if i == 2 {
			t_row, _ := strconv.Atoi(strings.Fields(indata)[1])
			t_col, _ := strconv.Atoi(strings.Fields(indata)[0])
			player_start = Cell{t_row, t_col}
		} else {
			t_row, _ := strconv.Atoi(strings.Fields(indata)[1])
			t_col, _ := strconv.Atoi(strings.Fields(indata)[0])
			t_cell := Cell{t_row, t_col}
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
	end_cell, all_cell_count := play(player_start, 0)
	//end_cell, all_cell_count := play(Cell{2, 7}, 0)
	fmt.Println(end_cell)
	fmt.Println(all_cell_count)
}

func play(start_cell Cell, coin_count int) (Cell, []int) {
	temp_cell := start_cell
	done_standard_jump := false
	num_coins := coin_count
	if coins[start_cell] {
		num_coins += 1
		fmt.Println("Added ", start_cell, " ", num_coins)
		coins[start_cell] = false
	}
	sub_coins := 0
	end_cell := temp_cell
	cell_count := make([]int, 0)
	cells := make([]Cell, 0)
	collected_cells := make([]Cell, 0)
	i := 0
	for !moves_finished(temp_cell) {
		sub_coins = 0
		//fmt.Println(temp_cell, " ", can_high_jump(temp_cell))
		if can_high_jump(temp_cell) {
			_, temp_end_cell := high_jump(temp_cell, collected_cells)
			if level_changed(temp_cell, temp_end_cell) {
				_, sub_cell_count := play(temp_end_cell, num_coins)
				//fmt.Println("---", temp_end_cell, " ", sub_cell_count)
				cell_count = append(cell_count, sub_cell_count...)
			}
		}
		if has_atleast_two_coins_above(temp_cell) && !done_standard_jump {
			cells, end_cell = standard_jump(temp_cell, collected_cells)
			sub_coins = len(get_unique_coins(cells))

			done_standard_jump = true
		} else if should_high_jump(temp_cell) {
			cells, end_cell = high_jump(temp_cell, collected_cells)
			sub_coins = len(get_unique_coins(cells))
			done_standard_jump = true
		} else if should_long_jump(temp_cell) {
			cells, end_cell = long_jump(temp_cell, collected_cells)
			sub_coins = len(get_unique_coins(cells))
			done_standard_jump = false
		} else {
			cells, end_cell = standard_move(temp_cell, collected_cells)
			sub_coins = len(get_unique_coins(cells))
			done_standard_jump = false
		}
		collected_cells = append(collected_cells, get_unique_coins(cells)...)
		collected_cells = get_unique_coins(collected_cells)
		num_coins += sub_coins
		temp_cell = end_cell
		i += 1
		/*if i == 30 {
			os.Exit(1)
		}*/
	}

	fmt.Println(collected_cells)
	cell_count = append(cell_count, num_coins)
	return temp_cell, cell_count
}

func can_high_jump(cell Cell) bool {
	max_cell := Cell{cell.row + 3, cell.col + 1}
	//fmt.Println("IS BLOCK ", " ", max_cell, " ", !is_block(max_cell))
	if !is_block(max_cell) {
		//if (is_block(Cell{cell.row, cell.col + 1}) || is_block(Cell{cell.row + 1, cell.col + 1})) && !is_block(Cell{cell.row + 2, cell.col + 1}) {
		if is_block(Cell{cell.row + 1, cell.col + 1}) {
			return true
		}
		if (is_block(Cell{cell.row, cell.col + 1}) || is_block(Cell{cell.row + 1, cell.col + 1})) {
			return true
		}
	}
	return false
}

/*
func can_long_jump(cell Cell) bool {
	max_cell := Cell{cell.row + 1, cell.col + 3}
	diagonal_cell := Cell{cell.row + 1, cell.col + 1}
	if is_block(Cell{cell.row, cell.col + 1}) {
	}
}*/

func should_long_jump(cell Cell) bool {
	//max_cell := Cell{cell.row + 1, cell.col + 3}
	diagonal_cell := Cell{cell.row + 1, cell.col + 1}
	if is_block(Cell{cell.row, cell.col + 1}) {
		if !is_block(diagonal_cell) {
			return true
		}
	}
	return false
}

func should_high_jump(cell Cell) bool {
	max_cell := Cell{cell.row + 3, cell.col + 1}
	if !is_block(Cell{cell.row + 1, cell.col}) && !is_block(Cell{cell.row + 2, cell.col}) {
		if coins[max_cell] {
			return true
		}
	}
	if is_block(Cell{cell.row + 2, cell.col + 1}) {
		return true
	}
	return false
}

func level_changed(start_cell Cell, end_cell Cell) bool {
	if start_cell.row != end_cell.row {
		return true
	}
	return false
}

func can_change_level(cell Cell) bool {
	if blocks[Cell{cell.row + 1, cell.col + 2}] {
		return true
	}
	if blocks[Cell{cell.row + 3, cell.col}] {
		return true
	}
	return false
}

func inside_grid(cell Cell) bool {
	if (cell.row > 0 && cell.row < rows) && (cell.col >= 0 && cell.col <= cols) {
		return true
	}
	return false
}

func falling_move(return_cell []Cell, jump_cell Cell, collected_cells []Cell) ([]Cell, Cell) {
	for jump_cell.row >= 0 {
		if coins[jump_cell] && !in_array(jump_cell, collected_cells) {
			return_cell = append(return_cell, jump_cell)
			collected_cells = append(collected_cells, jump_cell)
		}
		jump_cell = Cell{jump_cell.row - 1, jump_cell.col}
		if blocks[jump_cell] {
			return return_cell, Cell{jump_cell.row + 1, jump_cell.col}
		}
	}
	if jump_cell.row < 0 {
		jump_cell = Cell{0, jump_cell.col}
	}
	return return_cell, jump_cell
}

func standard_move(cell Cell, collected_cells []Cell) ([]Cell, Cell) {
	return_cell := make([]Cell, 0)
	if can_move_horizontally(cell) {
		cell = Cell{cell.row, cell.col + 1}
		if coins[cell] && !in_array(cell, collected_cells) {
			return_cell = append(return_cell, cell)
			collected_cells = append(collected_cells, cell)
		}
	} else {
		return return_cell, cell
	}
	falling_return_cell, falling_jump_cell := falling_move(return_cell, cell, collected_cells)
	return_cell = append(return_cell, falling_return_cell...)
	return return_cell, falling_jump_cell
}

func standard_jump(cell Cell, collected_cells []Cell) ([]Cell, Cell) {
	return_cell := make([]Cell, 0)
	temp_cell := cell
	if can_standard_jump(temp_cell) {
		for i := 0; i < 2; i++ {
			temp_cell = Cell{temp_cell.row + 1, temp_cell.col}
			if coins[temp_cell] && !in_array(temp_cell, collected_cells) {
				return_cell = append(return_cell, temp_cell)
				collected_cells = append(collected_cells, temp_cell)
			}
		}
	} else if has_one_coin_above(temp_cell) {
		temp_cell = Cell{temp_cell.row + 1, temp_cell.col}
		if coins[temp_cell] && !in_array(temp_cell, collected_cells) {
			return_cell = append(return_cell, temp_cell)
			collected_cells = append(collected_cells, temp_cell)
		}
	} else {
		return return_cell, cell
	}

	falling_return_cell, falling_jump_cell := falling_move(return_cell, cell, collected_cells)
	return_cell = append(return_cell, falling_return_cell...)
	return return_cell, falling_jump_cell
}

func high_jump(cell Cell, collected_cells []Cell) ([]Cell, Cell) {
	temp_cell := cell
	return_cell := make([]Cell, 0)
	for i := 0; i < 2 && temp_cell.row < rows; i++ {
		temp_cell = Cell{temp_cell.row + 1, temp_cell.col}
		if coins[temp_cell] && !in_array(temp_cell, collected_cells) {
			return_cell = append(return_cell, temp_cell)
			collected_cells = append(collected_cells, temp_cell)
		}
		if blocks[temp_cell] {
			return return_cell, Cell{temp_cell.row - 1, temp_cell.col}
		}
	}
	diagonal_cell := Cell{temp_cell.row + 1, temp_cell.col + 1}
	//fmt.Println("Diagonal: ", diagonal_cell)
	if coins[diagonal_cell] && !in_array(diagonal_cell, collected_cells) {
		temp_cell = diagonal_cell
		return_cell = append(return_cell, temp_cell)
		collected_cells = append(collected_cells, temp_cell)
	}
	if blocks[diagonal_cell] {
		return return_cell, Cell{temp_cell.row - 1, temp_cell.col}
	} else {
		temp_cell = diagonal_cell
	}
	falling_return_cell, falling_jump_cell := falling_move(return_cell, temp_cell, collected_cells)
	return_cell = append(return_cell, falling_return_cell...)
	return return_cell, falling_jump_cell
}

func long_jump(cell Cell, collected_cells []Cell) ([]Cell, Cell) {
	diagonal_cell := Cell{cell.row + 1, cell.col + 1}
	return_cell := make([]Cell, 0)
	temp_cell := cell

	if coins[diagonal_cell] && !in_array(diagonal_cell, collected_cells) {
		temp_cell = diagonal_cell
		return_cell = append(return_cell, temp_cell)
		collected_cells = append(collected_cells, temp_cell)
	}
	if blocks[diagonal_cell] {
		return return_cell, cell
	}
	end_cell := diagonal_cell
	for diagonal_cell.col+2 >= end_cell.col && can_move_horizontally(end_cell) {
		end_cell = Cell{end_cell.row, end_cell.col + 1}
		if coins[end_cell] && !in_array(end_cell, collected_cells) {
			return_cell = append(return_cell, end_cell)
			collected_cells = append(collected_cells, temp_cell)
		}
		if blocks[end_cell] {
			end_cell = Cell{end_cell.row, end_cell.col - 1}
			break
		}
	}
	falling_return_cell, falling_jump_cell := falling_move(return_cell, Cell{end_cell.row, end_cell.col}, collected_cells)
	return_cell = append(return_cell, falling_return_cell...)
	return return_cell, falling_jump_cell
}

func can_standard_jump(cell Cell) bool {
	if (is_block(Cell{cell.row, cell.col + 1})) {
		return false
	}
	if has_atleast_two_coins_above(cell) {
		return true
	}
	return false
}

func has_one_coin_above(cell Cell) bool {
	if (coins[Cell{cell.row + 1, cell.col}]) {
		return true
	}
	return false
}

func has_atleast_two_coins_above(cell Cell) bool {
	if (has_one_coin_above(cell) || coins[Cell{cell.row + 2, cell.col}]) {
		return true
	}
	return false
}

func can_move_up(cell Cell) bool {
	//fmt.Println("V: ", cell, " ", Cell{cell.row + 1, cell.col}, " ", !can_high_jump(cell), " ", !coins[Cell{cell.row + 1, cell.col}])
	if cell.row >= rows {
		return false
	}
	if blocks[Cell{cell.row + 1, cell.col}] {
		return false
	}
	if !coins[Cell{cell.row + 1, cell.col}] {
		return false
	}
	return true
}

func can_move_horizontally(cell Cell) bool {
	if cell.col >= cols {
		return false
	}
	if is_block(Cell{cell.row, cell.col + 1}) {
		return false
	}
	return true
}

func moves_finished(cell Cell) bool {
	//fmt.Println("Horizontal ", can_move_horizontally(cell), " Up: ", can_move_up(cell), " Cell: ", cell)
	if can_move_horizontally(cell) || can_move_up(cell) {
		return false
	}
	return true
}

func get_unique_coins(cells []Cell) []Cell {
	unique_cells := make([]Cell, 0)
	for _, v := range cells {
		if !in_array(v, unique_cells) {
			////fmt.Println("Adding ", v)
			unique_cells = append(unique_cells, v)
		}
	}
	return unique_cells
}

func in_array(cell Cell, cell_elements []Cell) bool {
	for _, cell_element := range cell_elements {
		if cell.row == cell_element.row && cell.col == cell_element.col {
			return true
		}
	}
	return false
}

func clear_coins(cells []Cell) int {
	count := 0
	for _, v := range cells {
		if coins[v] {
			//fmt.Println("Adding: ", v)
			coins[v] = false
			count += 1
		}
	}
	//fmt.Println("Count: ", count)
	return count
}

func is_block(cell Cell) bool {
	return blocks[cell]
}
