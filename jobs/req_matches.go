package jobs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"sekiro_echo/model"

	"github.com/fatih/structs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongodb *mgo.Session
)

const (
	MONGO_ADDR = "localhost"
	DATA_URL   = "https://api.football-data.org/v2/matches"
	AUTH_TOKEN = "4958466805ba41f680595be4fc92ac87"
)

func init() {
	var err error
	mongodb, err = mgo.Dial(MONGO_ADDR)
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

	f, err := os.OpenFile("cron.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer f.Close()
	log.SetOutput(f)
}

type MatchesRep struct {
	Count   uint8
	Filters interface{}
	Matches []model.Match
}

//每天运行一次，获取未来七天的赛程
func AddScheduledMatch() {
	log.Println("=====>Start running job: AddScheduledMatch")

	req, _ := http.NewRequest("GET", DATA_URL, nil)
	req.Header.Add("X-Auth-Token", AUTH_TOKEN)

	q := req.URL.Query()
	q.Add("dateFrom", time.Now().Format("2006-01-02"))
	q.Add("dateTo", time.Now().AddDate(0, 0, 7).Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Connecting to the server Error,Waiting for next run")
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	matchesRep := MatchesRep{}
	json.Unmarshal(body, &matchesRep)

	var matchesCollection *mgo.Collection
	matchesCollection = mongodb.DB("football_data").C("matches")
	for _, match := range matchesRep.Matches {
		match.ID = bson.NewObjectId()
		var matchFind model.Match
		if err := matchesCollection.Find(bson.M{"matchid": match.MatchID}).One(&matchFind); err == mgo.ErrNotFound {
			//if not found insert
			log.Printf("Added match: %s vs %s \n", match.HomeTeam.Name, match.AwayTeam.Name)
			if err = matchesCollection.Insert(&match); err != nil {
				log.Fatal(err)
				return
			}
		}
	}

	log.Println("=====>End running job: AddScheduledMatch")
}

//每分钟运行一次，更新比分
func UpdateScore() {
	log.Println("=====>Start running job: UpdateScore")

	req, _ := http.NewRequest("GET", DATA_URL, nil)
	req.Header.Add("X-Auth-Token", AUTH_TOKEN)

	q := req.URL.Query()
	q.Add("status", "IN_PLAY")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Connecting to the server Error,Waiting for next run")
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	matchesRep := MatchesRep{}
	json.Unmarshal(body, &matchesRep)

	var matchesCollection *mgo.Collection
	matchesCollection = mongodb.DB("football_data").C("matches")
	for _, match := range matchesRep.Matches {
		var matchFind model.Match
		if err := matchesCollection.Find(bson.M{"matchid": match.MatchID}).One(&matchFind); err == mgo.ErrNotFound {
			log.Println("match not found")
		} else {
			//update score
			scoreMap := structs.Map(match.Score)
			log.Printf("%s %d : %d %s \n", match.HomeTeam.Name, match.Score.FullTime.HomeTeam, match.Score.FullTime.AwayTeam, match.AwayTeam.Name)
			matchesCollection.Update(bson.M{"matchid": matchFind.MatchID}, bson.M{"$set": bson.M{"score": scoreMap}})
		}

	}

	log.Println("=====>End running job: UpdateScore")

}
