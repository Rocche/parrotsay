package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func randomPartyPhrase() string {
	return "YOOOO FINALLY YOU DECIDED TO PARTYYYYYY!"
}

func PrintUsage() {
	fmt.Println("The 'parrotsay' command can operate in three modes:")
	fmt.Println("1. No input at all. Parrot will say a random party phrase!")
	fmt.Println("Example: parrotsay")
	fmt.Println()
	fmt.Println("2. Input as stdin. The input is passed through a pipe")
	fmt.Println("Example: forutne | parrotsay")
	fmt.Println("Example: echo '3dyc0n3G p075' | parrotsay")
	fmt.Println()
	fmt.Println("3. Program argument.")
	fmt.Println("Example: parrotsay 'YOOOO LET'S PARTYYY'")
}

func argsMode(args []string) ([]string, error) {
	if len(args) == 0 {
		return nil, errors.New("Cannot use args mode with no arguments!")
	}
	return []string{args[0]}, nil
}

func stdinMode() ([]string, error) {
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

func ParseUserInput() ([]string, error) {
	// first check if there are args: if so, args mode have precedence
	args := os.Args[1:]
	if len(args) > 0 {
		return argsMode(args)
	}
	info, _ := os.Stdin.Stat()
	if info.Mode()&os.ModeCharDevice == 0 {
		return stdinMode()
	}
	return []string{randomPartyPhrase()}, nil
}
