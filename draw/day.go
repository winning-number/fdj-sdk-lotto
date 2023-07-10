package draw

import (
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	csvDayMonday         = "LUNDI   "
	csvDayTuesday        = "MARDI   "
	csvDayWednesday      = "MERCREDI"
	csvDayThursday       = "JEUDI   "
	csvDayFriday         = "VENDREDI"
	csvDaySaturday       = "SAMEDI  "
	csvDaySunday         = "DIMANCHE"
	csvDayShortMonday    = "LU"
	csvDayShortTuesday   = "MA"
	csvDayShortWednesday = "ME"
	csvDayShortThursday  = "JE"
	csvDayShortFriday    = "VE"
	csvDayShortSaturday  = "SA"
	csvDayShortSunday    = "DI"
)

// Day list
const (
	DayMonday    Day = "MONDAY"
	DayTuesday   Day = "TUESDAY"
	DayWednesday Day = "WEDNESDAY"
	DayThursday  Day = "THURSDAY"
	DayFriday    Day = "FRIDAY"
	DaySaturday  Day = "SATURDAY"
	DaySunday    Day = "SUNDAY"
)

type Day string

// dateConverter detect the date format and convert it to a date type [time.Time].
// The input date support two format:
//   - 20060102 default
//   - 02/01/2006 (need to be reverse)
func dateConverter(date string) (time.Time, error) {
	var t time.Time
	var err error

	format := "20060102"
	if strings.ContainsRune(date, '/') {
		format = "02/01/2006"
	}
	if t, err = time.Parse(format, date); err != nil {
		return time.Time{}, errors.Wrap(ErrCSVDate, err.Error())
	}

	return t, nil
}

// dayConverter detect the day format from a csv and convert it to a [Day] type.
// If the day is unknown a [ErrDayUnknown] is return.
func dayConverter(day string) (Day, error) {
	switch day {
	case csvDayMonday, csvDayShortMonday:
		return DayMonday, nil
	case csvDayTuesday, csvDayShortTuesday:
		return DayTuesday, nil
	case csvDayWednesday, csvDayShortWednesday:
		return DayWednesday, nil
	case csvDayThursday, csvDayShortThursday:
		return DayThursday, nil
	case csvDayFriday, csvDayShortFriday:
		return DayFriday, nil
	case csvDaySaturday, csvDayShortSaturday:
		return DaySaturday, nil
	case csvDaySunday, csvDayShortSunday:
		return DaySunday, nil
	default:
		return "", ErrCSVDay
	}
}
