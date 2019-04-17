package handler

import (
	"log"

	"gopkg.in/mgo.v2"
)

var (
	mongodb *mgo.Session
)

func init() {
	// Database connection
	var err error
	mongodb, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}

	// Create indices
	if err := mongodb.Copy().DB("football_data").C("matches").EnsureIndex(mgo.Index{
		Key:    []string{"matchid"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

}
