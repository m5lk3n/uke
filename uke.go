package main

import "fmt"

func replaceAtIndex(s string, i int, r string) string {
	return s[:i] + r + s[i+1:]
}

type SingleChord struct {
	string int    // 0 = G, 1 = C, 2 = E, 3 = A
	fret   int    // 0-based number of fret space
	finger string // 1 = index, 2 = middle, 3 = ring, 4 = pinky
}

type Chord struct {
	name         string
	singleChords []SingleChord
}

type fretboardMatrix [9]string

var blankFretboard fretboardMatrix = fretboardMatrix{
	"+==+==+==+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
}

type Fretboard struct {
	fretboard fretboardMatrix
}

// f must not be a pointer to Fretboard as we don't modify the underlying matrix so that we get a blank board for every print
func (f Fretboard) printFingers(c Chord, printKey bool) {
	// indicate finger(s) on string(s) in matrix
	for _, sc := range c.singleChords {
		stringPos := sc.string * 3
		fretPos := sc.fret*2 + 1
		f.fretboard[fretPos] = replaceAtIndex(f.fretboard[fretPos], stringPos, sc.finger)
	}

	// print board matrix
	fmt.Println(c.name)
	for i := 0; i < len(f.fretboard); i++ {
		fmt.Println(f.fretboard[i])
	}
	if printKey {
		fmt.Println("1 = index finger, 2 = middle finger, 3 = ring finger, 4 = pinky")
	}
	fmt.Println()
}

func main() {
	C := Chord{"C", []SingleChord{{3, 2, "3"}}}
	Am := Chord{"Am", []SingleChord{{0, 1, "2"}}}
	F := Chord{"F", []SingleChord{{2, 0, "1"}, {0, 1, "2"}}}
	G := Chord{"G", []SingleChord{{1, 1, "1"}, {3, 1, "2"}, {2, 2, "3"}}}

	f := Fretboard{blankFretboard}
	f.printFingers(C, true)
	f.printFingers(Am, false)
	f.printFingers(F, false)
	f.printFingers(G, false)
}
