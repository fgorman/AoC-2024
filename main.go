package main

import (
	"fmt"
	"flag"
	"time"
)

func main() {
	// Get current day
	_, _, day := time.Now().Date()

	// Set CLI flags and parse
	aocDay := flag.Int("day", day, "Day for solution. Defaults to the current AoC day.")
	flag.Parse()
}