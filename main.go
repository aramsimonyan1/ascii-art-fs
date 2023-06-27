package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Input text or/end banner missing")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(1)
	}

	text := string(os.Args[1])
	if text == "" {
		fmt.Println("Input text is required")
		os.Exit(2)
	}

	if !isASCII(text) {
		fmt.Println("Invalid input text: Non-ASCII characters not allowed")
		os.Exit(3)
	}

	banner := string(os.Args[2])
	// if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
	//	fmt.Println("banner name not correct")
	//}

	var filename string
	switch banner {
	case "shadow":
		filename = "shadow.txt"
	case "standard":
		filename = "standard.txt"
	case "thinkertoy":
		filename = "thinkertoy.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("The banner name is not correct")
		os.Exit(4)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	preLine := []rune(text)
	for m := 0; m < len(preLine); m++ {
		arrayMiddle := "n3wL!Ne"
		if preLine[m] == 92 && preLine[m+1] == 'n' {
			array1 := preLine[0:m]
			array2 := preLine[m+2:]
			s1 := string([]rune(array1))
			s2 := string([]rune(array2))
			text = s1 + arrayMiddle + s2
			preLine = ([]rune(text))
		}
	}
	// split the text into lines if required
	nextStep := string(preLine)
	line := strings.Split(nextStep, "n3wL!Ne")
	// loop to work on lines
	for j := 0; j < len(line); j++ {
		// to make or not make new lines in situations with no other text
		if len(text) < 1 {
			break
		}
		if len(line[j]) < 1 && j == 0 {
			continue
		}
		if len(line[j]) < 1 {
			fmt.Println()
			continue
		}
		word := []rune(line[j])
		// row by row loop
		for k := 1; k < 9; k++ {
			// character by character loop
			for i := 0; i < len(word); i++ {
				m := rune(k)
				// reduce each character value by 32 in ascii table,
				// multiply by the 9 rows each character uses in standard.txt,
				// add the row number
				asciiFetch := ((word[i] - 32) * 9) + m
				fmt.Printf(fileLines[asciiFetch])
			}
			fmt.Println()
		}
	}
}
