package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var deletePattern string

	// Bind -d to deletePattern variable (default 0, description)
	flag.StringVar(&deletePattern, "d", "0", "Line number to delete")

	// Parse the command-line arguments
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	if deletePattern != "0" {
		deletePatternHandler(scanner, deletePattern)

		os.Exit(0)
	}

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage:\n\t%s \n", os.Args[0])
		fmt.Println("\t\t<regex> => Print Matching lines")
		fmt.Println("\t\t<regex> <replacement> => Replace Matching lines")
		fmt.Println("\t\t-d n => remove nth line (counting from 1)")
		os.Exit(1)
	}

	pattern := args[1]

	pattern = strings.ReplaceAll(pattern, "\\n", "\n")

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex: %v\n", err)
		os.Exit(1)
	}

	if len(args) == 3 {
		replacement := args[2]

		replacement = strings.ReplaceAll(replacement, "\\n", "\n")

		fmt.Printf("REPLACEMENT: %s\n", replacement)

		for scanner.Scan() {
			line := scanner.Text()

			if re.MatchString(line) {
				result := re.ReplaceAllString(line, replacement)
				fmt.Printf("%s\n", result)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	} else {
		for scanner.Scan() {
			line := scanner.Text()
			if re.MatchString(line) {
				fmt.Println(line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	}

}

// FIX: works also for negative numbers
func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func deletePatternHandler(scanner *bufio.Scanner, deletePattern string) {
	lineCount := 1

	var lineToDelete int

	rangeRegex := regexp.MustCompile(`^(\d+):(\d+)$`)
	// rangeRegex := regexp.MustCompile(`^([=><])(\d+):(\d+)$`)

	if isInt(deletePattern) {
		lineToDelete, _ = strconv.Atoi(deletePattern)

		for scanner.Scan() {
			line := scanner.Text()

			if lineCount != lineToDelete {
				fmt.Println(line)
			}

			lineCount++
		}
	} else if rangeRegex.MatchString(deletePattern) {
		matches := rangeRegex.FindStringSubmatch(deletePattern)

		start, _ := strconv.Atoi(matches[2])
		end, _ := strconv.Atoi(matches[3])

		for scanner.Scan() {
			line := scanner.Text()

			if !(lineCount >= start && lineCount <= end) {
				fmt.Println(line)
			}

			lineCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}
}
