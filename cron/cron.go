package main

import (
	"sekiro_echo/jobs"

	"github.com/jasonlvhit/gocron"
)

func main() {

	s := gocron.NewScheduler()

	//未来7日赛程
	s.Every(1).Day().Do(jobs.AddScheduledMatch)

	//实时比分
	s.Every(1).Minute().Do(jobs.UpdateScore)

	<-s.Start()
}
