package lotto

import "github.com/winning-number/fdj-sdk-lotto/draw"

// Endpoint to contact the FDJ history service
const (
	APIBasePath = "https://media.fdj.fr/static/csv/loto"
)

// Source is the type of lotto history
type Source int16

// File names for the different lotto history from the FDJ API
const (
	GrandLotoZipName       = "grandloto_201912.zip"
	GrandLotoNoelZipName   = "lotonoel_201703.zip"
	SuperLoto199605ZipName = "superloto_199605.zip"
	SuperLoto200810ZipName = "superloto_200810.zip"
	SuperLoto201703ZipName = "superloto_201703.zip"
	SuperLoto201907ZipName = "superloto_201907.zip"
	Loto197605ZipName      = "loto_197605.zip"
	Loto200810ZipName      = "loto_200810.zip"
	Loto201703ZipName      = "loto_201703.zip"
	Loto201902ZipName      = "loto_201902.zip"
	Loto201911ZipName      = "loto_201911.zip"
)

// Source constants for the different lotto history from the FDJ API
const (
	GrandLoto Source = iota
	GrandLotoNoel
	SuperLoto199605
	SuperLoto200810
	SuperLoto201703
	SuperLoto201907
	Loto197605
	Loto200810
	Loto201703
	Loto201902
	// most recent lotto history
	Loto201911
)

// SourceInfo contains the information about the source
// APIZipName is the name of the zip file containing the history from the FDJ api
// Type is the type of draw
// Version is the version of the draw for conversion
// Name is the name of the source
type SourceInfo struct {
	APIZipName string
	Type       draw.Type
	Version    draw.Version
	Name       Source
}

// URL returns the url to download the source
func (s SourceInfo) URL() string {
	return APIBasePath + "/" + s.APIZipName
}

// GetSourceInfo returns the source info for the given source
// If the source is not found, it returns the most recent source info
// the most recent source info is Loto201911 (from november 2019)
func GetSourceInfo(source Source) SourceInfo {
	switch source {
	case GrandLoto:
		return SourceInfo{
			APIZipName: GrandLotoZipName,
			Type:       draw.GrandLottoType,
			Version:    draw.V3,
			Name:       GrandLoto,
		}
	case GrandLotoNoel:
		return SourceInfo{
			APIZipName: GrandLotoNoelZipName,
			Type:       draw.XmasLottoType,
			Version:    draw.V3,
			Name:       GrandLotoNoel,
		}
	case SuperLoto199605:
		return SourceInfo{
			APIZipName: SuperLoto199605ZipName,
			Type:       draw.SuperLottoType,
			Version:    draw.V0,
			Name:       SuperLoto199605,
		}
	case SuperLoto200810:
		return SourceInfo{
			APIZipName: SuperLoto200810ZipName,
			Type:       draw.SuperLottoType,
			Version:    draw.V2,
			Name:       SuperLoto200810,
		}
	case SuperLoto201703:
		return SourceInfo{
			APIZipName: SuperLoto201703ZipName,
			Type:       draw.SuperLottoType,
			Version:    draw.V3,
			Name:       SuperLoto201703,
		}
	case SuperLoto201907:
		return SourceInfo{
			APIZipName: SuperLoto201907ZipName,
			Type:       draw.SuperLottoType,
			Version:    draw.V3,
			Name:       SuperLoto201907,
		}
	case Loto197605:
		return SourceInfo{
			APIZipName: Loto197605ZipName,
			Type:       draw.LottoType,
			Version:    draw.V1,
			Name:       Loto197605,
		}
	case Loto200810:
		return SourceInfo{
			APIZipName: Loto200810ZipName,
			Type:       draw.LottoType,
			Version:    draw.V2,
			Name:       Loto200810,
		}
	case Loto201703:
		return SourceInfo{
			APIZipName: Loto201703ZipName,
			Type:       draw.LottoType,
			Version:    draw.V3,
			Name:       Loto201703,
		}
	case Loto201902:
		return SourceInfo{
			APIZipName: Loto201902ZipName,
			Type:       draw.LottoType,
			Version:    draw.V3,
			Name:       Loto201902,
		}
	case Loto201911:
		return SourceInfo{
			APIZipName: Loto201911ZipName,
			Type:       draw.LottoType,
			Version:    draw.V4,
			Name:       Loto201911,
		}
	default:
		return SourceInfo{
			APIZipName: Loto201911ZipName,
			Type:       draw.LottoType,
			Version:    draw.V4,
			Name:       Loto201911,
		}
	}
}

// SourceInfoAll returns all the source info
func SourceInfoAll() []SourceInfo {
	sources := SourceAll()
	infos := make([]SourceInfo, len(sources))

	for i, source := range sources {
		infos[i] = GetSourceInfo(source)
	}

	return infos
}

// SourceAll returns all the sources available
func SourceAll() []Source {
	return []Source{
		GrandLoto,
		GrandLotoNoel,
		SuperLoto199605,
		SuperLoto200810,
		SuperLoto201703,
		SuperLoto201907,
		Loto197605,
		Loto200810,
		Loto201703,
		Loto201902,
		Loto201911,
	}
}
