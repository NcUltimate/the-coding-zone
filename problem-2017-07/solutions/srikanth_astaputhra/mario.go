package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var width, height int

type Cell struct {
	x int
	y int
}

type Result struct {
	cell      Cell
	coins     int
	end_cells []Cell
}

var coins map[Cell]bool
var blocks map[Cell]bool
var player_start Cell
var resultant_coins map[int]map[int][]Cell
var coin_result map[Cell]map[string]Result

func main() {
	file, err := os.Open("../../sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	coins = make(map[Cell]bool)
	blocks = make(map[Cell]bool)
	coin_result = make(map[Cell]map[string]Result)
	input := bufio.NewScanner(file)

	i := 0
	for input.Scan() {
		indata := input.Text()
		i += 1
		if len(strings.TrimSpace(indata)) == 0 {
			break
		}

		if i == 1 {
			width, _ = strconv.Atoi(strings.Fields(indata)[0])
			height, _ = strconv.Atoi(strings.Fields(indata)[1])
		} else if i == 2 {
			t_x, _ := strconv.Atoi(strings.Fields(indata)[0])
			t_y, _ := strconv.Atoi(strings.Fields(indata)[1])
			player_start = Cell{t_x, t_y}
		} else {
			t_x, _ := strconv.Atoi(strings.Fields(indata)[0])
			t_y, _ := strconv.Atoi(strings.Fields(indata)[1])
			t_cell := Cell{t_x, t_y}
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
	play(player_start, 0)
	resultant_coins = make(map[int]map[int][]Cell, 0)
	for i := player_start.x; i < width; i++ {
		resultant_coins[i] = make(map[int][]Cell)
	}
	fmt.Println()
	find_coins(player_start)
	/*for k, v := range coin_result {
		fmt.Println(k, " ", v["standard_jump"])
	}*/
}

func find_coins(start_from Cell) {
	keys := make([]Cell, 0)
	for i := start_from.x; i < width; i++ {
		for j := start_from.y; j < height; j++ {
			if len(coin_result[Cell{i, j}]) > 0 {
				keys = append(keys, Cell{i, j})
			}
		}
	}
	moves := [4]string{"high_jump", "standard_move", "standard_jump", "long_jump"}
	for _, cell := range keys {
		//result_map := coin_result[cell]
		for _, move := range moves {
			val := coin_result[cell][move]
			if len(resultant_coins[val.cell.x][val.cell.y]) > 0 {
				temp := get_unique_coins(append(resultant_coins[cell.x][cell.y], val.end_cells...))
				if len(temp) > len(resultant_coins[val.cell.x][val.cell.y]) {
					resultant_coins[val.cell.x][val.cell.y] = temp
				} /*else {
				}*/
			} else {
				resultant_coins[val.cell.x][val.cell.y] = get_unique_coins(append(resultant_coins[cell.x][cell.y], val.end_cells...))
			}
		}
	}

	num := 0
	for _, v := range keys {
		if len(resultant_coins[v.x][v.y]) > num {
			num = len(resultant_coins[v.x][v.y])
		}
	}
	fmt.Println(num)
}

func get_unique_coins(cells []Cell) []Cell {
	unique_cells := make([]Cell, 0)
	for _, v := range cells {
		if !in_array(v, unique_cells) {
			unique_cells = append(unique_cells, v)
		}
	}
	return unique_cells
}

func in_array(cell Cell, cell_elements []Cell) bool {
	for _, cell_element := range cell_elements {
		if cell.x == cell_element.x && cell.y == cell_element.y {
			return true
		}
	}
	return false
}

func play(start_cell Cell, coin_count int) (Cell, []int) {
	temp_cell := start_cell
	coins := make([]int, 0)
	for i := start_cell.x; i < width; i++ {
		//diagonal_cell := Cell{temp_cell.x + 1, temp_cell.y + 1}
		temp_cell = Cell{i, temp_cell.y}
		_, temp_cell = falling_move(temp_cell)
		sm_cells, sm_end_cell := standard_move(temp_cell)
		sj_cells, sj_end_cell := standard_jump(temp_cell)
		hj_cells, hj_end_cell := high_jump(temp_cell)
		lj_cells, lj_end_cell := long_jump(temp_cell)
		sm_result := Result{sm_end_cell, len(sm_cells), sm_cells}
		sj_result := Result{sj_end_cell, len(sj_cells), sj_cells}
		hj_result := Result{hj_end_cell, len(hj_cells), hj_cells}
		lj_result := Result{lj_end_cell, len(lj_cells), lj_cells}

		result_collection := make(map[string]Result)
		result_collection["standard_move"] = sm_result
		result_collection["standard_jump"] = sj_result
		result_collection["high_jump"] = hj_result
		result_collection["long_jump"] = lj_result
		coin_result[temp_cell] = result_collection

		if hj_end_cell.y != temp_cell.y && hj_end_cell.y != 0 {
			if _, ok := coin_result[hj_end_cell]; !ok {
				play(hj_end_cell, hj_result.coins+coin_count)
			}
		}
		if lj_end_cell.y != temp_cell.y && lj_end_cell.y != 0 {
			if _, ok := coin_result[lj_end_cell]; !ok {
				play(lj_end_cell, lj_result.coins+coin_count)
			}
		}
	}
	//sort.Sort(keys)
	return temp_cell, coins
}

func falling_move(cell Cell) ([]Cell, Cell) {
	temp_cell := cell
	coins_cell := make([]Cell, 0)
	for temp_cell.y > 0 {
		if blocks[Cell{temp_cell.x, temp_cell.y - 1}] {
			return coins_cell, temp_cell
		}
		temp_cell = Cell{temp_cell.x, temp_cell.y - 1}
		if coins[temp_cell] {
			coins_cell = append(coins_cell, temp_cell)
		}
		if temp_cell.y == 0 {
			return coins_cell, temp_cell
		}
	}
	return coins_cell, temp_cell
}

func standard_move(cell Cell) ([]Cell, Cell) {
	temp_cell := Cell{cell.x + 1, cell.y}
	coins_cell := make([]Cell, 0)
	if blocks[temp_cell] || temp_cell.x == width {
		return coins_cell, cell
	}
	if coins[temp_cell] {
		coins_cell = append(coins_cell, temp_cell)
	}
	temp_coins_cell := make([]Cell, 0)
	temp_coins_cell, temp_cell = falling_move(temp_cell)
	coins_cell = append(coins_cell, temp_coins_cell...)
	return get_unique_coins(coins_cell), temp_cell
}

func standard_jump(cell Cell) ([]Cell, Cell) {
	temp_cell := cell
	coins_cell := make([]Cell, 0)

	y := cell.y + 2

	for y > temp_cell.y {
		if blocks[Cell{temp_cell.x, temp_cell.y + 1}] {
			return coins_cell, cell
		}
		temp_cell = Cell{temp_cell.x, temp_cell.y + 1}
		if coins[temp_cell] {
			coins_cell = append(coins_cell, temp_cell)
		}
	}
	return coins_cell, cell
}

func high_jump(cell Cell) ([]Cell, Cell) {
	temp_cell := cell
	coins_cell := make([]Cell, 0)

	y := cell.y + 2
	for y > temp_cell.y && temp_cell.y < height {
		if blocks[Cell{temp_cell.x, temp_cell.y + 1}] || temp_cell.y+1 == height {
			temp_coins_cell := make([]Cell, 0)
			temp_coins_cell, temp_cell = falling_move(temp_cell)
			coins_cell = append(coins_cell, temp_coins_cell...)
			return coins_cell, temp_cell
		}
		temp_cell = Cell{temp_cell.x, temp_cell.y + 1}
		if coins[temp_cell] {
			coins_cell = append(coins_cell, temp_cell)
		}
	}

	diagonal_cell := Cell{temp_cell.x + 1, temp_cell.y + 1}
	if diagonal_cell.x == width {
		diagonal_cell.x -= 1
	}
	if diagonal_cell.y == height {
		diagonal_cell.y -= 1
	}
	if blocks[diagonal_cell] {
		return coins_cell, cell
	}
	temp_cell = diagonal_cell
	if coins[temp_cell] {
		coins_cell = append(coins_cell, temp_cell)
	}
	temp_coins_cell := make([]Cell, 0)
	temp_coins_cell, temp_cell = falling_move(temp_cell)
	coins_cell = append(coins_cell, temp_coins_cell...)
	return coins_cell, temp_cell
}

func long_jump(cell Cell) ([]Cell, Cell) {
	temp_cell := cell
	coins_cell := make([]Cell, 0)

	diagonal_cell := Cell{temp_cell.x + 1, temp_cell.y + 1}
	if blocks[diagonal_cell] || diagonal_cell.x == width || diagonal_cell.x == height {
		return coins_cell, temp_cell
	}
	temp_cell = diagonal_cell

	if coins[temp_cell] {
		coins_cell = append(coins_cell, temp_cell)
	}

	x := temp_cell.x + 2
	for x > temp_cell.x && temp_cell.x < width {
		if blocks[Cell{temp_cell.x + 1, temp_cell.y}] || temp_cell.x+1 == width {
			temp_coins_cell := make([]Cell, 0)
			temp_coins_cell, temp_cell = falling_move(temp_cell)
			coins_cell = append(coins_cell, temp_coins_cell...)
			return coins_cell, temp_cell
		}
		temp_cell = Cell{temp_cell.x + 1, temp_cell.y}
		if coins[temp_cell] {
			coins_cell = append(coins_cell, temp_cell)
		}
	}

	temp_coins_cell := make([]Cell, 0)
	temp_coins_cell, temp_cell = falling_move(temp_cell)
	coins_cell = append(coins_cell, temp_coins_cell...)
	return coins_cell, temp_cell
}
