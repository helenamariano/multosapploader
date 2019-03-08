package loader

import (
	"io/ioutil"
	"os"
)

// Alc is the raw bytes of the Application Load Certificate
type Alc []byte

// ParseAlc readers an ALC file. It is assumed that the
// file is in format as defined by section ALC_DATA in
// https://www.multos.com/uploads/FIF.pdf
func ParseAlc(filename string) (Alc, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	alc, err := ioutil.ReadAll(file)
	return alc, nil
}
