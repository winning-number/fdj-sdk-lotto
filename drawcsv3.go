package lotto

import "strings"

// DrawCSV3 define the model for the draws from 2017 march to 2019 february
// Compatible with the draws from 2019 february to 2019 november
type DrawCSV3 struct {
	ID             string `csv:"annee_numero_de_tirage"`
	Date           string `csv:"date_de_tirage"`
	ForclosureDate string `csv:"date_de_forclusion"`
	Day            string `csv:"jour_de_tirage"`
	B1             int32  `csv:"boule_1"`
	B2             int32  `csv:"boule_2"`
	B3             int32  `csv:"boule_3"`
	B4             int32  `csv:"boule_4"`
	B5             int32  `csv:"boule_5"`
	LuckyBall      int32  `csv:"numero_chance"`
	JokerPlus      string `csv:"numero_jokerplus"`
	WinOrder       string `csv:"combinaison_gagnante_en_ordre_croissant"`
	WinnerR1       int32  `csv:"nombre_de_gagnant_au_rang1"`
	WinnerR2       int32  `csv:"nombre_de_gagnant_au_rang2"`
	WinnerR3       int32  `csv:"nombre_de_gagnant_au_rang3"`
	WinnerR4       int32  `csv:"nombre_de_gagnant_au_rang4"`
	WinnerR5       int32  `csv:"nombre_de_gagnant_au_rang5"`
	WinnerR6       int32  `csv:"nombre_de_gagnant_au_rang6"`
	WinnerR7       int32  `csv:"nombre_de_gagnant_au_rang7"`
	WinnerR8       int32  `csv:"nombre_de_gagnant_au_rang8"`
	WinnerR9       int32  `csv:"nombre_de_gagnant_au_rang9"`
	GainR1         string `csv:"rapport_du_rang1"`
	GainR2         string `csv:"rapport_du_rang2"`
	GainR3         string `csv:"rapport_du_rang3"`
	GainR4         string `csv:"rapport_du_rang4"`
	GainR5         string `csv:"rapport_du_rang5"`
	GainR6         string `csv:"rapport_du_rang6"`
	GainR7         string `csv:"rapport_du_rang7"`
	GainR8         string `csv:"rapport_du_rang8"`
	GainR9         string `csv:"rapport_du_rang9"`
	GainCode       string `csv:"rapport_codes_gagnants"`
	Currency       string `csv:"devise"`
	NumberWinCodes int32  `csv:"nombre_de_codes_gagnants"`
	WinCodes       string `csv:"codes_gagnants"`
}

//nolint:dupl // match with csv4
func (d DrawCSV3) ConvertDraw(drawType DrawType) (Draw, error) {
	var draw Draw
	var err error

	if draw.Metadata, err = d.ConvertMetadata(drawType); err != nil {
		return Draw{}, err
	}
	draw.DrawBall = d.ConvertDrawBalls()
	draw.Joker = d.ConvertJoker()
	if draw.WinStats, err = d.ConvertWinStats(); err != nil {
		return Draw{}, err
	}
	if draw.WinCode, err = d.ConvertWinCode(); err != nil {
		return Draw{}, err
	}

	return draw, nil
}

func (d DrawCSV3) ConvertWinCode() (WinCode, error) {
	var wCode WinCode
	var err error

	if wCode.Price, err = MoneyConverter(d.GainCode); err != nil {
		return WinCode{}, err
	}
	wCode.Number = int(d.NumberWinCodes)
	wCode.Codes = strings.Split(d.WinCodes, ",")

	return wCode, nil
}

func (d DrawCSV3) ConvertWinStats() (WinStats, error) {
	var err error
	wStat := WinStats{
		WinNumber: []int32{
			d.WinnerR1,
			d.WinnerR2,
			d.WinnerR3,
			d.WinnerR4,
			d.WinnerR5,
			d.WinnerR6,
			d.WinnerR7,
			d.WinnerR8,
			d.WinnerR9,
		},
		SecondRoll: false,
	}

	if wStat.WinRate, err = AnyMoneyConverter(
		d.GainR1,
		d.GainR2,
		d.GainR3,
		d.GainR4,
		d.GainR5,
		d.GainR6,
		d.GainR7,
		d.GainR8,
		d.GainR9,
	); err != nil {
		return WinStats{}, err
	}

	return wStat, nil
}

func (d DrawCSV3) ConvertJoker() Joker {
	return Joker{
		Plus: d.JokerPlus,
	}
}

func (d DrawCSV3) ConvertDrawBalls() DrawBall {
	var balls DrawBall

	balls.Balls = []int32{d.B1, d.B2, d.B3, d.B4, d.B5}
	balls.LuckyBall = d.LuckyBall

	return balls
}

//nolint:dupl // match with csv0/2/3/4
func (d DrawCSV3) ConvertMetadata(drawType DrawType) (Metadata, error) {
	var meta Metadata
	var err error

	meta.DrawType = drawType
	meta.OldType = false
	meta.Version = DrawV3
	meta.TirageOrder = 1
	meta.ID = d.ID
	if meta.Date, err = DateFormat("/", d.Date, true); err != nil {
		return Metadata{}, err
	}
	if meta.ForclosureDate, err = DateFormat("/", d.ForclosureDate, true); err != nil {
		return Metadata{}, err
	}
	if meta.Day, err = DayConverter(d.Day); err != nil {
		return Metadata{}, err
	}
	meta.Currency = CurrencyEur

	return meta, nil
}
