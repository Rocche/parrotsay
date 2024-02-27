package characters

import (
	"fmt"
)

var frame = 0

// DrawParrot draws... well... a parrot?
func DrawParrot() {
	parrot := ParrotFrames[0]
	for _, c := range parrot {
		fmt.Printf(runeToColoredBlock(c, 0))
	}
}

func DrawAnimatedParrot() {
	parrot := ParrotFrames[frame]

	for _, c := range parrot {
		fmt.Printf(runeToColoredBlock(c, frame))
	}
	frame = (frame + 1) % len(ParrotFrames)
}
