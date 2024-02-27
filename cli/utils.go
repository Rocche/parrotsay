package cli

import "fmt"

func clearScreen() {
	fmt.Print("\x1b[2J")
}

func moveCursorToHomePosition() {
	fmt.Print("\x1b[H")
}

// Clear flushes the screen
func Clear() {
	clearScreen()
	moveCursorToHomePosition()
}
