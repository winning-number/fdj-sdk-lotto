package lotto

type Filter struct {
	Day Day

	SuperLotto   bool
	GrandLotto   bool
	XmasLotto    bool
	ClassicLotto bool

	OldLotto bool
}

// MatchWithDraw valid the draw parameter with the Filter.
func (o Filter) MatchWithDraw(draw *Draw) bool {
	if !o.matchDay(draw.Metadata.Day) {
		return false
	}
	if !o.matchOldLotto(draw.Metadata) {
		return false
	}
	if !o.matchDrawType(draw.Metadata) {
		return false
	}

	return true
}

func (o Filter) matchDay(day Day) bool {
	if string(o.Day) != "" && o.Day != day {
		return false
	}

	return true
}

func (o Filter) matchOldLotto(meta Metadata) bool {
	if !o.OldLotto && meta.OldType {
		return false
	}

	return true
}

func (o Filter) matchDrawType(meta Metadata) bool {
	switch meta.DrawType {
	case DrawSuperLottoType:
		if !o.SuperLotto {
			return false
		}
	case DrawGrandLottoType:
		if !o.GrandLotto {
			return false
		}
	case DrawXmasLottoType:
		if !o.XmasLotto {
			return false
		}
	case DrawLottoType:
		if !o.ClassicLotto {
			return false
		}
	}

	return true
}
