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
