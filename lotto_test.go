package lotto

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/winning-number/fdjapi-lotto/csvparser"
)

func TestNew(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		l, err := New()

		if assert.NoError(t, err) {
			assert.NotNil(t, l)
		}
	})
}

func TestLotto_LoadCSV(t *testing.T) {
	t.Run("Should return an error because DrawVersion is invalid", func(t *testing.T) {
		l := &lotto{}

		warn, err := l.LoadCSV(nil, DrawSuperLottoType, "plop")
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrDrawVersion)
			assert.Nil(t, warn)
		}
	})
	t.Run("Should return an error reader is nil", func(t *testing.T) {
		l := &lotto{}

		warn, err := l.LoadCSV(nil, DrawSuperLottoType, DrawV0)
		if assert.Error(t, err) {
			assert.Nil(t, warn)
			assert.ErrorIs(t, err, csvparser.ErrNilReader)
		}
	})
	t.Run("Should be ok for the super lotto (draw v0)", func(t *testing.T) {
		f := helpOpenFile(t, superLottoTestFileV0)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawSuperLottoType, DrawV0)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, TestDataSuperLottoV0, l.draws)
		}
	})
	t.Run("Should be ok for the classic lotto (draw v1)", func(t *testing.T) {
		f := helpOpenFile(t, classicLottoTestFileV1)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawLottoType, DrawV1)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, TestDataClassicLottoV1, l.draws)
		}
	})
	t.Run("Should be ok for the super lotto (draw v2)", func(t *testing.T) {
		f := helpOpenFile(t, superLottoTestFileV2)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawSuperLottoType, DrawV2)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, TestDataSuperLottoV2, l.draws)
		}
	})
	t.Run("Should be ok for the grand lotto (draw v3)", func(t *testing.T) {
		f := helpOpenFile(t, grandLottoTestFileV3)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawGrandLottoType, DrawV3)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, TestDataGrandLottoV3, l.draws)
		}
	})
	t.Run("Should be ok for the classic lotto (draw v4)", func(t *testing.T) {
		f := helpOpenFile(t, classicLottoTestFileV4)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawLottoType, DrawV4)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, TestDataClassicLottoV4, l.draws)
		}
	})
}

// assertLotto_LoadAPI test one case of TestLotto_LoadAPI
func assertLotto_LoadAPI(
	t *testing.T,
	option LoadAPIOption,
	expectedDraws []Draw,
	csvFile,
	pathHandler string) {

	buf, err := os.ReadFile(csvFile)
	if err != nil {
		t.Error(err)
	}
	body := helpCreateZipReader(t, zipContent{content: buf, name: filepath.Base(csvFile)})
	l := helpCreateLottoWithFakeHTTPClient(t, body, pathHandler)

	err = l.LoadAPI(option)
	if assert.NoError(t, err) {
		assert.EqualValues(t, expectedDraws, l.draws)
	}
}

func TestLotto_LoadAPI(t *testing.T) {
	t.Run("Should be ok for the superLoto199605", func(t *testing.T) {
		csvFile := superLottoTestFileV0
		pathHandler := SuperLoto199605
		expectedDraws := TestDataSuperLottoV0
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.SuperLoto199605 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the superLoto200810", func(t *testing.T) {
		csvFile := superLottoTestFileV2
		pathHandler := SuperLoto200810
		expectedDraws := TestDataSuperLottoV2
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.SuperLoto200810 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the superLoto201703", func(t *testing.T) {
		csvFile := superLottoTestFileV3
		pathHandler := SuperLoto201703
		expectedDraws := TestDataSuperLottoV3
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.SuperLoto201703 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the superLoto201907", func(t *testing.T) {
		csvFile := superLottoTestFileV3
		pathHandler := SuperLoto201907
		expectedDraws := TestDataSuperLottoV3
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.SuperLoto201907 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto197605", func(t *testing.T) {
		csvFile := classicLottoTestFileV1
		pathHandler := Loto197605
		expectedDraws := TestDataClassicLottoV1
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto197605 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto200810", func(t *testing.T) {
		csvFile := classicLottoTestFileV2
		pathHandler := Loto200810
		expectedDraws := TestDataClassicLottoV2
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto200810 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto201703", func(t *testing.T) {
		csvFile := classicLottoTestFileV3
		pathHandler := Loto201703
		expectedDraws := TestDataClassicLottoV3
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto201703 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto201902", func(t *testing.T) {
		csvFile := classicLottoTestFileV3
		pathHandler := Loto201902
		expectedDraws := TestDataClassicLottoV3
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto201902 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto201911", func(t *testing.T) {
		csvFile := classicLottoTestFileV4
		pathHandler := Loto201911
		expectedDraws := TestDataClassicLottoV4
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto201911 = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the GrandLoto", func(t *testing.T) {
		csvFile := grandLottoTestFileV3
		pathHandler := GrandLoto
		expectedDraws := TestDataGrandLottoV3
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.GrandLoto = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the GrandLotoNoel", func(t *testing.T) {
		csvFile := grandLottoNoelTestFileV3
		pathHandler := GrandLotoNoel
		expectedDraws := TestDataGrandLottoNoelV3
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.GrandLotoNoel = false

		assertLotto_LoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
}

func TestLotto_DrawCount(t *testing.T) {
	l := &lotto{}
	l.draws = append(l.draws, TestDataSuperLottoV0...)
	l.draws = append(l.draws, TestDataSuperLottoV2...)
	l.draws = append(l.draws, TestDataSuperLottoV3...)
	l.draws = append(l.draws, TestDataClassicLottoV1...)
	l.draws = append(l.draws, TestDataClassicLottoV2...)
	l.draws = append(l.draws, TestDataClassicLottoV3...)
	l.draws = append(l.draws, TestDataClassicLottoV4...)
	l.draws = append(l.draws, TestDataGrandLottoNoelV3...)
	l.draws = append(l.draws, TestDataGrandLottoV3...)

	globalFilter := Filter{
		OldLotto:     true,
		ClassicLotto: true,
		XmasLotto:    true,
		GrandLotto:   true,
		SuperLotto:   true,
	}
	t.Run("Should find nothing which match with the filter", func(t *testing.T) {
		expectedNBDraws := 0
		nbDraw := l.DrawCount(Filter{})

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Monday", func(t *testing.T) {
		expectedNBDraws := 1
		globalFilter.Day = DayMonday
		nbDraw := l.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Tuesday", func(t *testing.T) {
		expectedNBDraws := 2
		globalFilter.Day = DayTuesday
		nbDraw := l.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Wednesday", func(t *testing.T) {
		expectedNBDraws := 3
		globalFilter.Day = DayWednesday
		nbDraw := l.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Thursday", func(t *testing.T) {
		expectedNBDraws := 0
		globalFilter.Day = DayThursday
		nbDraw := l.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Friday", func(t *testing.T) {
		expectedNBDraws := 9
		globalFilter.Day = DayFriday
		nbDraw := l.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Saturday", func(t *testing.T) {
		expectedNBDraws := 4
		globalFilter.Day = DaySaturday
		nbDraw := l.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Sunday", func(t *testing.T) {
		expectedNBDraws := 0
		globalFilter.Day = DaySunday
		nbDraw := l.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match all", func(t *testing.T) {
		expectedNBDraws := len(l.draws)
		globalFilter.Day = ""
		nbDraw := l.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
}

func TestLotto_Draws(t *testing.T) {
	l := &lotto{}
	l.draws = append(l.draws, TestDataSuperLottoV0...)
	l.draws = append(l.draws, TestDataSuperLottoV2...)
	l.draws = append(l.draws, TestDataSuperLottoV3...)
	l.draws = append(l.draws, TestDataClassicLottoV1...)
	l.draws = append(l.draws, TestDataClassicLottoV2...)
	l.draws = append(l.draws, TestDataClassicLottoV3...)
	l.draws = append(l.draws, TestDataClassicLottoV4...)
	l.draws = append(l.draws, TestDataGrandLottoNoelV3...)
	l.draws = append(l.draws, TestDataGrandLottoV3...)

	t.Run("Should find nothing which match with the filter", func(t *testing.T) {
		expectedDraws := []Draw{}
		draws := l.Draws(Filter{})

		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match all draws", func(t *testing.T) {
		expectedDraws := l.draws
		draws := l.Draws(Filter{
			SuperLotto:   true,
			ClassicLotto: true,
			OldLotto:     true,
			GrandLotto:   true,
			XmasLotto:    true,
		})

		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match only the new classic draws", func(t *testing.T) {
		expectedDraws := TestDataClassicLottoV2
		expectedDraws = append(expectedDraws, TestDataClassicLottoV3...)
		expectedDraws = append(expectedDraws, TestDataClassicLottoV4...)
		draws := l.Draws(Filter{
			ClassicLotto: true,
		})
		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match only the new super draws", func(t *testing.T) {
		expectedDraws := TestDataSuperLottoV2
		expectedDraws = append(expectedDraws, TestDataSuperLottoV3...)
		draws := l.Draws(Filter{
			SuperLotto: true,
		})
		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match only the new grand lotto type draws", func(t *testing.T) {
		expectedDraws := TestDataGrandLottoV3
		draws := l.Draws(Filter{
			GrandLotto: true,
		})
		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match only the new xmax lotto type draws", func(t *testing.T) {
		expectedDraws := TestDataGrandLottoNoelV3
		draws := l.Draws(Filter{
			XmasLotto: true,
		})
		assert.ElementsMatch(t, expectedDraws, draws)
	})
}
