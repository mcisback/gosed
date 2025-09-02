package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var deletePattern string
	var path string

	boldMode := false

	bold := color.New(color.Bold).SprintFunc()

	// Bind -d to deletePattern variable (default 0, description)
	flag.StringVar(&deletePattern, "d", "0", "Line number to delete")
	flag.StringVar(&path, "f", "", "File to open (Defaults to stdin)")

	// Parse the command-line arguments
	flag.Parse()

	args := flag.Args()

	//fmt.Println("ARGS: ", args)

	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\t%s \n", os.Args[0])
		fmt.Println("\t\t<regex> => Print Matching lines")
		fmt.Println("\t\t<regex> <replacement> => Replace Matching lines")
		fmt.Println("\t\t-d n => remove nth line (counting from 1)")
		fmt.Println("\t\t-f path => use file or path instead of STDIN")
		os.Exit(1)
	}

	inputFile := os.Stdin

	if path != "" {
		pathExists, pathMode, _ := PathExists(path)

		fmt.Println("PATH: ", path)

		if !pathExists {
			fmt.Fprintf(os.Stderr, "[!] Error, path %s doesn't exists", path)

			os.Exit(1)
		}

		if pathMode.IsDir() {
			fmt.Println("Directory Mode Not Implemented YET")

			os.Exit(1)
		} else if pathMode.IsRegular() {
			file, err := os.Open(path)

			if err != nil {
				fmt.Println("Error opening file:", err)

				os.Exit(1)
			}

			inputFile = file
		}

		defer inputFile.Close()
	}

	scanner := bufio.NewScanner(inputFile)

	// Delete mode
	if deletePattern != "0" {
		deletePatternHandler(scanner, deletePattern)

		os.Exit(0)
	}

	pattern := args[0]

	if strings.HasSuffix(pattern, "/d") {
		deletePatternHandler(scanner, pattern)

		os.Exit(0)
	}

	if strings.HasSuffix(pattern, "/b") {
		boldMode = true

		pattern = pattern[:len(pattern)-2]
	}

	pattern = strings.ReplaceAll(pattern, "\\n", "\n")

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex: %v\n", err)

		os.Exit(1)
	}

	if len(args) >= 2 {
		// Replace
		replacement := args[1]

		replacement = strings.ReplaceAll(replacement, "\\n", "\n")

		// fmt.Printf("REPLACEMENT: %s\n", replacement)

		for scanner.Scan() {
			line := scanner.Text()

			if re.MatchString(line) {
				line = re.ReplaceAllString(line, replacement)

				// if boldMode {
				line = bold(line)
				// }
			}

			fmt.Println(line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Print Matching
		for scanner.Scan() {
			line := scanner.Text()

			if boldMode {
				if re.MatchString(line) {
					fmt.Println(bold(line))
				} else {
					fmt.Println(line)
				}

				continue
			}

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

func PathExists(path string) (bool, os.FileMode, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, 0, nil
		}
		return false, 0, err
	}
	return true, info.Mode(), nil
}

// FIX: works also for negative numbers
func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func deletePatternHandler(scanner *bufio.Scanner, deletePattern string) {
	// fmt.Println("deletePatternHandler", deletePattern)

	lineCount := 1

	var lineToDelete int

	rangeRegex := regexp.MustCompile(`^(\d+):(\d+)$`)
	deleteMatchingLinesRegex := regexp.MustCompile(`^(.+)/d$`)
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
	} else if deleteMatchingLinesRegex.MatchString(deletePattern) {
		// fmt.Println("deleteMatchingLinesRegex")

		matches := deleteMatchingLinesRegex.FindStringSubmatch(deletePattern)

		pattern := matches[1]

		re := regexp.MustCompile(pattern)

		for scanner.Scan() {
			line := scanner.Text()

			if !re.MatchString(line) {
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
