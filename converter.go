package lotto

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const bitSize = 64

func DateFormat(separator string, date string, revert bool) (time.Time, error) {
	var dateTime time.Time
	var err error

	format := fmt.Sprintf("2006%s01%s02", separator, separator)
	if revert {
		format = fmt.Sprintf("02%s01%s2006", separator, separator)
	}
	if dateTime, err = time.Parse(format, date); err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}

// Date converter value
// Day generic value
const (
	dayConverterMonday         = "LUNDI   "
	dayConverterTuesday        = "MARDI   "
	dayConverterWednesday      = "MERCREDI"
	dayConverterThursday       = "JEUDI   "
	dayConverterFriday         = "VENDREDI"
	dayConverterSaturday       = "SAMEDI  "
	dayConverterSunday         = "DIMANCHE"
	dayConverterShortMonday    = "LU"
	dayConverterShortTuesday   = "MA"
	dayConverterShortWednesday = "ME"
	dayConverterShortThursday  = "JE"
	dayConverterShortFriday    = "VE"
	dayConverterShortSaturday  = "SA"
	dayConverterShortSunday    = "DI"
)

func DayConverter(day string) (Day, error) {
	switch day {
	case dayConverterMonday, dayConverterShortMonday:
		return DayMonday, nil
	case dayConverterTuesday, dayConverterShortTuesday:
		return DayTuesday, nil
	case dayConverterWednesday, dayConverterShortWednesday:
		return DayWednesday, nil
	case dayConverterThursday, dayConverterShortThursday:
		return DayThursday, nil
	case dayConverterFriday, dayConverterShortFriday:
		return DayFriday, nil
	case dayConverterSaturday, dayConverterShortSaturday:
		return DaySaturday, nil
	case dayConverterSunday, dayConverterShortSunday:
		return DaySunday, nil
	default:
		return "", errors.New("day type unknown")
	}
}

func AnyMoneyConverter(inputs ...string) ([]float64, error) {
	var ret []float64
	var err error

	ret = make([]float64, len(inputs))
	for i, input := range inputs {
		var v float64
		if v, err = MoneyConverter(input); err != nil {
			return nil, err
		}
		ret[i] = v
	}

	return ret, nil
}

func MoneyConverter(input string) (float64, error) {
	var ret float64
	var err error

	if input == "" {
		return 0.0, nil
	}
	input = strings.Replace(input, ",", ".", 1)
	if ret, err = strconv.ParseFloat(input, bitSize); err != nil {
		return 0.0, err
	}

	return ret, nil
}

func DrawTypeConverter(name string) (DrawType, error) {
	switch name {
	case string(DrawGrandLottoType):
		return DrawGrandLottoType, nil
	case string(DrawXmasType):
		return DrawXmasType, nil
	case string(DrawSuperLottoType):
		return DrawSuperLottoType, nil
	case string(DrawLottoType):
		return DrawLottoType, nil
	default:
		return "", ErrDrawTypeDecode
	}
}
