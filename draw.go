package lotto

import (
	"sort"
	"time"
)

// Order type
const (
	OrderASC OrderType = iota
	OrderDESC
)

// DrawType
const (
	DrawSuperLottoType DrawType = "super-lotto"
	DrawGrandLottoType DrawType = "grand-lotto"
	DrawXmasLottoType  DrawType = "xmas-lotto"
	DrawLottoType      DrawType = "new-lotto"
)

// DrawVersion v0 for the oldest
const (
	DrawV0 DrawVersion = "v0"
	DrawV1 DrawVersion = "v1"
	DrawV2 DrawVersion = "v2"
	DrawV3 DrawVersion = "v3"
	DrawV4 DrawVersion = "v4"
)

// Day generic value
const (
	DayMonday    Day = "MONDAY"
	DayTuesday   Day = "TUESDAY"
	DayWednesday Day = "WEDNESDAY"
	DayThursday  Day = "THURSDAY"
	DayFriday    Day = "FRIDAY"
	DaySaturday  Day = "SATURDAY"
	DaySunday    Day = "SUNDAY"
)

const (
	CurrencyEur Currency = "EUR"
)

type OrderType uint8

type DrawRecorder interface {
	ConvertDraw(drawType DrawType) (Draw, error)
}

type Day string
type DrawVersion string
type DrawType string

type Currency string

type Draw struct {
	Metadata Metadata
	DrawBall DrawBall
	Joker    Joker
	WinStats WinStats
	WinCode  WinCode
}

type WinCode struct {
	Number int
	Codes  []string
	Price  float64
}

type WinStats struct {
	// WinNumber present number of winner by Rank. Index 0 is the Rank1.
	WinNumber []int32
	// WinRate present amount money win by each winner. Index 0 is the Rank1.
	WinRate []float64

	SecondRoll bool
	// WinNumberSecondRoll present number of winner by Rank for the second roll. Index 0 is the Rank1.
	WinNumberSecondRoll []int32
	// WinRateSecondRoll present amount money win by each winner for the second roll. Index 0 is the Rank1.
	WinRateSecondRoll []float64
}

type Joker struct {
	Base string
	Plus string
}

type DrawBall struct {
	Balls           []int32
	LuckyBall       int32
	BallsSecondRoll []int32
}

type Metadata struct {
	Version        DrawVersion
	OldType        bool
	DrawType       DrawType
	TirageOrder    int32
	ID             string
	Date           time.Time
	ForclosureDate time.Time
	Day            Day
	Currency       Currency
}

func OrderDraws(draws *[]Draw, order OrderType) {
	if order == OrderASC {
		orderDrawsASC(draws)
	} else {
		orderDrawsDESC(draws)
	}
}

// orderDrawsASC from less to more
// order from recentest to oldest
//
//nolint:dupl // no duplicate with orderDrawsDESC
func orderDrawsASC(draws *[]Draw) {
	sort.SliceStable(*draws, func(i, j int) bool {
		if (*draws)[i].Metadata.Date.After((*draws)[j].Metadata.Date) {
			return true
		}
		if (*draws)[i].Metadata.Date.Equal((*draws)[j].Metadata.Date) &&
			(*draws)[i].Metadata.TirageOrder > (*draws)[j].Metadata.TirageOrder {
			return true
		}

		return false
	})
}

// orderDrawsDESC from more to less
// order from oldest to recentest
//
//nolint:dupl // no duplicate with orderDrawsASC
func orderDrawsDESC(draws *[]Draw) {
	sort.SliceStable(*draws, func(i, j int) bool {
		if (*draws)[i].Metadata.Date.Before((*draws)[j].Metadata.Date) {
			return true
		}
		if (*draws)[i].Metadata.Date.Equal((*draws)[j].Metadata.Date) &&
			(*draws)[i].Metadata.TirageOrder < (*draws)[j].Metadata.TirageOrder {
			return true
		}

		return false
	})
}
