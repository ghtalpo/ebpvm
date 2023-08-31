package fileutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// ReadByteArray reads all contents from file
// this may does not work properly in every platform
// should embed images into go file
// https://ebiten.org/tour/image.html
func ReadByteArray(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err?", err)
		log.Fatal(err)
	}
	data, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		fmt.Println("read file err?", err2)
		log.Fatal(err2)
	}
	return data
}

// IsExists ...
func IsExists(path string) bool {
	return IsValid(path)
}

// IsValid ...
func IsValid(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// log.Fatal(err)
		return false
	}
	return true
}

// IsInvalid ...
func IsInvalid(path string) bool {
	return !IsValid(path)
}

// Create ...
func Create(filename string) (*OFS, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	ofs := OFS{}
	ofs.f = f

	return &ofs, nil
}

// Write ...
func Write(ofs *OFS, s string, l int) {
	if ofs.f == nil {
		panic(false)
	}
	b := []byte(s)
	if len(b) < l {
		for i := 0; i < l-len(b); i++ {
			b = append(b, 0)
		}
	} else if len(b) > l {
		b = b[0:l]
	}

	ofs.f.Write(b)
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
