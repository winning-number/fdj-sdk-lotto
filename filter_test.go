package lotto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_MatchWithDraw(t *testing.T) {
	t.Run("Should not match the day", func(t *testing.T) {
		filter := Filter{
			Day: "bad-day",
		}
		draw := Draw{Metadata: Metadata{
			Day: DayMonday,
		}}
		ok := filter.MatchWithDraw(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match old type", func(t *testing.T) {
		filter := Filter{
			OldLotto: false,
		}
		draw := Draw{Metadata: Metadata{
			Day:     DayMonday,
			OldType: true,
		}}
		ok := filter.MatchWithDraw(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match super lotto type", func(t *testing.T) {
		filter := Filter{
			SuperLotto: false,
		}
		draw := Draw{Metadata: Metadata{
			Day:      DayMonday,
			DrawType: DrawSuperLottoType,
		}}
		ok := filter.MatchWithDraw(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match grand lotto type", func(t *testing.T) {
		filter := Filter{
			GrandLotto: false,
		}
		draw := Draw{Metadata: Metadata{
			Day:      DayMonday,
			DrawType: DrawGrandLottoType,
		}}
		ok := filter.MatchWithDraw(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match xmas lotto type", func(t *testing.T) {
		filter := Filter{
			XmasLotto: false,
		}
		draw := Draw{Metadata: Metadata{
			Day:      DayMonday,
			DrawType: DrawXmasLottoType,
		}}
		ok := filter.MatchWithDraw(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match classic lotto type", func(t *testing.T) {
		filter := Filter{
			ClassicLotto: false,
		}
		draw := Draw{Metadata: Metadata{
			Day:      DayMonday,
			DrawType: DrawLottoType,
		}}
		ok := filter.MatchWithDraw(&draw)
		assert.False(t, ok)
	})
	t.Run("Should match specific draw", func(t *testing.T) {
		filter := Filter{
			Day:          DayMonday,
			ClassicLotto: true,
			OldLotto:     true,
		}
		draw := Draw{Metadata: Metadata{
			Day:      DayMonday,
			DrawType: DrawLottoType,
			OldType:  true,
		}}
		ok := filter.MatchWithDraw(&draw)
		assert.True(t, ok)
	})
}
