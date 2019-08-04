package adacore

import (
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
func (md *Markdown) GetMarkdownString() string {
	return md.str
}

// AppendString - append string
func (md *Markdown) AppendString(str string) string {
	md.str = adacorebase.AppendString(md.str, str+"\n\n")

	return md.str
}
