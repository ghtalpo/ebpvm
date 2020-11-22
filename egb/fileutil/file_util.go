package fileutil

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghtalpo/egb/egb/konst"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// ReadByteArray reads all contents from file
// this may does not work properly in every platform
// should embed images into go file
// https://ebiten.org/tour/image.html
func ReadByteArray(path string) []byte {
	f, err := ebitenutil.OpenFile(path)
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

// IsExists is
func IsExists(path string) bool {
	return isValid(path)
}

func isValid(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// log.Fatal(err)
		return false
	}
	return true
}

func isInvalid(path string) bool {
	return !isValid(path)
}

func getPathForTestData() string {
	return "_resources/dat/testdata.dat"
}

// CheckErrorsInDatFiles checks existence of all dat files
func CheckErrorsInDatFiles(lang string) bool {
	if isInvalid(getPathForTestData()) {
		return true
	}
	return false
}

// Create is
func Create(filename string) (*OFS, int) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	ofs := OFS{}
	ofs.f = f

	return &ofs, konst.OK
}

// Write is
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
