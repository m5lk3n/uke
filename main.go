package main

import (
	"net/http"
	"strings"

	uke "lttl.dev/ukeapi/uke"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

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

// ChordHandler prints the fretboard with the chord indicated (if chord is found)
func ChordHandler(c *gin.Context) {
	chord := c.Param("chord")

	f := uke.Fretboard{Fretboard: uke.BlankFretboard}
	ff := f.SetFingers(chord)

	accept := c.Request.Header.Get("Accept")
	switch {
	case strings.Contains(accept, "text/html"):
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "UkeAPI", "chord": chord, "fretboard": ff})
	case strings.Contains(accept, "json"):
		c.JSON(http.StatusOK, gin.H{"chord": chord, "fretboard": ff, "status": http.StatusOK})
	default:
		c.Data(http.StatusOK, "application/text; charset=utf-8", []byte(ff))
	}
}

// SetupRouter is published here to allow setup of tests
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil) // https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies

	// to debug: router.Use(gindump.Dump())

	router.Use(gin.Recovery()) // "recover from any panics", write 500 if any

	router.LoadHTMLGlob("templates/*")
	// router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.NoRoute(NotFoundHandler)

	// public, generic API
	router.GET("/healthy", LivenessHandler)
	router.GET("/ready", ReadinessHandler)

	// public, functional API
	router.GET("/:chord", ChordHandler)

	return router
}

func main() {
	router := SetupRouter()

	log.Infoln("ukeapi start...")
	defer log.Infoln("ukeapi shutdown!")

	// set port via PORT environment variable
	router.Run() // default port is 8080
}
