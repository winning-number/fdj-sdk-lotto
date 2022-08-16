package lotto

type Filter struct {
	Day Day

	SuperLotto   bool
	GrandLotto   bool
	XmaxLotto    bool
	ClassicLotto bool

	OldLotto bool
}

// MatchWithDraw valid the draw parameter with the Filter.
func (o Filter) MatchWithDraw(draw *Draw) bool {
	if !o.MatchDay(draw.Metadata.Day) {
		return false
	}
	if !o.MatchOldLotto(draw.Metadata) {
		return false
	}
	if !o.MatchDrawType(draw.Metadata) {
		return false
	}

	return true
}

func (o Filter) MatchDay(day Day) bool {
	if string(o.Day) != "" && o.Day != day {
		return false
	}

	return true
}

func (o Filter) MatchOldLotto(meta Metadata) bool {
	if !o.OldLotto && meta.OldType {
		return false
	}

	return true
}

func (o Filter) MatchDrawType(meta Metadata) bool {
	switch meta.DrawType {
	case DrawSuperLottoType:
		if !o.SuperLotto {
			return false
		}
	case DrawGrandLottoType:
		if !o.GrandLotto {
			return false
		}
	case DrawXmasType:
		if !o.XmaxLotto {
			return false
		}
	case DrawLottoType:
		if !o.ClassicLotto {
			return false
		}
	}

	return true
}
