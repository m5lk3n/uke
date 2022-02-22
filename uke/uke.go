package uke

import (
	"fmt"
	"strings"
)

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

var BlankFretboard fretboardMatrix = fretboardMatrix{
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
	Fretboard fretboardMatrix
}

// f must not be a pointer to Fretboard as we don't modify the underlying matrix so that we get a blank board for every call
func (f Fretboard) SetFingers(c *Chord) string {
	// indicate finger(s) on string(s) in matrix
	for _, sc := range c.singleChords {
		stringPos := sc.string * 3
		fretPos := sc.fret*2 + 1
		f.Fretboard[fretPos] = replaceAtIndex(f.Fretboard[fretPos], stringPos, sc.finger)
	}

	// return string with newline-separated fretboard where fingers are indicated
	return strings.Join(f.Fretboard[:], "\n")
}

func (f *Fretboard) getKey() string {
	k := fmt.Sprintf("%s = index finger, %s = middle finger, %s = ring finger, %s = pinky\n", indexFinger, middleFinger, ringFinger, pinky)

	return k
}

func (f *Fretboard) PrintFingers(c *Chord, printKey bool) {
	// print board matrix, labeled with chord
	fmt.Println(c.name)
	fmt.Println(f.SetFingers(c))

	if printKey {
		fmt.Println(f.getKey())
	}

	fmt.Println()
}
