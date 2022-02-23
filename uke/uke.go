package uke

import (
	"fmt"
	"strings"

	common "lttl.dev/ukeapi/common"
)

type singleChord struct {
	string int    // 0 = G, 1 = C, 2 = E, 3 = A
	fret   int    // 0-based number of fret space
	finger string // 1 = index, 2 = middle, 3 = ring, 4 = pinky
}

type chord struct {
	name         string
	singleChords []singleChord
}

// finger labels for key
const indexFinger = "1"
const middleFinger = "2"
const ringFinger = "3"
const pinky = "4"

var m = make(map[string]chord) // chord name -> chord

func init() {
	// basic chords
	m["C"] = chord{"C", []singleChord{{3, 2, ringFinger}}}
	m["Am"] = chord{"Am", []singleChord{{0, 1, middleFinger}}}
	m["F"] = chord{"F", []singleChord{{2, 0, indexFinger}, {0, 1, middleFinger}}}
	m["G"] = chord{"G", []singleChord{{1, 1, indexFinger}, {3, 1, middleFinger}, {2, 2, ringFinger}}}
	m["A"] = chord{"A", []singleChord{{1, 0, indexFinger}, {0, 1, middleFinger}}}
	m["Em"] = chord{"Em", []singleChord{{3, 1, indexFinger}, {2, 2, middleFinger}, {1, 3, ringFinger}}}
	m["D"] = chord{"D", []singleChord{{0, 1, indexFinger}, {1, 1, middleFinger}, {2, 1, ringFinger}}}
	m["Dm"] = chord{"Dm", []singleChord{{2, 0, indexFinger}, {0, 1, middleFinger}, {1, 1, ringFinger}}}
	m["E"] = chord{"E", []singleChord{{0, 0, indexFinger}, {3, 1, middleFinger}, {1, 3, pinky}}}
	m["Gbm"] = chord{"F#/Gbm", []singleChord{{1, 0, indexFinger}, {0, 1, middleFinger}, {2, 1, ringFinger}}}
	m["Bm"] = chord{"Bm", []singleChord{{1, 1, indexFinger}, {2, 1, indexFinger}, {3, 1, indexFinger}, {0, 3, ringFinger}}}
	m["Cm"] = chord{"Cm", []singleChord{{1, 2, indexFinger}, {2, 2, indexFinger}, {3, 2, indexFinger}}}
	m["Gm"] = chord{"Gm", []singleChord{{3, 0, indexFinger}, {1, 1, middleFinger}, {2, 2, ringFinger}}}
	m["G7"] = chord{"G7", []singleChord{{2, 0, indexFinger}, {1, 1, middleFinger}, {3, 1, ringFinger}}}
	m["C7"] = chord{"C7", []singleChord{{3, 0, indexFinger}}}
	m["A7"] = chord{"A7", []singleChord{{1, 0, indexFinger}}}
}

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
func (f Fretboard) SetFingers(s string) string {
	if c, ok := m[s]; ok {
		// indicate finger(s) on string(s) in matrix
		for _, sc := range c.singleChords {
			stringPos := sc.string * 3
			fretPos := sc.fret*2 + 1
			f.Fretboard[fretPos] = common.ReplaceAtIndex(f.Fretboard[fretPos], stringPos, sc.finger)
		}
	}

	// return string with newline-separated fretboard where fingers are indicated
	return strings.Join(f.Fretboard[:], "\n")
}

func (f *Fretboard) getKey() string {
	k := fmt.Sprintf("%s = index finger, %s = middle finger, %s = ring finger, %s = pinky\n", indexFinger, middleFinger, ringFinger, pinky)

	return k
}

func (f *Fretboard) PrintFingers(s string, printKey bool) {
	// print board matrix, labeled with chord
	fmt.Println(s)
	fmt.Println(f.SetFingers(s))

	if printKey {
		fmt.Println(f.getKey())
	}

	fmt.Println()
}
