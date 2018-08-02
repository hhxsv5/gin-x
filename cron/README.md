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
```

3. Import package `crons`
```Go
import _ "xxx/crons"
```

## License

[MIT](https://github.com/hhxsv5/gin-x/blob/master/LICENSE)
