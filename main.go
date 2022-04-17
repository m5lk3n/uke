package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	uke "lttl.dev/ukeapi/uke"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// NotFoundHandler indicates that a requested resource could not be found.
func NotFoundHandler(c *gin.Context) {
	// log this event as it could be an attempt to break in...
	log.Infoln("Not found, requested URL path:", c.Request.URL.Path)
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "requested resource not found", "status": http.StatusNotFound})
}

// LivenessHandler always returns HTTP 200.
func LivenessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "alive", "status": http.StatusOK})
}

// ReadinessHandler always returns HTTP 200.
func ReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ready", "status": http.StatusOK})
}

// ChordHandler prints the fretboard with the chord indicated (if chord is
// found).
func ChordHandler(c *gin.Context) {
	chordParam := c.Param("chord")
	keyParam := c.Param("key")

	f := uke.Fretboard{Fretboard: uke.BlankFretboard}
	ff := f.SetFingers(chordParam)
	var k string
	if strings.EqualFold(keyParam, "/key") {
		k = f.GetKey()
	}

	accept := c.Request.Header.Get("Accept")
	switch {
	case strings.Contains(accept, "text/html"):
		c.HTML(http.StatusOK, "single.tmpl", gin.H{"title": "UkeAPI", "label": chordParam, "content": ff, "key": k})
	case strings.Contains(accept, "json"):
		c.JSON(http.StatusOK, gin.H{"chord": chordParam, "fretboard": ff, "key": k, "status": http.StatusOK})
	default:
		text := fmt.Sprintf("%s\n%s\n%s", chordParam, ff, k)
		c.Data(http.StatusOK, "application/text; charset=utf-8", []byte(text))
	}
}

// ChordsHandler prints the fretboard with up to 4 chord(s) indicated (if chord
// is found). Supported output format is HTML only, any other will return "not
// implemented/supported".
func ChordsHandler(c *gin.Context) {
	accept := c.Request.Header.Get("Accept")
	switch {
	case strings.Contains(accept, "text/html"):
		break
	case strings.Contains(accept, "json"):
		c.JSON(http.StatusNotImplemented, gin.H{"message": "requested format not supported", "status": http.StatusNotImplemented})
	default:
		c.Data(http.StatusNotImplemented, "application/text; charset=utf-8", []byte("requested format not supported"))
	}

	chordsParam := c.Param("chords")
	keyParam := c.Param("key")

	const max = 4 // some arbitrary upper bound for fretboards with chords to display next to each other, goes along with label and content entries in multi.tmpl
	l := make([]string, max)
	b := make([]string, max)
	var f uke.Fretboard
	chords := strings.Split(chordsParam, "-")
	for i, chord := range chords {
		if i == max {
			break
		}

		f = uke.Fretboard{Fretboard: uke.BlankFretboard}
		ff := f.SetFingers(chord)
		l[i] = chord
		b[i] = ff
	}

	var k string
	if strings.EqualFold(keyParam, "/key") {
		k = f.GetKey()
	}

	c.HTML(http.StatusOK, "multi.tmpl", gin.H{"title": "UkeAPI", "label0": l[0], "content0": b[0], "label1": l[1], "content1": b[1], "label2": l[2], "content2": b[2], "label3": l[3], "content3": b[3], "key": k})
}

// ChordNamesHandler prints the supported chord names.
func ChordNamesHandler(c *gin.Context) {
	cn := uke.GetChordNames()

	accept := c.Request.Header.Get("Accept")
	switch {
	case strings.Contains(accept, "text/html"):
		c.HTML(http.StatusOK, "single.tmpl", gin.H{"title": "UkeAPI", "label": "Supported Chords", "content": cn})
	case strings.Contains(accept, "json"):
		c.JSON(http.StatusOK, gin.H{"chord": "Supported Chords", "fretboard": cn, "status": http.StatusOK})
	default:
		c.Data(http.StatusOK, "application/text; charset=utf-8", []byte(cn))
	}
}

// SetupRouter is published to allow setup of tests.
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil) // https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies

	// to debug: router.Use(gindump.Dump())

	router.Use(gin.Recovery()) // "recover from any panics", write 500 if any

	router.LoadHTMLGlob("templates/*")
	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.NoRoute(NotFoundHandler)

	v1 := router.Group("/api/v1")
	{
		// public, generic API
		v1.GET("/healthy", LivenessHandler)
		v1.GET("/ready", ReadinessHandler)

		// public, functional API
		v1.GET("/chord/:chord", ChordHandler)
		v1.GET("/chord/:chord/*key", ChordHandler)
		v1.GET("/chords/:chords", ChordsHandler)
		v1.GET("/chords/:chords/*key", ChordsHandler)
		v1.GET("/chordNames", ChordNamesHandler)
	}

	return router
}

func main() {
	serve := flag.Bool("serve", false, "run in HTTP server mode (false by default)")
	chord := flag.String("chord", "", "display single chord, e.g. C or BM (case-sensitive) (ignored in server mode)")
	key := flag.Bool("key", false, "show key (false by default)")
	flag.Parse()

	if *serve {
		router := SetupRouter()

		log.Infoln("UkeAPI starting...")
		defer log.Infoln("UkeAPI shutdown!")

		// set port via PORT environment variable
		router.Run() // default port is 8080
	} else if *chord != "" {
		f := uke.Fretboard{Fretboard: uke.BlankFretboard}
		f.PrintFingers(*chord, *key)
	} else {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}
