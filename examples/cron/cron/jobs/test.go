package jobs

import (
	"log"
	"time"
)

type Test struct {
}

func (Test) Frequency() string {
	return "30 * * * * *" //every 30s
}

func (Test) Run() {
	log.Println("Test Crontab", time.Now().Unix())
}
