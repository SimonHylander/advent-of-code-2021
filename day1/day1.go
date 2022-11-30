package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part")
	flag.Parse()
	fmt.Println(fmt.Sprintf("Running part %d", part))

	if part == 1 {
		answer := part1()
		fmt.Println("Part1 output:", answer)
	} else {
		answer := part2()
		fmt.Println("Part2 output:", answer)
	}
}

func part1() int {
	input, err := readInput()

	if err != nil {
		panic(err)
	}

	increased := 0

	for i := 0; i < len(input); i++ {

		if i > 0 {
			previousMeassurement := input[i-1]

			if input[i] > previousMeassurement {
				increased++
			}
		}
	}

	return increased
}

func part2() int {

	input, err := readInput()

	if err != nil {
		panic(err)
	}

	var increased int

	for i := 0; i < len(input)-3; i++ {
		windowA := input[i] + input[i+1] + input[i+2]
		windowB := input[i+1] + input[i+2] + input[i+3]

		if windowB > windowA {
			increased++
		}
	}

	fmt.Println(fmt.Sprintf("amount: %d", increased))

	return increased
}

func readInput() ([]int, error) {
	// weogo:embed input.txt
	// var f embed.FS
	// data, _ := f.ReadFile("input.txt")

	file, err := os.Open("./day1/input.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fmt.Println(file)

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Not a number")
			return nil, err
		}

		lines = append(lines, number)
	}

	return lines, scanner.Err()
}
