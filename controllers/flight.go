package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/swhite24/vianca_api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	FlightController struct {
		session *mgo.Session
	}
)

func NewFlightController(s *mgo.Session) *FlightController {
	return &FlightController{s}
}

func (fc FlightController) GetFlight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	f := models.Flight{}

	if err := fc.session.DB("vianca_api").C("flights").FindId(oid).One(&f); err != nil {
		w.WriteHeader(404)
		return
	}

	fj, _ := json.Marshal(f)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", fj)
}

func (fc FlightController) CreateFlight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f := models.Flight{}

	json.NewDecoder(r.Body).Decode(&f)

	f.Id = bson.NewObjectId()

	fc.session.DB("vianca_api").C("flights").Insert(f)

	fj, _ := json.Marshal(f)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", fj)
}

func (fc FlightController) RemoveFlight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("vianca_api").C("flights").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}
