package fileutil

import (
	"encoding/json"
	"fmt"
	"path"
)

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

func getFilenameForLanguage(filename string, language string) string {
	if len(language) != 0 {
		ext := path.Ext(filename)
		base := filename[0 : len(filename)-len(ext)]
		return fmt.Sprintf("%s_%s%s", base, language, ext)
	}
	return filename
}

// LoadJSON load
func LoadJSON(path string, lang string) map[string]interface{} {
	data := ReadByteArray(getFilenameForLanguage(path, lang))

	j := make(map[string]interface{})
	if err := json.Unmarshal(data, &j); err != nil {
		panic(err)
	}
	return j
}

// LoadTestData load
func LoadTestData() {
	testData = ReadByteArray(getPathForTestData())
}
