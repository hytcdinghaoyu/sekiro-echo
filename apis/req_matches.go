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

	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create indices
	if err = db.Copy().DB("football_data").C("matches").EnsureIndex(mgo.Index{
		Key:    []string{"matchid"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	url := "https://api.football-data.org/v2/matches"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", "4958466805ba41f680595be4fc92ac87")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	matchesRep := MatchesRep{}
	json.Unmarshal(body, &matchesRep)

	var matchesCollection *mgo.Collection
	matchesCollection = db.DB("football_data").C("matches")
	for _, match := range matchesRep.Matches {
		match.ID = bson.NewObjectId()
		fmt.Println(match)

		var matchFind model.Match
		if err = matchesCollection.Find(bson.M{"matchid": match.MatchID}).One(&matchFind); err != nil {
			//if not found insert
			if err == mgo.ErrNotFound {
				//return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
				if err = matchesCollection.Insert(&match); err != nil {
					log.Fatal(err)
					return
				}
			}
		} else {
			fmt.Println(match.Score.FullTime.HomeTeam)
			fmt.Println(matchFind.Status)
			matchesCollection.Update(bson.M{"matchid": matchFind.MatchID}, bson.M{"$set": bson.M{"status": "changed1"}})

		}

	}

	os.Exit(0)

}
