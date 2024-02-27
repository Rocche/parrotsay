// Thank you Flavio for the amazing tutorial "Building a CLI command with Go: cowsay"
// https://flaviocopes.com/go-tutorial-cowsay/
// Please go check it out!
package main

import (
	"fmt"
	"os"
	"parrotsay/balloon"
	"parrotsay/characters"
	"parrotsay/cli"
	"time"
)

func main() {
	lines, err := cli.ParseUserInput()
	if err != nil {
		fmt.Println(err.Error())
		cli.PrintUsage()
		os.Exit(1)
	}
	for {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		balloon.PrintBalloon(lines)
		characters.DrawAnimatedParrot()
		d, _ := time.ParseDuration("50ms")
		time.Sleep(d)
	}
}
