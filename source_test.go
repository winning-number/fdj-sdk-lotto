package lotto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winning-number/fdj-sdk-lotto/draw"
)

func TestSourceInfo_URL(t *testing.T) {
	t.Run("Should return the url for the source 'Grand lotto'", func(t *testing.T) {
		expected := "https://media.fdj.fr/static/csv/loto/grandloto_201912.zip"

		got := GetSourceInfo(GrandLoto).URL()

		assert.EqualValues(t, expected, got)
	})
}

func TestGetSourceInfo(t *testing.T) {
	t.Run("Should return the source info 'Grand lotto'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: GrandLotoZipName,
			Type:       draw.GrandLottoType,
			Version:    draw.V3,
			Name:       GrandLoto,
		}

		got := GetSourceInfo(GrandLoto)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Grand lotto noel'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: GrandLotoNoelZipName,
			Type:       draw.XmasLottoType,
			Version:    draw.V3,
			Name:       GrandLotoNoel,
		}

		got := GetSourceInfo(GrandLotoNoel)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Super lotto 1996-05'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: SuperLoto199605ZipName,
			Type:       draw.SuperLottoType,
			Version:    draw.V0,
			Name:       SuperLoto199605,
		}

		got := GetSourceInfo(SuperLoto199605)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Super lotto 2008-10'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: SuperLoto200810ZipName,
			Type:       draw.SuperLottoType,
			Version:    draw.V2,
			Name:       SuperLoto200810,
		}

		got := GetSourceInfo(SuperLoto200810)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Super lotto 2017-03'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: SuperLoto201703ZipName,
			Type:       draw.SuperLottoType,
			Version:    draw.V3,
			Name:       SuperLoto201703,
		}

		got := GetSourceInfo(SuperLoto201703)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Super lotto 2019-07'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: SuperLoto201907ZipName,
			Type:       draw.SuperLottoType,
			Version:    draw.V3,
			Name:       SuperLoto201907,
		}

		got := GetSourceInfo(SuperLoto201907)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Lotto 1976-05'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: Loto197605ZipName,
			Type:       draw.LottoType,
			Version:    draw.V1,
			Name:       Loto197605,
		}

		got := GetSourceInfo(Loto197605)

		assert.EqualValues(t, expected, got)
	})

	t.Run("Should return the source info 'Lotto 2008-10'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: Loto200810ZipName,
			Type:       draw.LottoType,
			Version:    draw.V2,
			Name:       Loto200810,
		}

		got := GetSourceInfo(Loto200810)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Lotto 2017-03'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: Loto201703ZipName,
			Type:       draw.LottoType,
			Version:    draw.V3,
			Name:       Loto201703,
		}

		got := GetSourceInfo(Loto201703)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Lotto 2019-02'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: Loto201902ZipName,
			Type:       draw.LottoType,
			Version:    draw.V3,
			Name:       Loto201902,
		}

		got := GetSourceInfo(Loto201902)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the source info 'Lotto 20019-11'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: Loto201911ZipName,
			Type:       draw.LottoType,
			Version:    draw.V4,
			Name:       Loto201911,
		}

		got := GetSourceInfo(Loto201911)

		assert.EqualValues(t, expected, got)
	})
	t.Run("Should return the default source info 'unknow source'", func(t *testing.T) {
		expected := SourceInfo{
			APIZipName: Loto201911ZipName,
			Type:       draw.LottoType,
			Version:    draw.V4,
			Name:       Loto201911,
		}

		got := GetSourceInfo(9999)

		assert.EqualValues(t, expected, got)
	})
}

func TestSourceInfoAll(t *testing.T) {
	t.Run("Should return all the sources", func(t *testing.T) {
		infos := SourceInfoAll()

		assert.Len(t, infos, 11)
	})
}

func TestSourceAll(t *testing.T) {
	t.Run("Should return all the sources", func(t *testing.T) {
		expected := []Source{
			GrandLoto, GrandLotoNoel, SuperLoto199605, SuperLoto200810, SuperLoto201703, SuperLoto201907,
			Loto197605, Loto200810, Loto201703, Loto201902, Loto201911,
		}

		got := SourceAll()

		assert.EqualValues(t, expected, got)
	})
}
