package draw

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winning-number/fdj-sdk-lotto/draw/testdata"
)

func TestCSV4_Convert(t *testing.T) {
	t.Run("Should convert csv4 to draw", func(t *testing.T) {
		data := make([]CSV4, 2)
		loadCSV(t, testdata.FileClassicLotoV4, &data)
		expected := DataSetClassicLottoV4()

		for i, csv := range data {
			draw, err := Convert(csv, LottoType)
			if assert.NoError(t, err) {
				assert.Equal(t, expected[i], draw)
			}
		}
	})
	t.Run("Should return an error on the first draw because metadata has failed", func(t *testing.T) {
		data := make([]CSV4, 2)
		loadCSV(t, testdata.FileClassicLotoV4, &data)
		expected := DataSetClassicLottoV4()

		data[0].Date = "invalid date"
		for i, csv := range data {
			draw, err := Convert(csv, LottoType)
			if i == 0 && assert.Error(t, err) {
				assert.ErrorIs(t, err, ErrCSVDate)
				assert.Empty(t, draw)
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, expected[i], draw)
				}
			}

		}
	})
	t.Run(
		"Should return an error on the first draw because winStat has failed on gain rank",
		func(t *testing.T) {
			data := make([]CSV4, 2)
			loadCSV(t, testdata.FileClassicLotoV4, &data)
			expected := DataSetClassicLottoV4()

			data[0].GainR1 = "invalid GainR1"
			for i, csv := range data {
				draw, err := Convert(csv, LottoType)
				if i == 0 && assert.Error(t, err) {
					assert.ErrorIs(t, err, ErrCSVPrice)
					assert.Empty(t, draw)
				} else {
					if assert.NoError(t, err) {
						assert.Equal(t, expected[i], draw)
					}
				}

			}
		})
	t.Run(
		"Should return an error on the first draw because winStat has failed on gain rank second roll",
		func(t *testing.T) {
			data := make([]CSV4, 2)
			loadCSV(t, testdata.FileClassicLotoV4, &data)
			expected := DataSetClassicLottoV4()

			data[0].GainR1SecondRoll = "invalid GainR1 Second Roll"
			for i, csv := range data {
				draw, err := Convert(csv, LottoType)
				if i == 0 && assert.Error(t, err) {
					assert.ErrorIs(t, err, ErrCSVPrice)
					assert.Empty(t, draw)
				} else {
					if assert.NoError(t, err) {
						assert.Equal(t, expected[i], draw)
					}
				}

			}
		})
	t.Run("Should return an error on the first draw because winCode has failed", func(t *testing.T) {
		data := make([]CSV4, 2)
		loadCSV(t, testdata.FileClassicLotoV4, &data)
		expected := DataSetClassicLottoV4()

		data[0].GainCode = "invalid Win1"
		for i, csv := range data {
			draw, err := Convert(csv, LottoType)
			if i == 0 && assert.Error(t, err) {
				assert.ErrorIs(t, err, ErrCSVPrice)
				assert.Empty(t, draw)
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, expected[i], draw)
				}
			}

		}
	})
}
