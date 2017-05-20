package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
)

type (
	Flight struct {
		Id bson.ObjectId `json:"id" bson:"_id"`
		FlightCode string `json:"flightCode" bson:"flightCode"`
		Origin string `json:"origin" bson:"origin"`
    Destination string `json:"destination" bson:"destination"`
    Price float32 `json:"price" bson:"price"`
    Currency string `json:"currency" bson:"currency"`
    date time.Time `json:"date" bson:"date"`
		Passengers int `json:"passengers" bson:"passengers"`
	}
)
