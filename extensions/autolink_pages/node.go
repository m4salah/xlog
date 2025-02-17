package autolink_pages

import (
	. "github.com/m4salah/dlog"
	"github.com/yuin/goldmark/ast"
)

var KindPageLink = ast.NewNodeKind("PageLink")

type PageLink struct {
	ast.BaseInline
	page Page
}

func (_ *PageLink) Kind() ast.NodeKind {
	return KindPageLink
}

func (p *PageLink) Dump(source []byte, level int) {
	m := map[string]string{
		"value": p.page.Name(),
	}
	ast.DumpHelper(p, source, level, m, nil)
}
