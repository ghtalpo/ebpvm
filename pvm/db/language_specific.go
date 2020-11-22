package db

import (
	"fmt"

	"github.com/ghtalpo/ebpvm/pvm/fileutil"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

var (
	lang               string
	valDepedentForLang int
)

// SetGameLanguage is
func SetGameLanguage(inLang string) {
	if lang == "" {
		lang = inLang
		setupVariables()
	} else {
		panic("Do not call SetGameLanguage twice.")
	}
}

// GetGameLanguage is
func GetGameLanguage() string {
	return getGameLanguage()
}

func getGameLanguage() string {
	return lang
}

// CheckDatFiles is
func CheckDatFiles() {
	if fileutil.CheckErrorsInDatFiles(getGameLanguage()) {
		panic(false)
	}
}

// LoadDataFiles is
func LoadDataFiles() {
	fileutil.LoadTestData()
}

func setupVariables() {
	switch getGameLanguage() {
	case "en":
		valDepedentForLang = 0x1
	case "ko":
		valDepedentForLang = 0x2
	default:
		panic(fmt.Sprintf("unsupported language:%s", getGameLanguage()))
	}
}

func getValDepedentForLang() int {
	if lang == "" {
		panic("call SetGameLanguage first")
	}
	return valDepedentForLang
}

func decodeToUtf8(src string) string {
	switch getGameLanguage() {
	case "en":
		return src
	case "ko":
		// cp949 to utf8
		got, _, err := transform.String(korean.EUCKR.NewDecoder(), src)
		if err != nil {
			panic(err)
		}
		return got
	default:
		panic(fmt.Sprintf("unsupported language:%s", getGameLanguage()))
	}
}

func encodeFromUtf8(src string) string {
	switch getGameLanguage() {
	case "en":
		return src
	case "ko":
		// utf8 to cp949
		got, _, err := transform.String(korean.EUCKR.NewEncoder(), src)
		if err != nil {
			panic(err)
		}
		return got
	default:
		panic(fmt.Sprintf("unsupported language:%s", getGameLanguage()))
	}
}
