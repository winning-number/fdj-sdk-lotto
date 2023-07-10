package lotto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winning-number/fdj-sdk-lotto/draw"
)

func TestFilter_Match(t *testing.T) {
	t.Run("Should not match with day ", func(t *testing.T) {
		filter := Filter{
			Day: "bad-day",
		}
		draw := draw.Draw{Metadata: draw.Metadata{
			Day: draw.DayMonday,
		}}
		ok := filter.Match(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match with old lotto", func(t *testing.T) {
		filter := Filter{
			OldLotto: false,
		}
		draw := draw.Draw{Metadata: draw.Metadata{
			Day:     draw.DayMonday,
			OldType: true,
		}}
		ok := filter.Match(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match with the super lotto type", func(t *testing.T) {
		filter := Filter{
			SuperLotto: false,
		}
		draw := draw.Draw{Metadata: draw.Metadata{
			Day:      draw.DayMonday,
			DrawType: draw.SuperLottoType,
		}}
		ok := filter.Match(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match with the grand lotto type", func(t *testing.T) {
		filter := Filter{
			GrandLotto: false,
		}
		draw := draw.Draw{Metadata: draw.Metadata{
			Day:      draw.DayMonday,
			DrawType: draw.GrandLottoType,
		}}
		ok := filter.Match(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match with the xmas lotto type", func(t *testing.T) {
		filter := Filter{
			XmasLotto: false,
		}
		draw := draw.Draw{Metadata: draw.Metadata{
			Day:      draw.DayMonday,
			DrawType: draw.XmasLottoType,
		}}
		ok := filter.Match(&draw)
		assert.False(t, ok)
	})
	t.Run("Should not match with the classic lotto type", func(t *testing.T) {
		filter := Filter{
			ClassicLotto: false,
		}
		draw := draw.Draw{Metadata: draw.Metadata{
			Day:      draw.DayMonday,
			DrawType: draw.LottoType,
		}}
		ok := filter.Match(&draw)
		assert.False(t, ok)
	})
	t.Run("Should match the draw", func(t *testing.T) {
		filter := Filter{
			Day:          draw.DayMonday,
			ClassicLotto: true,
			OldLotto:     true,
		}
		draw := draw.Draw{Metadata: draw.Metadata{
			Day:      draw.DayMonday,
			DrawType: draw.LottoType,
			OldType:  true,
		}}
		ok := filter.Match(&draw)
		assert.True(t, ok)
	})
}
