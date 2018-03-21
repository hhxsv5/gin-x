package crons

import (
	"log"
	"github.com/hhxsv5/gin-x/cron"
	"github.com/hhxsv5/gin-x/examples/cron/cron/jobs"
	"time"
)

var (
	CronManger cron.Manager
)

func init() {

	CronManger := cron.NewManager()
	CronManger.Register(jobs.Test{}) // Register Job
	CronManger.RegisterFunc("0 * * * * *", func() { // Register Func
		log.Println(time.Now().Unix())
	})
	//...

	CronManger.Start()
	log.Println("Cron started success")
}
