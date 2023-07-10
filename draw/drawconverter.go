package draw

// DrawConverter is an interface to convert a specific draw version / type to a Draw.
type DrawConverter interface {
	winCode() (WinCode, error)
	winStat() (WinStats, error)
	joker() Joker
	roll() Roll
	metadata(drawT Type) (Metadata, error)
}

// Convert converts a DrawConverter to a Draw.
// It returns an error if the conversion fails.
func Convert[T DrawConverter](csv T, drawT Type) (Draw, error) {
	var draw Draw
	var err error

	if draw.Metadata, err = csv.metadata(drawT); err != nil {
		return Draw{}, err
	}
	draw.Roll = csv.roll()
	draw.Joker = csv.joker()
	if draw.WinStats, err = csv.winStat(); err != nil {
		return Draw{}, err
	}
	if draw.WinCode, err = csv.winCode(); err != nil {
		return Draw{}, err
	}

	return draw, nil
}

type coreCSV struct {
	ID             string `csv:"annee_numero_de_tirage"`
	Date           string `csv:"date_de_tirage"`
	ForclosureDate string `csv:"date_de_forclusion"`
	Day            string `csv:"jour_de_tirage"`
	Currency       string `csv:"devise"`
}

func (c coreCSV) metadata(drawT Type) (Metadata, error) {
	var meta Metadata
	var err error

	meta.DrawType = drawT
	meta.FDJID = c.ID
	if meta.Date, err = dateConverter(c.Date); err != nil {
		return Metadata{}, err
	}
	if meta.ForclosureDate, err = dateConverter(c.ForclosureDate); err != nil {
		return Metadata{}, err
	}
	if meta.Day, err = dayConverter(c.Day); err != nil {
		return Metadata{}, err
	}
	meta.Currency = currencyConverter(c.Currency)

	return meta, nil
}
