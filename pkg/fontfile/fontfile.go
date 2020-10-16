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
}`, fontFile.Name, relativePath)
}
