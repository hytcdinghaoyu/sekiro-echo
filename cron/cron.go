package main

import (
	"sekiro_echo/jobs"

	"github.com/jasonlvhit/gocron"
)

func main() {
	s := gocron.NewScheduler()
	//未来7日赛程
	s.Every(1).Minute().Do(jobs.AddScheduledMatch)

	//进行中比赛的比分

	<-s.Start()
	select {}
}
