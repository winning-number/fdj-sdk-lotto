package lotto

import "github.com/winning-number/fdj-sdk-lotto/draw"

// Filter is the filter of draw.
// if day is empty, it will match all day.
// if old lotto is false, it will not match old lotto, otherwise it will match all.
// if super lotto is false, it will not match super lotto.
// if grand lotto is false, it will not match grand lotto.
// if xmas lotto is false, it will not match xmas lotto.
// if classic lotto is false, it will not match classic lotto.
// classic lotto is the default draw type (from november 2019).
type Filter struct {
	Day draw.Day

	SuperLotto   bool
	GrandLotto   bool
	XmasLotto    bool
	ClassicLotto bool

	OldLotto bool
}

// Match returns true if the draw match the filter.
func (f Filter) Match(draw *draw.Draw) bool {
	if !f.matchDay(draw.Metadata.Day) {
		return false
	}
	if !f.matchOldLotto(draw.Metadata) {
		return false
	}
	if !f.matchDrawType(draw.Metadata) {
		return false
	}

	return true
}

func (f Filter) matchDay(day draw.Day) bool {
	if string(f.Day) != "" && f.Day != day {
		return false
	}

	return true
}

func (f Filter) matchOldLotto(meta draw.Metadata) bool {
	if !f.OldLotto && meta.OldType {
		return false
	}

	return true
}

func (f Filter) matchDrawType(meta draw.Metadata) bool {
	switch meta.DrawType {
	case draw.SuperLottoType:
		if !f.SuperLotto {
			return false
		}
	case draw.GrandLottoType:
		if !f.GrandLotto {
			return false
		}
	case draw.XmasLottoType:
		if !f.XmasLotto {
			return false
		}
	case draw.LottoType:
		if !f.ClassicLotto {
			return false
		}
	}

	return true
}
