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

// finger labels for key
const indexFinger = "1"
const middleFinger = "2"
const ringFinger = "3"
const pinky = "4"

// basic chords
var C = Chord{"C", []SingleChord{{3, 2, ringFinger}}}
var Am = Chord{"Am", []SingleChord{{0, 1, middleFinger}}}
var F = Chord{"F", []SingleChord{{2, 0, indexFinger}, {0, 1, middleFinger}}}
var G = Chord{"G", []SingleChord{{1, 1, indexFinger}, {3, 1, middleFinger}, {2, 2, ringFinger}}}
var A = Chord{"A", []SingleChord{{1, 0, indexFinger}, {0, 1, middleFinger}}}
var Em = Chord{"Em", []SingleChord{{3, 1, indexFinger}, {2, 2, middleFinger}, {1, 3, ringFinger}}}
var D = Chord{"D", []SingleChord{{0, 1, indexFinger}, {1, 1, middleFinger}, {2, 1, ringFinger}}}
var Dm = Chord{"Dm", []SingleChord{{2, 0, indexFinger}, {0, 1, middleFinger}, {1, 1, ringFinger}}}
var E = Chord{"E", []SingleChord{{0, 0, indexFinger}, {3, 1, middleFinger}, {1, 3, pinky}}}
var Gbm = Chord{"F#/Gbm", []SingleChord{{1, 0, indexFinger}, {0, 1, middleFinger}, {2, 1, ringFinger}}}
var Bm = Chord{"Bm", []SingleChord{{1, 1, indexFinger}, {2, 1, indexFinger}, {3, 1, indexFinger}, {0, 3, ringFinger}}}
var Cm = Chord{"Cm", []SingleChord{{1, 2, indexFinger}, {2, 2, indexFinger}, {3, 2, indexFinger}}}
var Gm = Chord{"Gm", []SingleChord{{3, 0, indexFinger}, {1, 1, middleFinger}, {2, 2, ringFinger}}}
var G7 = Chord{"G7", []SingleChord{{2, 0, indexFinger}, {1, 1, middleFinger}, {3, 1, ringFinger}}}
var C7 = Chord{"C7", []SingleChord{{3, 0, indexFinger}}}
var A7 = Chord{"A7", []SingleChord{{1, 0, indexFinger}}}

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
func (f Fretboard) printFingers(c *Chord, printKey bool) {
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
		fmt.Printf("%s = index finger, %s = middle finger, %s = ring finger, %s = pinky\n", indexFinger, middleFinger, ringFinger, pinky)
	}
	fmt.Println()
}

func main() {
	f := Fretboard{blankFretboard}
	f.printFingers(&C, true)
	f.printFingers(&Am, false)
	f.printFingers(&F, false)
	f.printFingers(&G, false)
	f.printFingers(&A, false)
	f.printFingers(&Em, false)
	f.printFingers(&D, false)
	f.printFingers(&Dm, false)
	f.printFingers(&E, false)
	f.printFingers(&Gbm, false)
	f.printFingers(&Bm, false)
	f.printFingers(&Cm, false)
	f.printFingers(&Gm, false)
	f.printFingers(&G7, false)
	f.printFingers(&C7, false)
	f.printFingers(&A7, false)
}
