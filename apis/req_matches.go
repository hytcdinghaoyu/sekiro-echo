package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"sekiro_echo/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MatchesRep struct {
	Count   uint8
	Filters interface{}
	Matches []model.Match
}

func main() {

	// match := &model.Match{
	// 	ID: bson.NewObjectId(),
	// }
	// match.Score.Winner = "barca"
	// match.MatchID = 12201

	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}

	// Create indices
	if err = db.Copy().DB("football_data").C("matches").EnsureIndex(mgo.Index{
		Key:    []string{"matchid"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	url := "http://api.football-data.org/v2/matches"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", "4958466805ba41f680595be4fc92ac87")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	matchesRep := MatchesRep{}
	json.Unmarshal(body, &matchesRep)

	for _, match := range matchesRep.Matches {
		match.ID = bson.NewObjectId()
		fmt.Println(match)
		if err = db.DB("football_data").C("matches").Insert(&match); err != nil {
			fmt.Println(err)
			return
		}
	}

	os.Exit(0)

}
