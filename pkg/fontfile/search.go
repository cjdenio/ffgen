package fontfile

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func SearchDirectory(directory string) ([]FontFile, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var fontFiles []FontFile

	for i := 0; i < len(files); i++ {
		parsed, valid := parseFile(directory, files[i])
		if valid {
			fontFiles = append(fontFiles, parsed)
		}
	}

	return fontFiles, nil
}

func parseFile(basepath string, file os.FileInfo) (FontFile, bool) {
	return FontFile{
		Path: filepath.Join(basepath, file.Name()),
	}, !file.IsDir()
}
