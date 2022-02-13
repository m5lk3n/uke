package main

import "fmt"

type chord struct {
	name   string
	string int    // 0 = G, 1 = C, 2 = E, 3 = A
	fret   int    // 0-based number of fret space
	finger string // 1 = index, 2 = middle, 3 = ring, 4 = pinky
}

func replaceAtIndex(s string, i int, r string) string {
	return s[:i] + r + s[i+1:]
}

func printFretboard(c chord, printKey bool) {

	var fretboard [9]string
	fretboard[0] = "+==+==+==+"
	fretboard[1] = "|  |  |  |"
	fretboard[2] = "+--+--+--+"
	fretboard[3] = "|  |  |  |"
	fretboard[4] = "+--+--+--+"
	fretboard[5] = "|  |  |  |"
	fretboard[6] = "+--+--+--+"
	fretboard[7] = "|  |  |  |"
	fretboard[8] = "+--+--+--+"

	// indicate finger on string
	stringPos := c.string * 3
	fretPos := c.fret*2 + 1
	fretboard[fretPos] = replaceAtIndex(fretboard[fretPos], stringPos, c.finger)

	fmt.Println(c.name)
	for i := 0; i < len(fretboard); i++ {
		fmt.Println(fretboard[i])
	}
	if printKey {
		fmt.Println("1 = index finger, 2 = middle finger, 3 = ring finger, 4 = pinky")
	}
	fmt.Println()
}

func main() {
	C := chord{"C", 3, 2, "3"}
	Am := chord{"Am", 0, 1, "2"}

	printFretboard(C, true)
	printFretboard(Am, false)
}
