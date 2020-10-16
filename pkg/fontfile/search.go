package fontfile

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	if file.IsDir() {
		return FontFile{}, false
	}

	re := regexp.MustCompile(`(?i)(\w+)-(\w+)\.(otf|ttf|woff|woff2)`)
	match := re.FindStringSubmatch(file.Name())

	if match == nil {
		return FontFile{}, false
	}

	weightRe := regexp.MustCompile(`(?i)(thin|extralight|light|regular|medium|semibold|bold|extrabold|black)?(italic)?`)
	weightMatch := weightRe.FindStringSubmatch(match[2])

	wString := weightMatch[1]

	var weight int

	if strings.EqualFold(wString, "thin") {
		weight = 100
	} else if strings.EqualFold(wString, "extralight") {
		weight = 200
	} else if strings.EqualFold(wString, "light") {
		weight = 300
	} else if strings.EqualFold(wString, "regular") {
		weight = 400
	} else if strings.EqualFold(wString, "medium") {
		weight = 500
	} else if strings.EqualFold(wString, "semibold") {
		weight = 600
	} else if strings.EqualFold(wString, "bold") {
		weight = 700
	} else if strings.EqualFold(wString, "extrabold") {
		weight = 800
	} else if strings.EqualFold(wString, "black") {
		weight = 900
	} else if wString == "" {
		weight = 400
	}

	if weightMatch == nil {
		return FontFile{}, false
	}

	return FontFile{
		Path:   filepath.Join(basepath, file.Name()),
		Name:   match[1],
		Weight: weight,
	}, true
}
