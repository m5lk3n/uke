package main

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
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

// f must not be a pointer to Fretboard as we don't modify the underlying matrix so that we get a blank board for every call
func (f Fretboard) setFingers(c *Chord) string {
	// indicate finger(s) on string(s) in matrix
	for _, sc := range c.singleChords {
		stringPos := sc.string * 3
		fretPos := sc.fret*2 + 1
		f.fretboard[fretPos] = replaceAtIndex(f.fretboard[fretPos], stringPos, sc.finger)
	}

	// return string with newline-separated fretboard where fingers are indicated
	return strings.Join(f.fretboard[:], "\n")
}

func (f *Fretboard) getKey() string {
	k := fmt.Sprintf("%s = index finger, %s = middle finger, %s = ring finger, %s = pinky\n", indexFinger, middleFinger, ringFinger, pinky)

	return k
}

func (f *Fretboard) printFingers(c *Chord, printKey bool) {
	// print board matrix, labeled with chord
	fmt.Println(c.name)
	fmt.Println(f.setFingers(c))

	if printKey {
		fmt.Println(f.getKey())
	}

	fmt.Println()
}

// NotFoundHandler to indicate that requested resource could not be found
func NotFoundHandler(c *gin.Context) {
	// log this event as it could be an attempt to break in...
	log.Infoln("Not found, requested URL path:", c.Request.URL.Path)
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "requested resource not found", "status": http.StatusNotFound})
}

// LivenessHandler always returns HTTP 200
func LivenessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "alive", "status": http.StatusOK})
}

// ReadinessHandler always returns HTTP 200
func ReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ready", "status": http.StatusOK})
}

// C
func CHandler(ctx *gin.Context) {
	f := Fretboard{blankFretboard}
	c := f.setFingers(&C)

	accept := ctx.Request.Header.Get("Accept")
	switch {
	case strings.Contains(accept, "text/html"):
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "UkeAPI", "chord": "C", "fretboard": c})
	case strings.Contains(accept, "json"):
		ctx.JSON(http.StatusOK, gin.H{"chord": "C", "fretboard": c, "status": http.StatusOK})
	default:
		ctx.Data(http.StatusOK, "application/text; charset=utf-8", []byte(c))
	}
}

// SetupRouter is published here to allow setup of tests
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil) // https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies

	// to debug: router.Use(gindump.Dump())

	router.Use(gin.Recovery()) // "recover from any panics", write 500 if any

	router.LoadHTMLGlob("templates/*")
	//	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.NoRoute(NotFoundHandler)

	// public, generic API
	router.GET("/healthy", LivenessHandler)
	router.GET("/ready", ReadinessHandler)

	// public, functional API
	router.GET("/C", CHandler)

	return router
}
func main() {
	/*
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
	*/
	router := SetupRouter()

	log.Infoln("UkeAPI start...")
	defer log.Infoln("UkeAPI shutdown!")

	// set port via PORT environment variable
	router.Run() // default port is 8080
}
