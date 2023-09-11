package embed

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/m4salah/xlog"
	"github.com/m4salah/xlog/extensions/shortcode"
)

func init() {
	shortcode.ShortCode("embed", embedShortcode)
}

func embedShortcode(in xlog.Markdown) template.HTML {
	p := xlog.NewPage(strings.TrimSpace(string(in)))
	if p == nil || !p.Exists() {
		return template.HTML(fmt.Sprintf("Page: %s doesn't exist", in))
	}

	return p.Render()
}
