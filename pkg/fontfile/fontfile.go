package fontfile

import (
	"fmt"
	"path/filepath"
)

type FontFile struct {
	Name   string
	Weight int
	Italic bool
	Path   string
}

func (fontFile FontFile) CssRule(cssFilePath string) string {
	relativePath, _ := filepath.Rel(cssFilePath, fontFile.Path)

	return fmt.Sprintf(`@font-face {
  font-family: "%s";
  src: url("%s");
  font-weight: %d;
  font-style: %s;
}`, fontFile.Name, relativePath, fontFile.Weight, map[bool]string{true: "italic", false: "normal"}[fontFile.Italic])
}
