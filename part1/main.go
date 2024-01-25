package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("/Users/kaankahveci/GolandProjects/adventOfCodeDay1/input.txt")
	// print the current working directory
	dir, _ := os.Getwd()
	fmt.Println(dir)
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += getCalibrationValue(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file")
		return
	}

	fmt.Println("Total sum of calibration values: ", sum)
}

func getCalibrationValue(s string) int {
	firstDigit := -1
	lastDigit := -1

	// Find the first digit
	for _, r := range s {
		if unicode.IsDigit(r) {
			firstDigit = int(r - '0')
			break
		}
	}

	// Find the last digit
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			lastDigit = int(s[i] - '0')
			break
		}
	}

	if firstDigit != -1 || lastDigit != -1 {
		return firstDigit*10 + lastDigit
	}

	return 0
}
