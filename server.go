package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"time"
)

func UnixTimestamp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	now := time.Now().Unix()
    fmt.Fprintf(w, "%d", now)
}

func TimeIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	now := time.Now()
	tz := ps.ByName("tz")
	timezone, err := time.LoadLocation(tz)

	if err != nil {
		fmt.Fprintf(w, "Invalid timezone.")
		return
	}

	fmt.Fprintf(w, "%s", now.In(timezone).String())
}

func main() {
	router := httprouter.New()
	// Routes
	router.GET("/", UnixTimestamp)
	router.GET("/:tz", TimeIn)

	// Start server
	fmt.Printf("Listening on port 3001...")
	log.Fatal(http.ListenAndServe(":3001", router))
}
