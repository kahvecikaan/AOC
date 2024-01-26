package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Number to digit mapping
	numMap := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}

	workingDir, _ := os.Getwd()
	file, err := os.Open(workingDir + "/day1/input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstNum := getFirstNumber(line, numMap)
		lastNum := getLastNumber(line, numMap)
		combinedValue := firstNum + lastNum
		intValue, _ := strconv.Atoi(combinedValue)
		sum += intValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file")
		return
	}

	fmt.Println("Sum of calibration values:", sum)
}

func getFirstNumber(line string, numMap map[string]string) string {
	lowestIndex := len(line)
	var value string

	for word, digit := range numMap {
		index := strings.Index(line, word)
		if index != -1 && index < lowestIndex {
			lowestIndex = index
			value = digit
		}

		index = strings.Index(line, digit)
		if index != -1 && index < lowestIndex {
			lowestIndex = index
			value = digit
		}
	}

	return value
}

func getLastNumber(line string, numMap map[string]string) string {
	highestIndex := -1
	var value string

	for word, digit := range numMap {
		index := strings.LastIndex(line, word)
		if index != -1 && index > highestIndex {
			highestIndex = index
			value = digit
		}

		index = strings.LastIndex(line, digit)
		if index != -1 && index > highestIndex {
			highestIndex = index
			value = digit
		}
	}

	return value
}
