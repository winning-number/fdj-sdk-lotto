package draw

import (
	"time"
)

func DataSetSuperLottoV2() []Draw {
	return []Draw{
		{
			Metadata: Metadata{
				Version:        V2,
				OldType:        false,
				DrawType:       SuperLottoType,
				TirageOrder:    1,
				ID:             "2017001",
				Date:           time.Date(2017, time.January, 13, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2017, time.March, 15, 0, 0, 0, 0, time.UTC),
				Day:            DayFriday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{10, 12, 30, 24, 6},
				HasLucky:  true,
				HasSecond: false,
				LuckyBall: 2,
			},
			Joker: Joker{
				Base: "",
				Plus: "5 163 249",
			},
			WinStats: WinStats{
				WinNumber:           []int32{0, 14, 2970, 89335, 1014094, 1162890},
				WinRate:             []float64{0, 47971, 486.6, 7, 4.4, 2},
				SecondRoll:          false,
				WinNumberSecondRoll: []int32(nil),
				WinRateSecondRoll:   []float64(nil),
			},
			WinCode: WinCode{
				Number: 0,
				Codes:  []string(nil),
				Price:  0,
			},
		},
		{
			Metadata: Metadata{
				Version:        V2,
				OldType:        false,
				DrawType:       SuperLottoType,
				TirageOrder:    1,
				ID:             "2016002",
				Date:           time.Date(2016, time.September, 16, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2016, time.November, 16, 0, 0, 0, 0, time.UTC),
				Day:            DayFriday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{23, 31, 14, 7, 36},
				LuckyBall: 2,
				HasLucky:  true,
				HasSecond: false,
			},
			Joker: Joker{
				Base: "",
				Plus: "3 872 999",
			},
			WinStats: WinStats{
				WinNumber:           []int32{1, 2, 1158, 48738, 670731, 851371},
				WinRate:             []float64{13550865.0, 251668.5, 935.4, 9.6, 5, 2},
				SecondRoll:          false,
				WinNumberSecondRoll: []int32(nil),
				WinRateSecondRoll:   []float64(nil),
			},
			WinCode: WinCode{
				Number: 0,
				Codes:  []string(nil),
				Price:  0,
			},
		},
	}
}

func DataSetClassicLottoV2() []Draw {
	return []Draw{
		{
			Metadata: Metadata{
				Version:        V2,
				OldType:        false,
				DrawType:       LottoType,
				TirageOrder:    1,
				FDJID:          "2017027",
				ID:             "201702714-14-28-32-37+4",
				Date:           time.Date(2017, time.March, 4, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2017, time.May, 4, 0, 0, 0, 0, time.UTC),
				Day:            DaySaturday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{28, 14, 37, 32, 4},
				LuckyBall: 4,
				HasLucky:  true,
				HasSecond: false,
			},
			Joker: Joker{Base: "", Plus: "6 036 389"},
			WinStats: WinStats{
				WinNumber:           []int32{0, 1, 617, 26910, 394534, 585454},
				WinRate:             []float64{0, 327742.5, 1143.2, 11.3, 5.5, 2},
				SecondRoll:          false,
				WinNumberSecondRoll: []int32(nil),
				WinRateSecondRoll:   []float64(nil),
			},
			WinCode: WinCode{Number: 0, Codes: []string(nil), Price: 0},
		},
		{
			Metadata: Metadata{
				Version:        V2,
				OldType:        false,
				DrawType:       LottoType,
				TirageOrder:    1,
				ID:             "2017026116-20-31-33-46+5",
				FDJID:          "2017026",
				Date:           time.Date(2017, time.March, 1, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2017, time.May, 1, 0, 0, 0, 0, time.UTC),
				Day:            DayWednesday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{33, 31, 16, 46, 20},
				LuckyBall: 5,
				HasLucky:  true,
				HasSecond: false,
			},
			Joker: Joker{Base: "", Plus: "4 447 172"},
			WinStats: WinStats{
				WinNumber:           []int32{0, 0, 359, 17202, 265105, 494834},
				WinRate:             []float64{0, 0, 2077.1, 12.8, 5.9, 2},
				SecondRoll:          false,
				WinNumberSecondRoll: []int32(nil),
				WinRateSecondRoll:   []float64(nil),
			},
			WinCode: WinCode{Number: 0, Codes: []string(nil), Price: 0},
		},
	}
}
