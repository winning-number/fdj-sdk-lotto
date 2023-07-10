package draw

import (
	"time"
)

func DataSetClassicLottoV4() []Draw {
	return []Draw{
		{
			Metadata: Metadata{
				Version:        V4,
				OldType:        false,
				DrawType:       LottoType,
				TirageOrder:    1,
				ID:             "2022209315-10-32-40-45+5",
				FDJID:          "20222093",
				Date:           time.Date(2022, time.August, 3, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2022, time.October, 3, 0, 0, 0, 0, time.UTC),
				Day:            DayWednesday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{45, 10, 5, 40, 32},
				LuckyBall: 5,
				HasLucky:  true,
				HasSecond: true,
				Second:    []int32{10, 14, 23, 28, 31},
			},
			Joker: Joker{
				Base: "",
				Plus: "Jokerplus indisponible",
			},
			WinStats: WinStats{
				WinNumber: []int32{0, 6, 39, 264, 1961, 12757, 29242, 200413, 394977},
				WinRate: []float64{
					3000000.0,
					26992.4,
					1013.5,
					540,
					43.4,
					24,
					8.1,
					4.7,
					2.2,
				},
				SecondRoll:          true,
				WinNumberSecondRoll: []int32{3, 366, 12817, 160661},
				WinRateSecondRoll:   []float64{39016, 286.7, 21, 3},
			},
			WinCode: WinCode{
				Number: 10,
				Codes: []string{
					"M 7538 5852",
					"I 4633 1764",
					"F 2366 8173",
					"Q 2074 2014",
					"S 6907 0682",
					"A 5138 7991",
					"N 0539 0042",
					"B 4326 8515",
					"E 3096 8757",
					"P 7093 9983",
				},
				Price: 20000,
			},
		},
		{
			Metadata: Metadata{
				Version:        V4,
				OldType:        false,
				DrawType:       LottoType,
				TirageOrder:    1,
				ID:             "2022209211-4-23-30-38+6",
				FDJID:          "20222092",
				Date:           time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
				ForclosureDate: time.Date(2022, time.October, 1, 0, 0, 0, 0, time.UTC),
				Day:            DayMonday,
				Currency:       CurrencyEur,
			},
			Roll: Roll{
				First:     []int32{38, 1, 4, 23, 30},
				LuckyBall: 6,
				HasLucky:  true,
				HasSecond: true,
				Second:    []int32{15, 24, 36, 45, 47},
			},
			Joker: Joker{
				Base: "",
				Plus: "8 724 758",
			},
			WinStats: WinStats{
				WinNumber: []int32{0, 2, 25, 278, 1267, 12848, 19210, 189532, 257794},
				WinRate: []float64{
					2000000.0,
					72206.7,
					1409.9,
					457.3,
					59.9,
					21.2,
					11,
					4.4,
					2.2,
				},
				SecondRoll:          true,
				WinNumberSecondRoll: []int32{0, 155, 6784, 102720},
				WinRateSecondRoll:   []float64{100000, 754, 44.3, 3},
			},
			WinCode: WinCode{
				Number: 10,
				Codes: []string{
					"C 2091 3108",
					"D 4468 6572",
					"S 6907 2812",
					"F 2177 2238",
					"U 9469 7275",
					"M 2284 3713",
					"K 8911 1530",
					"L 9410 9846",
					"V 2671 8789",
					"C 7186 4533",
				},
				Price: 20000,
			},
		},
	}
}
