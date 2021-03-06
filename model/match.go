package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Match struct {
	ID            bson.ObjectId `json:"-" bson:"_id"`
	MatchID       uint64        `json:"id"`
	Competition   CompetitionSum
	Status        string `bson:"Status"`
	MatchDay      uint8
	UtcDate       time.Time
	MatchDate     string
	MatchDateTime string
	Score         ScoreSum
	HomeTeam      TeamSum
	AwayTeam      TeamSum
}

type ScoreSum struct {
	Winner    string    `bson:"Winner"`
	Duration  string    `bson:"Duration"`
	FullTime  ScoreDesc `bson:"FullTime"`
	HalfTime  ScoreDesc `bson:"HalfTime"`
	ExtraTime ScoreDesc `bson:"ExtraTime"`
	Penalties ScoreDesc `bson:"Penalties"`
}

type ScoreDesc struct {
	HomeTeam uint8 `bson:"HomeTeam"`
	AwayTeam uint8 `bson:"AwayTeam"`
}

type TeamSum struct {
	ID   uint16
	Name string
}

type CompetitionSum struct {
	ID   uint64
	Name string
}
