package models

import (
	"gorm.io/gorm"
	"time"
)

type Holiday struct {
	gorm.Model
	ID   int       `gorm:"primaryKey" json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func rewriteYear() {

	var rewriteYear struct {
		Year int
	}

	rY := rewriteYear
	rY.Year = year
}

var year int

const (
	dayJONe    = 1
	dayJTwo    = 2
	dayJThree  = 3
	dayJFour   = 4
	dayJFive   = 5
	dayJSix    = 6
	dayJR      = 7
	dayJEight  = 8
	dayF       = 23
	dayM       = 8
	dayMayOne  = 1
	dayMayNine = 9
	dayJun     = 12
	dayNov     = 4
	hour       = 00
	minute     = 00
	sec        = 00
	nSec       = 00
)

var (
	newYear                    = "Новый год"
	newYearHoliday             = "Новогодние каникулы"
	nativity                   = "Рождество Христово"
	defenderOfTheFatherlandDay = "День защитника Отечества"
	internationalWomenDay      = "Международный женский день"
	labourDay                  = "Праздник Весны и Труда"
	victoryDay                 = "День Победы"
	russiaDay                  = "День России"
	nationalUnityDay           = "День народного единства"
)

var Holidays = []Holiday{
	{ID: 1, Name: newYear, Date: time.Date(year, time.January, dayJONe, hour, minute, sec, nSec, time.Local)},
	{ID: 2, Name: newYearHoliday, Date: time.Date(year, time.January, dayJTwo, hour, minute, sec, nSec, time.Local)},
	{ID: 3, Name: newYearHoliday, Date: time.Date(year, time.January, dayJThree, hour, minute, sec, nSec, time.Local)},
	{ID: 4, Name: newYearHoliday, Date: time.Date(year, time.January, dayJFour, hour, minute, sec, nSec, time.Local)},
	{ID: 5, Name: newYearHoliday, Date: time.Date(year, time.January, dayJFive, hour, minute, sec, nSec, time.Local)},
	{ID: 6, Name: newYearHoliday, Date: time.Date(year, time.January, dayJSix, hour, minute, sec, nSec, time.Local)},
	{ID: 7, Name: nativity, Date: time.Date(year, time.January, dayJR, hour, minute, sec, nSec, time.Local)},
	{ID: 8, Name: newYearHoliday, Date: time.Date(year, time.January, dayJEight, hour, minute, sec, nSec, time.Local)},
	{ID: 9, Name: defenderOfTheFatherlandDay, Date: time.Date(year, time.February, dayF, hour, minute, sec, nSec, time.Local)},
	{ID: 10, Name: internationalWomenDay, Date: time.Date(year, time.March, dayM, hour, minute, sec, nSec, time.Local)},
	{ID: 11, Name: labourDay, Date: time.Date(year, time.May, dayMayOne, hour, minute, sec, nSec, time.Local)},
	{ID: 12, Name: victoryDay, Date: time.Date(year, time.May, dayMayNine, hour, minute, sec, nSec, time.Local)},
	{ID: 13, Name: russiaDay, Date: time.Date(year, time.June, dayJun, hour, minute, sec, nSec, time.Local)},
	{ID: 14, Name: nationalUnityDay, Date: time.Date(year, time.November, dayNov, hour, minute, sec, nSec, time.Local)},
}
