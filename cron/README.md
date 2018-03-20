gin-slim-cron
=============
A smart cron for gin framework

## Usage

1. Create Job
```Go
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
```

2. Create Job Manager & Register Job
```Go
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
```

## License

[MIT](https://github.com/hhxsv5/gin-slim-router/blob/master/LICENSE)
