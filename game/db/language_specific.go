package db

import (
	"fmt"

	"github.com/ghtalpo/egb/common/fileutil"
	// "golang.org/x/text/encoding/korean"
	// "golang.org/x/text/transform"
)

var (
	lang               string
	valDepedentForLang int
)

// SetGameLanguage ...
func SetGameLanguage(inLang string) {
	if lang == "" {
		lang = inLang
		setupVariables()
	} else {
		panic("Do not call SetGameLanguage twice.")
	}
}

// GetGameLanguage ...
func GetGameLanguage() string {
	return getGameLanguage()
}

func getGameLanguage() string {
	return lang
}

// CheckErrorsInDatFiles checks existence of all dat files
func CheckErrorsInDatFiles(lang string) bool {
	return fileutil.IsInvalid(getPathForTestData())
}

// CheckDatFiles ...
func CheckDatFiles() {
	if CheckErrorsInDatFiles(getGameLanguage()) {
		panic(false)
	}
}

var (
	testData []byte
)

// GetTestData get
func GetTestData() []byte {
	return testData
}

// GetTestDataSub get
func GetTestDataSub(start int, end int) []byte {
	return testData[start:end]
}

func getPathForTestData() string {
	return "_resources/dat/testdata.dat"
}

// LoadTestData ...
func LoadTestData() {
	testData = fileutil.ReadByteArray(getPathForTestData())
}

// LoadDataFiles ...
func LoadDataFiles() {
	LoadTestData()
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

// func decodeToUtf8(src string) string {
// 	switch getGameLanguage() {
// 	case "en":
// 		return src
// 	case "ko":
// 		// cp949 to utf8
// 		got, _, err := transform.String(korean.EUCKR.NewDecoder(), src)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return got
// 	default:
// 		panic(fmt.Sprintf("unsupported language:%s", getGameLanguage()))
// 	}
// }

// func encodeFromUtf8(src string) string {
// 	switch getGameLanguage() {
// 	case "en":
// 		return src
// 	case "ko":
// 		// utf8 to cp949
// 		got, _, err := transform.String(korean.EUCKR.NewEncoder(), src)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return got
// 	default:
// 		panic(fmt.Sprintf("unsupported language:%s", getGameLanguage()))
// 	}
// }
