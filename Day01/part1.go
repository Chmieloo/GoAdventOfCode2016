package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	for _, element := range instructions {
		fmt.Println(element)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
