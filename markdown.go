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

	md.AppendParagraph("# " + title)

	return md
}

// isTitle - is this line a title?
func isTitle(strline string) bool {
	ns := strings.TrimLeft(strline, " ")
	if len(ns) > 0 {
		return ns[0] == '#'
	}

	return false
}

// GetMarkdownString - get markdown string
func (md *Markdown) GetMarkdownString(lst *KeywordMappingList) string {
	if lst != nil && len(lst.Keywords) > 0 {
		lstline := strings.Split(md.str, "\n")
		for _, v := range lst.Keywords {
			if v.URL == "" {
				for i, cl := range lstline {
					if !isTitle(cl) {
						lstline[i] = strings.Replace(cl, v.Keyword,
							adacorebase.AppendString("``", v.Keyword, "``"), -1)
					}
				}
			} else {
				for i, cl := range lstline {
					if !isTitle(cl) {
						lstline[i] = strings.Replace(cl, v.Keyword,
							adacorebase.AppendString("[", v.Keyword, "]("+v.URL+")"), -1)
					}
				}
			}
		}

		md.str = strings.Join(lstline, "\n")
	}

	return md.str
}

// AppendParagraph - append paragraph
func (md *Markdown) AppendParagraph(str string) string {
	md.str = adacorebase.AppendString(md.str, str+"\n\n")

	return md.str
}

// AppendTable - append a table
func (md *Markdown) AppendTable(head []string, data [][]string) string {
	if len(head) > 0 {
		str := "|"

		for _, hv := range head {
			str += hv + "|"
		}

		str += "\n|"

		for range head {
			str += "---|"
		}

		str += "\n"

		for _, li := range data {
			str += "|"
			for _, ld := range li {
				str += ld + "|"
			}
			str += "\n"
		}

		md.str = adacorebase.AppendString(md.str, str+"\n\n")
	}

	return md.str
}
