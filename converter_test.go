package lotto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateFormat(t *testing.T) {
	t.Run("Should return an error because the format is invalid", func(t *testing.T) {
		expectedErr := "parsing time \"20210213\" as \"2006-01-02\": cannot parse \"0213\" as \"-\""
		date, err := DateFormat("-", "20210213", false)

		if assert.Error(t, err) {
			assert.EqualError(t, err, expectedErr)
			assert.Empty(t, date)
		}
	})
	t.Run("Should be ok without separator", func(t *testing.T) {
		expectedDate := "2021-02-13 00:00:00 +0000 UTC"
		date, err := DateFormat("", "20210213", false)

		if assert.NoError(t, err) {
			assert.Equal(t, expectedDate, date.String())
		}
	})
	t.Run("Should be with a '/' separator", func(t *testing.T) {
		expectedDate := "2021-02-13 00:00:00 +0000 UTC"
		date, err := DateFormat("/", "13/02/2021", true)

		if assert.NoError(t, err) {
			assert.Equal(t, expectedDate, date.String())
		}
	})
}

func TestDayConverter(t *testing.T) {
	t.Run("Should return an error because the day type is unknow", func(t *testing.T) {
		day, err := DayConverter("Monday")

		if assert.Error(t, err) {
			assert.EqualError(t, err, ErrDayUnknown.Error())
			assert.EqualValues(t, day, "")
		}
	})
	t.Run("Should be ok with Monday (short and long", func(t *testing.T) {
		day, err := DayConverter(dayConverterMonday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayMonday)
		}
		day, err = DayConverter(dayConverterShortMonday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayMonday)
		}
	})
	t.Run("Should be ok with Tuesday (short and long", func(t *testing.T) {
		day, err := DayConverter(dayConverterTuesday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayTuesday)
		}
		day, err = DayConverter(dayConverterShortTuesday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayTuesday)
		}
	})
	t.Run("Should be ok with Wednesday (short and long", func(t *testing.T) {
		day, err := DayConverter(dayConverterWednesday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayWednesday)
		}
		day, err = DayConverter(dayConverterShortWednesday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayWednesday)
		}
	})
	t.Run("Should be ok with Thursday (short and long", func(t *testing.T) {
		day, err := DayConverter(dayConverterThursday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayThursday)
		}
		day, err = DayConverter(dayConverterShortThursday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayThursday)
		}
	})
	t.Run("Should be ok with Friday (short and long", func(t *testing.T) {
		day, err := DayConverter(dayConverterFriday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayFriday)
		}
		day, err = DayConverter(dayConverterShortFriday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DayFriday)
		}
	})
	t.Run("Should be ok with Saturday (short and long", func(t *testing.T) {
		day, err := DayConverter(dayConverterSaturday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DaySaturday)
		}
		day, err = DayConverter(dayConverterShortSaturday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DaySaturday)
		}
	})
	t.Run("Should be ok with Sunday (short and long", func(t *testing.T) {
		day, err := DayConverter(dayConverterSunday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DaySunday)
		}
		day, err = DayConverter(dayConverterShortSunday)
		if assert.NoError(t, err) {
			assert.EqualValues(t, day, DaySunday)
		}
	})
}

func TestMoneyConverter(t *testing.T) {
	t.Run("Should return an error because the input is empty", func(t *testing.T) {
		expectedValue := 0.0
		val, err := MoneyConverter("")

		if assert.Error(t, err) {
			assert.EqualError(t, err, ErrInvalidMoney.Error())
			assert.EqualValues(t, val, expectedValue)
		}
	})
	t.Run("Should return an error because the input is not convertible", func(t *testing.T) {
		expectedValue := 0.0
		val, err := MoneyConverter("eleven.dot.two")

		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrMoneyConverter)
			assert.EqualValues(t, val, expectedValue)
		}
	})
	t.Run("Should be ok with dot separator", func(t *testing.T) {
		expectedValue := 4.2
		val, err := MoneyConverter("4.2")

		if assert.NoError(t, err) {
			assert.EqualValues(t, val, expectedValue)
		}
	})
	t.Run("Should be ok with comma separator", func(t *testing.T) {
		expectedValue := 4.2
		val, err := MoneyConverter("4,2")

		if assert.NoError(t, err) {
			assert.EqualValues(t, val, expectedValue)
		}
	})
	t.Run("Should be ok without separator", func(t *testing.T) {
		expectedValue := 7501991.0
		val, err := MoneyConverter("7501991")

		if assert.NoError(t, err) {
			assert.EqualValues(t, val, expectedValue)
		}
	})
}

func TestAnyMoneyConverter(t *testing.T) {
	t.Run("Should return an error from MoneyConverter", func(t *testing.T) {
		values, err := AnyMoneyConverter("4.2", "four-dot-two")

		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrMoneyConverter)
			assert.Empty(t, values)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		expectedValue := []float64{4.2, 21.0}
		values, err := AnyMoneyConverter("4.2", "21")

		if assert.NoError(t, err) {
			assert.EqualValues(t, values, expectedValue)
		}
	})
}

func TestDrawTypeConverter(t *testing.T) {
	t.Run("Should return an error because type unknown", func(t *testing.T) {
		drawType, err := DrawTypeConverter("Draw")

		if assert.Error(t, err) {
			assert.EqualError(t, err, ErrDrawTypeDecode.Error())
			assert.EqualValues(t, "", drawType)
		}
	})
	t.Run("Should be ok with Grand loto", func(t *testing.T) {
		drawType, err := DrawTypeConverter(string(DrawGrandLottoType))

		if assert.NoError(t, err) {
			assert.EqualValues(t, DrawGrandLottoType, drawType)
		}
	})
	t.Run("Should be ok with Xmas loto", func(t *testing.T) {
		drawType, err := DrawTypeConverter(string(DrawXmasLottoType))

		if assert.NoError(t, err) {
			assert.EqualValues(t, DrawXmasLottoType, drawType)
		}
	})
	t.Run("Should be ok with Super loto", func(t *testing.T) {
		drawType, err := DrawTypeConverter(string(DrawSuperLottoType))

		if assert.NoError(t, err) {
			assert.EqualValues(t, DrawSuperLottoType, drawType)
		}
	})
	t.Run("Should be ok with classic loto", func(t *testing.T) {
		drawType, err := DrawTypeConverter(string(DrawLottoType))

		if assert.NoError(t, err) {
			assert.EqualValues(t, DrawLottoType, drawType)
		}
	})
}
