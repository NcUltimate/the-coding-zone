package main

import (
	"fmt"
	"os"
	"strconv"
)

func dieOn(err error) {
	if err != nil {
		fmt.Printf("Died:", err.Error())
		os.Exit(1)
	}
}

func atoi(s string) int {
	i, err := strconv.ParseUint(s, 10, 32)
	dieOn(err)
	return int(i)
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	dieOn(err)
	return f
}

func green(s string) string {
	var result string
	result += "\033[0;32m"
	result += s
	result += "\033[0m"
	return result
}

func block(s string) string {
	var result string
	result += "\033[7;34m"
	result += s
	result += "\033[0m"
	return result
}

func blue(s string) string {
	var result string
	result += "\033[0;35m"
	result += s
	result += "\033[0m"
	return result
}
