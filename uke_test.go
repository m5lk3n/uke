package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	uke "lttl.dev/ukeapi/uke"
)

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

func TestC(t *testing.T) {
	testChord(t, "C", fretboardC)
}

func TestBm(t *testing.T) {
	testChord(t, "Bm", fretboardBm)
}

func testChord(t *testing.T, s string, expectedOutput []string) {
	std := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f := uke.Fretboard{Fretboard: uke.BlankFretboard}
	f.PrintFingers(s, false) // this std output gets captured

	w.Close()
	captured, _ := ioutil.ReadAll(r)
	os.Stdout = std

	expected := strings.Join(expectedOutput, "\n")

	if string(captured) != expected {
		t.Errorf("Expected %s, got %s", expected, captured)
	}
}
