package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s <regex> [<optional_replacement>]\n", os.Args[0])
		os.Exit(1)
	}

	pattern := os.Args[1]

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) == 3 {
		replacement := os.Args[2]
		
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
