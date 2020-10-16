package gen

import (
	"sort"
	"strings"

	"github.com/cjdenio/ffgen/pkg/fontfile"
)

func GenerateRules(rules []fontfile.FontFile, cssFilePath string) string {
	var generatedRules []string

	sort.Slice(rules, func(a, b int) bool {
		return rules[a].Weight < rules[b].Weight
	})

	for _, v := range rules {
		generatedRules = append(generatedRules, v.CssRule(cssFilePath))
	}

	return strings.Join(generatedRules, "\n\n")
}
