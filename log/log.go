package log

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
)

func InitDailyLog(prefix, path string) error {
	d, err := NewDailyLogWriter(prefix, path)
	if err != nil {
		return err
	}
	log.SetOutput(d)
	log.SetLevel(log.DebugLevel)
	return nil
}

func entry() *log.Entry {
	_, f, l, ok := runtime.Caller(2)
	if !ok {
		f, l = "", 0
	}
	return log.WithFields(log.Fields{
		"file": f,
		"line": l,
	})
}

func Debug(data ...interface{}) {
	entry().Debug(data)
}

func Info(data ...interface{}) {
	entry().Info(data)
}

func Warn(data ...interface{}) {
	entry().Warn(data)
}

func Error(data ...interface{}) {
	entry().Error(data)
}

func Fatal(data ...interface{}) {
	entry().Fatal(data)
}

type DailyLogWriter struct {
	prefix string
	path   string
	file   *os.File
}

func (d *DailyLogWriter) Write(p []byte) (int, error) {
	return d.file.Write(p)
}

func (d *DailyLogWriter) init() error {
	if d.file != nil {
		d.file.Close()
	}
	if len(d.path) == 0 {
		return errors.New("invalid log path")
	}
	s, err := os.Stat(d.path)
	if err != nil {
		err := os.MkdirAll(d.path, 0777)
		if err != nil {
			return err
		}
	} else {
		if !s.IsDir() {
			return errors.New("conflict log path ")
		}
	}

	t := time.Now().Format("2006-01-02")
	if d.prefix != "" {
		t = "-" + t
	}
	f := fmt.Sprintf("%s/%s%s.log", d.path, d.prefix, t)
	file, err := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	d.file = file
	return nil
}

func NewDailyLogWriter(prefix string, path string) (*DailyLogWriter, error) {
	d := &DailyLogWriter{
		prefix: prefix,
		path:   path,
	}
	if err := d.init(); err != nil {
		return nil, err
	}
	return d, nil
}
