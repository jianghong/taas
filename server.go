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

//func TimeIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//now := time.Now()
	//locationString := ps.ByName("location")
	//location, err := time.LoadLocation(locationString)

	//fmt.Fprintf(w, "%d", now.In(location))
//}

func main() {
	router := httprouter.New()
	// Routes
	router.GET("/", UnixTimestamp)
	//router.GET("/:location", TimeIn)

	// Start server
	fmt.Printf("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
