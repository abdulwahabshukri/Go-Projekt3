package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	COUNT_PER_LINE int = 2
)

func main() {

	fmt.Printf("pinguinsay ")

	in := bufio.NewReader(os.Stdin)

	text, _ := in.ReadString('\n')

	text = strings.ReplaceAll(text, "\n", "")

	words := strings.Split(text, " ")

	var wordsWithLines [100][50]string

	var wordCount int
	var lineCount int
	for _, word := range words {
		wordsWithLines[lineCount][wordCount] = word

		wordCount++

		if wordCount == COUNT_PER_LINE {
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
		for j := 0; j < COUNT_PER_LINE; j++ {
			if j == COUNT_PER_LINE-1 {
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
