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
	numZeroes2 := 0
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
			numZeroes2 += (currentPosition + distance) / MOD
			currentPosition = (currentPosition + distance) % MOD

		case 'L':
			numZeroes2 += (distance / 100)
			rest := distance % 100

			if currentPosition > 0 && rest >= currentPosition {
				numZeroes2++
			}

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

	fmt.Printf("Part 1 : Final Password (Total Zeroes): %d\n", numZeroes)
	fmt.Printf("Part 2 : Final Password (Total Zeroes): %d\n", numZeroes2)
}
