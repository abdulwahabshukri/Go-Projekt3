package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Count int = 2
)

func main() {

	text := os.Args[1]

	text = strings.ReplaceAll(text, "\n", "")

	words := strings.Split(text, " ")

	var wordsWithLines [100][2]string

	var wordCount int
	var lineCount int
	for _, word := range words {
		wordsWithLines[lineCount][wordCount] = word

		wordCount++

		if wordCount == Count {
			lineCount++
			wordCount = 0
		}
	}

	for k := 0; k < 10; k++ {
		fmt.Printf(" _")
	}

	fmt.Println()

	fmt.Printf("/\t\t    \\")

	fmt.Println()

	for i := 0; i < lineCount; i++ {
		line := "     "
		for j := 0; j < Count; j++ {
			if j == Count-1 {
				line += wordsWithLines[i][j] + ""
			} else {
				line += wordsWithLines[i][j] + " "
			}
		}
		fmt.Printf(line + "\n")
	}

	fmt.Printf("\\\t\t    /\n")

	for k := 0; k < 10; k++ {
		fmt.Printf(" -")
	}

	printTux()
}

func printTux() {
	tux := `
		\
		 \
		    .---.
		   |o_o  |
		   |:_/  |
		  //   \  \
		 (|     |  )
		/'\_    /' \
		\___)==(___/`
	fmt.Println(tux)
}
