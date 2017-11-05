package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var instructions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = strings.Split(line, ",")
	}

	var totalDistance int64 = 0
	var distance int64 = 0
	var currentDirection int = 0
	var currentPositionX int = 0
	var currentPositionY int = 0
	var direction string;

	for _, element := range instructions {
		trimmedElement := strings.Trim(element, " ")
		// [starting index, ending index]
		direction = trimmedElement[0:1]
		distance, err = strconv.ParseInt(trimmedElement[1:], 10, 64)
		
		if direction == string("L") {
			currentDirection -= 90;
		} else {
			currentDirection += 90;
		}

		if currentDirection % 360 == 0 {
			currentDirection = 0
		}

		switch currentDirection {
		case 0:
			currentPositionY += int(distance)
		case 90, -270:
			currentPositionX += int(distance)
		case -90, 270:
			currentPositionX -= int(distance)
		case -180, 180:
			currentPositionY -= int(distance)
		}
	}

	totalDistance = abs(currentPositionX) + abs(currentPositionY)

	fmt.Println(totalDistance)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func abs(n int) int64 {
	if (n < 0) {
		return int64(-1 * n)
	} else {
		return int64(n)
	}
}
