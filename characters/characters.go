package characters

import (
	"fmt"
)

// local variable that tracks which parrot frame we are displaying
var frame = 0

// DrawParrot draws... well... a parrot?
func DrawParrot() {
	parrot := ParrotFrames[0]
	for _, c := range parrot {
		fmt.Printf(runeToColoredBlock(c, 0))
	}
}

// DrawAnimatedParrot draws... well... a parrot? BUT ANIMATED!
func DrawAnimatedParrot() {
	parrot := ParrotFrames[frame]

	for _, c := range parrot {
		fmt.Printf(runeToColoredBlock(c, frame))
	}
	frame = (frame + 1) % len(ParrotFrames)
}
