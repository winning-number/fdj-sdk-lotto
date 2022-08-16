// Package lotto is package to get all lotto draws from the fdj history
// see https://www.fdj.fr/jeux-de-tirage/loto/historique
// Like an api, could get any archives to parse them into a common Draw object
// Filters could be get the draw by Day of draw, type of lotto ...
package lotto

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/winning-number/fdjapi-lotto/csvparser"
	"github.com/winning-number/fdjapi-lotto/httpclient"
)

// DrawTypeKey for the csvparser.ContextRecord
const (
	DrawTypeKey = "DrawType"
)

type LoadAPISourceDisable struct {
	GrandLoto       bool
	GrandLotoNoel   bool
	SuperLoto199605 bool
	SuperLoto200810 bool
	SuperLoto201703 bool
	SuperLoto201907 bool
	Loto197605      bool
	Loto200810      bool
	Loto201703      bool
	Loto201902      bool
	Loto201911      bool
}

type LoadAPIOption struct {
	SourceDisable LoadAPISourceDisable
	OutputZipFile string
	OutputCSVFile string
}

// Lotto interface could get the full history from the fdj archive or parse them directly from a csv
// LoadAPI and LoadCSV add draws inside the internal draws list and order them by date and id
type Lotto interface {
	// Read history from the FDJ api with filter to select the sources
	LoadAPI(option LoadAPIOption) error
	LoadCSV(filepath string, drawType DrawType, drawVersion DrawVersion) error

	DrawCount(filter Filter) int
	Draws(filter Filter) []Draw
}

type lotto struct {
	draws      []Draw
	httpClient *http.Client
}

// New Lotto instance collect the Draw history
func New() (Lotto, error) {
	var driver lotto

	driver.httpClient = httpclient.New()

	return &driver, nil
}

func (l *lotto) LoadAPI(option LoadAPIOption) error {
	var err error

	if err = l.loadingSources(option); err != nil {
		return err
	}
	l.orderDraws()

	return nil
}

func (l *lotto) LoadCSV(filepath string, drawType DrawType, drawVersion DrawVersion) error {
	var file *os.File
	var err error
	var warn csvparser.Warning

	conf := l.csvParseConfiguration(drawType)
	switch drawVersion {
	case DrawV0:
		conf.CreateObject = func() any { return &DrawCSV0{} }
	case DrawV1:
		conf.CreateObject = func() any { return &DrawCSV1{} }
	case DrawV2:
		conf.CreateObject = func() any { return &DrawCSV2{} }
	case DrawV3:
		conf.CreateObject = func() any { return &DrawCSV3{} }
	case DrawV4:
		conf.CreateObject = func() any { return &DrawCSV4{} }
	default:
		return ErrDrawVersion
	}

	if file, err = os.Open(filepath); err != nil {
		return err
	}
	defer file.Close()
	if warn, err = csvparser.CSVParse(file, conf); err != nil {
		return err
	}
	printWarnDecode(warn, filepath)
	l.orderDraws()

	return nil
}

func (l *lotto) DrawCount(filter Filter) int {
	draws := l.Draws(filter)

	return len(draws)
}

func (l lotto) Draws(f Filter) []Draw {
	matchesDraws := []Draw{}

	for i := range l.draws {
		if !f.MatchWithDraw(&l.draws[i]) {
			continue
		}
		matchesDraws = append(matchesDraws, l.draws[i])
	}

	return matchesDraws
}

func (l *lotto) orderDraws() {
	sort.SliceStable(l.draws, func(i, j int) bool {
		if l.draws[i].Metadata.Date.After(l.draws[j].Metadata.Date) {
			return true
		}
		if strings.Compare(l.draws[i].Metadata.ID, l.draws[j].Metadata.ID) == 0 &&
			l.draws[i].Metadata.TirageOrder > l.draws[j].Metadata.TirageOrder {
			return true
		}

		return false
	})
}

//nolint:funlen,gocognit,gocyclo // prefer keep the loading sources together
func (l *lotto) loadingSources(option LoadAPIOption) error {
	// ...
	var err error

	if !option.SourceDisable.GrandLoto {
		conf := l.csvParseConfiguration(DrawGrandLottoType)
		conf.CreateObject = func() any { return &DrawCSV3{} }
		if err = l.loadingSource(option, conf, GrandLoto); err != nil {
			return err
		}
	}
	if !option.SourceDisable.GrandLotoNoel {
		conf := l.csvParseConfiguration(DrawXmasType)
		conf.CreateObject = func() any { return &DrawCSV3{} }
		if err = l.loadingSource(option, conf, GrandLotoNoel); err != nil {
			return err
		}
	}
	if !option.SourceDisable.SuperLoto199605 {
		conf := l.csvParseConfiguration(DrawSuperLottoType)
		conf.CreateObject = func() any { return &DrawCSV0{} }
		if err = l.loadingSource(option, conf, SuperLoto199605); err != nil {
			return err
		}
	}
	if !option.SourceDisable.SuperLoto200810 {
		conf := l.csvParseConfiguration(DrawSuperLottoType)
		conf.CreateObject = func() any { return &DrawCSV2{} }
		if err = l.loadingSource(option, conf, SuperLoto200810); err != nil {
			return err
		}
	}
	if !option.SourceDisable.SuperLoto201703 {
		conf := l.csvParseConfiguration(DrawSuperLottoType)
		conf.CreateObject = func() any { return &DrawCSV3{} }
		if err = l.loadingSource(option, conf, SuperLoto201703); err != nil {
			return err
		}
	}
	if !option.SourceDisable.SuperLoto201907 {
		conf := l.csvParseConfiguration(DrawSuperLottoType)
		conf.CreateObject = func() any { return &DrawCSV3{} }
		if err = l.loadingSource(option, conf, SuperLoto201907); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto197605 {
		conf := l.csvParseConfiguration(DrawLottoType)
		conf.CreateObject = func() any { return &DrawCSV1{} }
		if err = l.loadingSource(option, conf, Loto197605); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto200810 {
		conf := l.csvParseConfiguration(DrawLottoType)
		conf.CreateObject = func() any { return &DrawCSV2{} }
		if err = l.loadingSource(option, conf, Loto200810); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto201703 {
		conf := l.csvParseConfiguration(DrawLottoType)
		conf.CreateObject = func() any { return &DrawCSV3{} }
		if err = l.loadingSource(option, conf, Loto201703); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto201902 {
		conf := l.csvParseConfiguration(DrawLottoType)
		conf.CreateObject = func() any { return &DrawCSV3{} }
		if err = l.loadingSource(option, conf, Loto201902); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto201911 {
		conf := l.csvParseConfiguration(DrawLottoType)
		conf.CreateObject = func() any { return &DrawCSV4{} }
		if err = l.loadingSource(option, conf, Loto201911); err != nil {
			return err
		}
	}

	return nil
}

func (l *lotto) loadingSource(option LoadAPIOption, conf csvparser.ParseConfig, filePath string) error {
	var source SourceReader
	var err error
	var warn csvparser.Warning

	if source, err = l.sourceReader(option, filePath); err != nil {
		return err
	}
	defer source.Close()
	if warn, err = csvparser.CSVParse(source.CSVReader(), conf); err != nil {
		return err
	}
	printWarnDecode(warn, filePath)

	return nil
}

func printWarnDecode(warn csvparser.Warning, source string) {
	if warn == nil {
		return
	}
	log.Default().Printf("Warning for the source %s\n", source)
	for k, v := range warn {
		log.Default().Printf("header %s unused with the value %s\n", k, v)
	}
}

func (l *lotto) sourceReader(option LoadAPIOption, filePath string) (SourceReader, error) {
	var resp *http.Response
	var req *http.Request
	var source SourceReader
	var err error

	url := fmt.Sprintf("%s/%s", BasePath, filePath)
	if req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	if resp, err = l.httpClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if source, err = newSourceReader(resp.Body, option, filePath); err != nil {
		return nil, err
	}

	return source, nil
}

/*
** Private method use to configure the csv parser config
** Allow to enrich the draw list in a lotto instance
 */

func (l *lotto) drawRecorder(rec any, context csvparser.ContextRecord) error {
	var draw Draw
	var err error
	var drawType DrawType
	var obj DrawRecorder
	var ok bool
	var contextValue string

	if obj, ok = rec.(DrawRecorder); !ok {
		return ErrDrawTypeDecode
	}
	if contextValue, ok = context.Get(DrawTypeKey); !ok {
		return ErrDrawTypeKeyNotFound
	}
	if drawType, err = DrawTypeConverter(contextValue); err != nil {
		return err
	}
	if draw, err = obj.ConvertDraw(drawType); err != nil {
		return err
	}
	l.draws = append(l.draws, draw)

	return nil
}

func (l *lotto) csvParseConfiguration(drawType DrawType) csvparser.ParseConfig {
	conf := csvparser.ParseConfig{
		Context:        csvparser.NewContextRecord(),
		RecordObject:   l.drawRecorder,
		CommaSeparator: ';',
	}
	conf.Context.Set(DrawTypeKey, string(drawType))

	return conf
}
