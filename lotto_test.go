package lotto

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/winning-number/fdj-sdk-lotto/csvparser"
	"github.com/winning-number/fdj-sdk-lotto/helptest"
)

// datatest files
const (
	superLottoTestFileV0     = "testdata/super-loto-v0.csv"
	superLottoTestFileV2     = "testdata/super-loto-v2.csv"
	superLottoTestFileV3     = "testdata/super-loto-v3.csv"
	grandLottoTestFileV3     = "testdata/grand-loto-v3.csv"
	grandLottoNoelTestFileV3 = "testdata/grand-loto-noel-v3.csv"
	classicLottoTestFileV1   = "testdata/classic-loto-v1.csv"
	classicLottoTestFileV2   = "testdata/classic-loto-v2.csv"
	classicLottoTestFileV3   = "testdata/classic-loto-v3.csv"
	classicLottoTestFileV4   = "testdata/classic-loto-v4.csv"
)

// helpLoadAPIOptionSourceDisabled return LoadAPIOption without source enable
func helpLoadAPIOptionSourceDisabled() LoadAPIOption {
	return LoadAPIOption{
		SourceDisable: LoadAPISourceDisable{
			GrandLoto:       true,
			GrandLotoNoel:   true,
			SuperLoto199605: true,
			SuperLoto200810: true,
			SuperLoto201703: true,
			SuperLoto201907: true,
			Loto197605:      true,
			Loto200810:      true,
			Loto201703:      true,
			Loto201902:      true,
			Loto201911:      true,
		},
	}
}

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
		f := helptest.OpenFile(t, superLottoTestFileV0)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawSuperLottoType, DrawV0)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, dataSuperLottoV0(), l.draws)
		}
	})
	t.Run("Should be ok for the classic lotto (draw v1)", func(t *testing.T) {
		f := helptest.OpenFile(t, classicLottoTestFileV1)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawLottoType, DrawV1)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, dataClassicLottoV1(), l.draws)
		}
	})
	t.Run("Should be ok for the super lotto (draw v2)", func(t *testing.T) {
		f := helptest.OpenFile(t, superLottoTestFileV2)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawSuperLottoType, DrawV2)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, dataSuperLottoV2(), l.draws)
		}
	})
	t.Run("Should be ok for the grand lotto (draw v3)", func(t *testing.T) {
		f := helptest.OpenFile(t, grandLottoTestFileV3)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawGrandLottoType, DrawV3)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, dataGrandLottoV3(), l.draws)
		}
	})
	t.Run("Should be ok for the classic lotto (draw v4)", func(t *testing.T) {
		f := helptest.OpenFile(t, classicLottoTestFileV4)
		defer f.Close()
		l := &lotto{}

		warn, err := l.LoadCSV(f, DrawLottoType, DrawV4)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, dataClassicLottoV4(), l.draws)
		}
	})
}

// assertLottoLoadAPI test one case of TestLotto_LoadAPI
func assertLottoLoadAPI(
	t *testing.T,
	option LoadAPIOption,
	expectedDraws []Draw,
	csvFile,
	pathHandler string) {

	buf, err := os.ReadFile(csvFile)
	if err != nil {
		t.Error(err)
	}
	body := helptest.CreateZipReader(t, helptest.ZipContent{Content: buf, Name: filepath.Base(csvFile)})
	url := fmt.Sprintf("%s/%s", BasePath, pathHandler)
	lotto := &lotto{
		httpClient: helptest.CreateFakeHTTPClient(t, body, url),
	}

	err = lotto.LoadAPI(option)
	if assert.NoError(t, err) {
		assert.EqualValues(t, expectedDraws, lotto.draws)
	}
}

func TestLotto_LoadAPI(t *testing.T) {
	t.Run("Should be ok for the superLoto199605", func(t *testing.T) {
		csvFile := superLottoTestFileV0
		pathHandler := SuperLoto199605
		expectedDraws := dataSuperLottoV0()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.SuperLoto199605 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the superLoto200810", func(t *testing.T) {
		csvFile := superLottoTestFileV2
		pathHandler := SuperLoto200810
		expectedDraws := dataSuperLottoV2()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.SuperLoto200810 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the superLoto201703", func(t *testing.T) {
		csvFile := superLottoTestFileV3
		pathHandler := SuperLoto201703
		expectedDraws := dataSuperLottoV3()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.SuperLoto201703 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the superLoto201907", func(t *testing.T) {
		csvFile := superLottoTestFileV3
		pathHandler := SuperLoto201907
		expectedDraws := dataSuperLottoV3()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.SuperLoto201907 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto197605", func(t *testing.T) {
		csvFile := classicLottoTestFileV1
		pathHandler := Loto197605
		expectedDraws := dataClassicLottoV1()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto197605 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto200810", func(t *testing.T) {
		csvFile := classicLottoTestFileV2
		pathHandler := Loto200810
		expectedDraws := dataClassicLottoV2()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto200810 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto201703", func(t *testing.T) {
		csvFile := classicLottoTestFileV3
		pathHandler := Loto201703
		expectedDraws := dataClassicLottoV3()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto201703 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto201902", func(t *testing.T) {
		csvFile := classicLottoTestFileV3
		pathHandler := Loto201902
		expectedDraws := dataClassicLottoV3()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto201902 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the Loto201911", func(t *testing.T) {
		csvFile := classicLottoTestFileV4
		pathHandler := Loto201911
		expectedDraws := dataClassicLottoV4()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.Loto201911 = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the GrandLoto", func(t *testing.T) {
		csvFile := grandLottoTestFileV3
		pathHandler := GrandLoto
		expectedDraws := dataGrandLottoV3()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.GrandLoto = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
	t.Run("Should be ok for the GrandLotoNoel", func(t *testing.T) {
		csvFile := grandLottoNoelTestFileV3
		pathHandler := GrandLotoNoel
		expectedDraws := dataGrandLottoNoelV3()
		option := helpLoadAPIOptionSourceDisabled()
		option.SourceDisable.GrandLotoNoel = false

		assertLottoLoadAPI(t, option, expectedDraws, csvFile, pathHandler)
	})
}

func TestLotto_DrawCount(t *testing.T) {
	driver := &lotto{}
	driver.draws = append(driver.draws, dataSuperLottoV0()...)
	driver.draws = append(driver.draws, dataSuperLottoV2()...)
	driver.draws = append(driver.draws, dataSuperLottoV3()...)
	driver.draws = append(driver.draws, dataClassicLottoV1()...)
	driver.draws = append(driver.draws, dataClassicLottoV2()...)
	driver.draws = append(driver.draws, dataClassicLottoV3()...)
	driver.draws = append(driver.draws, dataClassicLottoV4()...)
	driver.draws = append(driver.draws, dataGrandLottoNoelV3()...)
	driver.draws = append(driver.draws, dataGrandLottoV3()...)

	globalFilter := Filter{
		OldLotto:     true,
		ClassicLotto: true,
		XmasLotto:    true,
		GrandLotto:   true,
		SuperLotto:   true,
	}
	t.Run("Should find nothing which match with the filter", func(t *testing.T) {
		expectedNBDraws := 0
		nbDraw := driver.DrawCount(Filter{})

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Monday", func(t *testing.T) {
		expectedNBDraws := 1
		globalFilter.Day = DayMonday
		nbDraw := driver.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Tuesday", func(t *testing.T) {
		expectedNBDraws := 2
		globalFilter.Day = DayTuesday
		nbDraw := driver.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Wednesday", func(t *testing.T) {
		expectedNBDraws := 3
		globalFilter.Day = DayWednesday
		nbDraw := driver.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Thursday", func(t *testing.T) {
		expectedNBDraws := 0
		globalFilter.Day = DayThursday
		nbDraw := driver.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Friday", func(t *testing.T) {
		expectedNBDraws := 9
		globalFilter.Day = DayFriday
		nbDraw := driver.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Saturday", func(t *testing.T) {
		expectedNBDraws := 4
		globalFilter.Day = DaySaturday
		nbDraw := driver.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match only the draws picked on Sunday", func(t *testing.T) {
		expectedNBDraws := 0
		globalFilter.Day = DaySunday
		nbDraw := driver.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
	t.Run("Should match all", func(t *testing.T) {
		expectedNBDraws := len(driver.draws)
		globalFilter.Day = ""
		nbDraw := driver.DrawCount(globalFilter)

		assert.Equal(t, expectedNBDraws, nbDraw)
	})
}

func TestLotto_Draws(t *testing.T) {
	driver := &lotto{}
	driver.draws = append(driver.draws, dataSuperLottoV0()...)
	driver.draws = append(driver.draws, dataSuperLottoV2()...)
	driver.draws = append(driver.draws, dataSuperLottoV3()...)
	driver.draws = append(driver.draws, dataClassicLottoV1()...)
	driver.draws = append(driver.draws, dataClassicLottoV2()...)
	driver.draws = append(driver.draws, dataClassicLottoV3()...)
	driver.draws = append(driver.draws, dataClassicLottoV4()...)
	driver.draws = append(driver.draws, dataGrandLottoNoelV3()...)
	driver.draws = append(driver.draws, dataGrandLottoV3()...)

	t.Run("Should find nothing which match with the filter", func(t *testing.T) {
		expectedDraws := []Draw{}
		draws := driver.Draws(Filter{})

		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match all draws", func(t *testing.T) {
		expectedDraws := driver.draws
		draws := driver.Draws(Filter{
			SuperLotto:   true,
			ClassicLotto: true,
			OldLotto:     true,
			GrandLotto:   true,
			XmasLotto:    true,
		})

		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match only the new classic draws", func(t *testing.T) {
		expectedDraws := dataClassicLottoV2()
		expectedDraws = append(expectedDraws, dataClassicLottoV3()...)
		expectedDraws = append(expectedDraws, dataClassicLottoV4()...)
		draws := driver.Draws(Filter{
			ClassicLotto: true,
		})
		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match only the new super draws", func(t *testing.T) {
		expectedDraws := dataSuperLottoV2()
		expectedDraws = append(expectedDraws, dataSuperLottoV3()...)
		draws := driver.Draws(Filter{
			SuperLotto: true,
		})
		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match only the new grand lotto type draws", func(t *testing.T) {
		expectedDraws := dataGrandLottoV3()
		draws := driver.Draws(Filter{
			GrandLotto: true,
		})
		assert.ElementsMatch(t, expectedDraws, draws)
	})
	t.Run("Should match only the new xmax lotto type draws", func(t *testing.T) {
		expectedDraws := dataGrandLottoNoelV3()
		draws := driver.Draws(Filter{
			XmasLotto: true,
		})
		assert.ElementsMatch(t, expectedDraws, draws)
	})
}
