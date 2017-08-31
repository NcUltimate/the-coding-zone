package main

import (
	"fmt"
	"os"
)

func dieOn(err error) {
	if err != nil {
		fmt.Printf("Died:", err.Error())
		os.Exit(1)
	}
}
