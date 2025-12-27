package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	currentPosition := 50
	numZeroes := 0
	const MOD = 100

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		direction := line[0]
		distanceStr := line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			continue
		}

		switch direction {
		case 'R':
			currentPosition = (currentPosition + distance) % MOD

		case 'L':
			currentPosition = ((currentPosition-distance)%MOD + MOD) % MOD
		}

		if currentPosition == 0 {
			numZeroes++
		}

		fmt.Printf("Line: %s | Count: %d \n", line, currentPosition)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading line: %v\n", err)
		return
	}
	fmt.Printf("Final Password (Total Zeroes): %d\n", numZeroes)
}
