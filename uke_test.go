package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	uke "lttl.dev/ukeapi/uke"
)

var fretboardEmpty []string = []string{
	"NA",
	"+==+==+==+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"",
	"",
}

var fretboardC []string = []string{
	"C",
	"+==+==+==+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  3",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"",
	"",
}

var fretboardBm []string = []string{
	"Bm",
	"+==+==+==+",
	"|  |  |  |",
	"+--+--+--+",
	"|  1  1  1",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"3  |  |  |",
	"+--+--+--+",
	"",
	"",
}

var fretboardFKey []string = []string{
	"F",
	"+==+==+==+",
	"|  |  1  |",
	"+--+--+--+",
	"2  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"|  |  |  |",
	"+--+--+--+",
	"1 = index finger, 2 = middle finger, 3 = ring finger, 4 = pinky",
	"",
	"",
	"",
}

func TestNotFound(t *testing.T) {
	testChord(t, "NA", false, fretboardEmpty)
}

func TestC(t *testing.T) {
	testChord(t, "C", false, fretboardC)
}

func TestBm(t *testing.T) {
	testChord(t, "Bm", false, fretboardBm)
}

func TestFKey(t *testing.T) {
	testChord(t, "F", true, fretboardFKey)
}

func testChord(t *testing.T, s string, k bool, expectedOutput []string) {
	std := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f := uke.Fretboard{Fretboard: uke.BlankFretboard}
	f.PrintFingers(s, k) // this std output gets captured

	w.Close()
	captured, _ := ioutil.ReadAll(r)
	os.Stdout = std

	expected := strings.Join(expectedOutput, "\n")

	if string(captured) != expected {
		t.Errorf("Expected %s, got %s", expected, captured)
	}
}
