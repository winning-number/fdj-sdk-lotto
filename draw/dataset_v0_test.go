package draw

import (
	"time"
)

func DataSetSuperLottoV0() []Draw {
	return []Draw{
		{
			Metadata: Metadata{
				Version:        V0,
				OldType:        true,
				DrawType:       SuperLottoType,
				TirageOrder:    1,
				ID:             "2008002115-25-34-36-37-42",
				FDJID:          "2008002",
				Date:           time.Date(2008, time.June, 13, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2008, time.August, 13, 0, 0, 0, 0, time.UTC),
				Day:            DayFriday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{34, 25, 37, 42, 15, 36, 49},
				HasLucky:  false,
				HasSecond: false,
			},
			Joker: Joker{
				Base: "",
				Plus: "8 651 700",
			},
			WinStats: WinStats{
				WinNumber:           []int32{2, 6, 225, 613, 14049, 17651, 253160},
				WinRate:             []float64{7501991.0, 78102.6, 5101.4, 242.4, 121.2, 16.2, 8.1},
				SecondRoll:          false,
				WinNumberSecondRoll: []int32(nil),
				WinRateSecondRoll:   []float64(nil),
			},
		},
		{
			Metadata: Metadata{
				Version:        V0,
				OldType:        true,
				DrawType:       SuperLottoType,
				TirageOrder:    1,
				FDJID:          "2008001",
				ID:             "200800111-13-15-17-30-49",
				Date:           time.Date(2008, time.February, 29, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2008, time.April, 30, 0, 0, 0, 0, time.UTC),
				Day:            DayFriday,
				Currency:       CurrencyFR,
			},
			Roll: Roll{
				First:     []int32{1, 49, 30, 17, 15, 13, 48},
				HasLucky:  false,
				HasSecond: false,
			},
			Joker: Joker{
				Base: "",
				Plus: "8 152 912",
			},
			WinStats: WinStats{
				WinNumber:           []int32{2, 8, 351, 574, 16166, 15033, 259505},
				WinRate:             []float64{10001036.0, 44953.9, 2493.7, 161.8, 80.9, 12.2, 6.1},
				SecondRoll:          false,
				WinNumberSecondRoll: []int32(nil),
				WinRateSecondRoll:   []float64(nil),
			},
		},
	}
}
