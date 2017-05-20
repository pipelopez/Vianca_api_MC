package models

import "gopkg.in/mgo.v2/bson"

type (
	AirLine struct {
		Id bson.ObjectId `json:"id" bson:"_id"`
		Code string `json:"code" bson:"code"`
		Name string `json:"name" bson:"name"`
		Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	}
)
