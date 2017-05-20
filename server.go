package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/swhite24/vianca_api/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	fc := controllers.NewFlightController(getSession())

	r.GET("/flight/:id", fc.GetFlight)

	// Create a new user
	r.POST("/flight", fc.CreateFlight)

	// Remove an existing user
	r.DELETE("/flight/:id", fc.RemoveFlight)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
