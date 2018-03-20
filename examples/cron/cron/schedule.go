package crons

import (
	"log"
	"github.com/hhxsv5/gin-x/cron"
	"github.com/hhxsv5/gin-x/examples/cron/cron/jobs"
)

var (
	CronManger cron.Manager
)

func init() {

	CronManger := cron.NewManager()
	CronManger.Register(jobs.Test{})
	//...

	CronManger.Start()
	log.Println("Cron started success")
}
