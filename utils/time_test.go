package utils

import (
	"testing"
)

func TestNowTime(t *testing.T) {
	nowTime := NowTime()
	t.Log(nowTime)
}

func TestNowTimeStr(t *testing.T) {
	nowTimeStr := NowTimeStr()
	t.Log(nowTimeStr)
}

func TestLastWeekRange(t *testing.T) {
	lastWeekStart, lastWeekEnd := LastWeekRange(true)
	t.Log(lastWeekStart, lastWeekEnd)
	t.Log()
}

func TestStartOfWeek(t *testing.T) {
	startOfWeek := StartOfWeek(true)
	t.Log(startOfWeek)
}

func TestEndOfWeek(t *testing.T) {
	endOfWeek := EndOfWeek(true)
	t.Log(endOfWeek)
}

func TestCurrentWeekRange(t *testing.T) {
	weekStart, weekEnd := CurrentWeekRange(true)
	t.Log(weekStart, weekEnd)
}
