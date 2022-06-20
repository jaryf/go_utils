package utils

import (
	"time"
)

const (
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

// NowTime 返回当前时间time对象
func NowTime() (now time.Time) {
	return time.Now().Local()
}

// NowTimeStr 返回当前时间字符串
func NowTimeStr() (now string) {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}

// StartOfMin 获取这一分钟的开始时间
func StartOfMin(t ...time.Time) time.Time {
	var nTime time.Time
	if len(t) > 0 {
		nTime = t[0]
	} else {
		nTime = NowTime()
	}
	return time.Date(nTime.Year(), nTime.Month(), nTime.Day(), nTime.Hour(), nTime.Minute(), 0, 0, nTime.Location())
}

// StartOfHour 获取这一小时的开始时间
func StartOfHour(t ...time.Time) time.Time {
	var nTime time.Time
	if len(t) > 0 {
		nTime = t[0]
	} else {
		nTime = NowTime()
	}
	return time.Date(nTime.Year(), nTime.Month(), nTime.Day(), nTime.Hour(), 0, 0, 0, nTime.Location())
}

// StartOfDay 获取这一天的开始时间
func StartOfDay(t ...time.Time) time.Time {
	var nTime time.Time
	if len(t) > 0 {
		nTime = t[0]
	} else {
		nTime = NowTime()
	}
	return time.Date(nTime.Year(), nTime.Month(), nTime.Day(), 0, 0, 0, 0, nTime.Location())
}

// StartOfWeek 周开始时间,mondayStart周一是否是开始时间
func StartOfWeek(mondayStart bool, t ...time.Time) time.Time {
	var nTime time.Time
	if len(t) > 0 {
		nTime = t[0]
	} else {
		nTime = NowTime()
	}

	weekDay := int(nTime.Weekday())
	if mondayStart {
		weekDay = weekDay - 1
	}
	return StartOfDay(nTime).AddDate(0, 0, -weekDay)
}

// StartOfMonth 月开始时间
func StartOfMonth(t ...time.Time) time.Time {
	var nTime time.Time
	if len(t) > 0 {
		nTime = t[0]
	} else {
		nTime = NowTime()
	}
	return time.Date(nTime.Year(), nTime.Month(), 1, 0, 0, 0, 0, nTime.Location())
}

// StartOfYear 年开始时间
func StartOfYear(t ...time.Time) time.Time {
	var nTime time.Time
	if len(t) > 0 {
		nTime = t[0]
	} else {
		nTime = NowTime()
	}
	return time.Date(nTime.Year(), 1, 1, 0, 0, 0, 0, nTime.Location())
}

// EndOfMin 获取这一分钟的结束时间
func EndOfMin(t ...time.Time) time.Time {
	return StartOfMin(t...).Add(time.Minute - time.Nanosecond)
}

// EndOfHour 获取这一小时的结束时间
func EndOfHour(t ...time.Time) time.Time {
	return StartOfHour(t...).Add(time.Hour - time.Nanosecond)
}

// EndOfDay 获取这一天的结束时间
func EndOfDay(t ...time.Time) time.Time {
	return StartOfDay(t...).Add(-time.Nanosecond)
}

// EndOfWeek 周结束时间,mondayStart周一是否是开始时间
func EndOfWeek(mondayStart bool, t ...time.Time) time.Time {
	return StartOfWeek(mondayStart, t...).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfMonth 月结束时间
func EndOfMonth(t ...time.Time) time.Time {
	return StartOfMonth(t...).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfYear 年结束时间
func EndOfYear(t ...time.Time) time.Time {
	return StartOfYear(t...).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// CurrentWeekRange 获取当前周的开始结束时间
func CurrentWeekRange(mondayStart bool) (currentWeekStart, currentWeekEnd time.Time) {
	currentWeekStart = StartOfWeek(mondayStart)
	currentWeekEnd = currentWeekStart.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return
}

// LastWeekRange 获取上一周的开始结束时间
func LastWeekRange(mondayStart bool) (lastWeekStart, lastWeekEnd time.Time) {
	currentWeekStart := StartOfWeek(mondayStart)
	lastWeekStart = currentWeekStart.AddDate(0, 0, -7)
	lastWeekEnd = lastWeekStart.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return
}
