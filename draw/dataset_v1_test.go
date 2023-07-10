package draw

import (
	"time"
)

func DataSetClassicLottoV1() []Draw {
	return []Draw{
		{
			Metadata: Metadata{
				Version:        V1,
				OldType:        true,
				DrawType:       LottoType,
				TirageOrder:    2,
				FDJID:          "2008080",
				ID:             "2008080215-16-32-33-42-49",
				Date:           time.Date(2008, time.October, 4, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2008, time.December, 4, 0, 0, 0, 0, time.UTC),
				Day:            DaySaturday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{33, 32, 42, 16, 15, 49, 37},
				HasLucky:  false,
				HasSecond: false,
			},
			Joker: Joker{
				HasBase: true,
				Base:    "",
				Plus:    "7 523 262",
			},
			WinStats: WinStats{
				WinNumber:           []int32{2, 8, 213, 589, 11911, 20552, 255211},
				WinRate:             []float64{904940, 10953.9, 1400.3, 61.6, 30.8, 5.4, 2.7},
				SecondRoll:          false,
				WinNumberSecondRoll: []int32(nil),
				WinRateSecondRoll:   []float64(nil),
			},
		},
		{
			Metadata: Metadata{
				Version:        V1,
				OldType:        true,
				DrawType:       LottoType,
				TirageOrder:    1,
				FDJID:          "2008080",
				ID:             "200808019-12-16-25-27-36",
				Date:           time.Date(2008, time.October, 4, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2008, time.December, 4, 0, 0, 0, 0, time.UTC),
				Day:            DaySaturday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{36, 9, 16, 27, 25, 12, 48},
				HasLucky:  false,
				HasSecond: false,
			},
			Joker: Joker{
				HasBase: true,
				Base:    "",
				Plus:    "7 523 262",
			},
			WinStats: WinStats{
				WinNumber:           []int32{3, 7, 461, 1033, 22997, 25950, 391755},
				WinRate:             []float64{282006, 12513.4, 662.5, 32, 16, 3.6, 1.8},
				SecondRoll:          false,
				WinNumberSecondRoll: []int32(nil),
				WinRateSecondRoll:   []float64(nil),
			},
		},
	}
}
