package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var deleteLine int

	// Bind -d to deleteLine variable (default 0, description)
	flag.IntVar(&deleteLine, "d", 0, "Line number to delete")

	// Parse the command-line arguments
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	if deleteLine != 0 {
		lineCount := 1
		for scanner.Scan() {
			line := scanner.Text()

			if lineCount != deleteLine {
				fmt.Println(line)
			}

			lineCount++
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}

		return
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

	// pattern = strings.ReplaceAll(pattern, ":nl:", "\n")

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex: %v\n", err)
		os.Exit(1)
	}

	if len(args) == 3 {
		replacement := args[2]

		// replacement = strings.ReplaceAll(replacement, ":nl:", "\n")

		for scanner.Scan() {
			line := scanner.Text()

			if re.MatchString(line) {
				result := re.ReplaceAllString(line, replacement)
				fmt.Println(result)
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
