package helper

import "testing"

func TestMonthStart(t *testing.T) {
	t.Log(MonthStart())
}

func TestTodayStart(t *testing.T) {
	t.Log(TodayStart())
}

func TestTodayEnd(t *testing.T) {
	t.Log(TodayEnd())
}

func TestNowUnix(t *testing.T) {
	unix := NowUnix()
	t.Log(unix)
}

func TestNowDate(t *testing.T) {
	nd := NowDate()
	t.Log(nd)
}

func TestNowDateTime(t *testing.T) {
	ndt := NowDateTime()
	t.Log(ndt)
}

func TestParseDate(t *testing.T) {
	tm, err := ParseDate("2018-01-08")
	if err != nil {
		t.Error(err)
	}
	t.Log(tm.String())
}

func TestParseDateTime(t *testing.T) {
	tm, err := ParseDateTime("2018-01-08 16:34:00")
	if err != nil {
		t.Error(err)
	}
	t.Log(tm.String())
}

func TestParseStringTime(t *testing.T) {
	tm, err := ParseStringTime("2018-01-08 14:23:00", "Asia/Shanghai")
	if err != nil {
		t.Error(err)
	}
	t.Log(tm.String())
}
