package jobs

import (
	"log"
	"time"
)

type Test struct {
}

func (Test) Frequency() string {
	return "30 * * * * *" //每30秒
}

func (Test) Run() {
	log.Println("Test Crontab", time.Now().Unix())
}
