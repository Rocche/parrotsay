// Thank you Flavio for the amazing tutorial "Building a CLI command with Go: cowsay"
// https://flaviocopes.com/go-tutorial-cowsay/
// Please go check it out!
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"parrotsay/balloon"
	"parrotsay/characters"
)

func printUsage() {
	fmt.Println("Usage: fortune | parrotsay")
}

// reads the input to the program and returns the list of input lines
func readUserInput() ([]string, error) {
	info, _ := os.Stdin.Stat()

	// make sure there is actually an input coming from stdin
	if info.Mode()&os.ModeCharDevice != 0 {
		return nil, errors.New("The command is intended to work with pipes.")
	}

	// let's read those lines, shall we?
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func main() {
	lines, err := readUserInput()
	if err != nil {
		fmt.Println(err.Error())
		printUsage()
		os.Exit(1)
	}
	balloon.PrintBalloon(lines)
	characters.DrawParrot()
	fmt.Println()
}
