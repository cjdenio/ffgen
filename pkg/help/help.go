package help

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintHelp() {
	white := color.New(color.FgWhite).Add(color.Bold)
	muted := color.New(color.FgHiBlack)

	fmt.Printf(`%s - quickly and easily generate @font-face rules for CSS!

%s
  ffgen <directory> <css file>
  %s ffgen ./fonts/Raleway ./css/fonts.css

%s
  https://github.com/cjdenio/ffgen
`, white.Sprint("ffgen 🎨"), white.Sprint("USAGE 🛠️"), muted.Sprint("Example:"), white.Sprint("GITHUB 🐙"))
}
