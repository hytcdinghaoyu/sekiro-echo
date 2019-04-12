package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Match struct {
	ID       bson.ObjectId `json:"-" bson:"_id"`
	MatchID  uint64        `json:"id"`
	Status   string
	MatchDay uint8
	UtcDate  time.Time
	Score    ScoreSum
	HomeTeam TeamSum
	AwayTeam TeamSum
}

type ScoreSum struct {
	Winner    string
	Duration  string
	FullTime  ScoreDesc
	HalfTime  ScoreDesc
	ExtraTime ScoreDesc
	Penalties ScoreDesc
}

type ScoreDesc struct {
	HomeTeam uint8
	AwayTeam uint8
}

type TeamSum struct {
	ID   uint16
	Name string
}
