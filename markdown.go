package adacore

import (
	"strings"

	adacorebase "github.com/zhs007/adacore/base"
)

// Markdown - markdown
type Markdown struct {
	// Title - title
	Title string
	// str - markdown string
	str string
}

// NewMakrdown - new Markdown
func NewMakrdown(title string) *Markdown {
	md := &Markdown{
		Title: title,
	}

	md.AppendString("# " + title)

	return md
}

// GetMarkdownString - get markdown string
func (md *Markdown) GetMarkdownString(lst *KeywordMappingList) string {
	if lst != nil && len(lst.Keywords) > 0 {
		for _, v := range lst.Keywords {
			if v.URL == "" {
				md.str = strings.Replace(md.str, v.Keyword,
					adacorebase.AppendString("``", v.Keyword, "``"), -1)
			} else {
				md.str = strings.Replace(md.str, v.Keyword,
					adacorebase.AppendString("``[", v.Keyword, "]("+v.URL+")``"), -1)
			}
		}
	}

	return md.str
}

// AppendString - append string
func (md *Markdown) AppendString(str string) string {
	md.str = adacorebase.AppendString(md.str, str+"\n\n")

	return md.str
}
