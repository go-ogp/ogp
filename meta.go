package ogp

import (
	"fmt"
	"html/template"
	"strings"
)

type metaBuilder struct {
	tags []string
}

func (b *metaBuilder) Add(ns, prop string, content interface{}) *metaBuilder {
	var tag string
	if prop != "" {
		tag = fmt.Sprintf(`<meta property="%s:%s" content="%v">`, ns, prop, content)
	} else {
		tag = fmt.Sprintf(`<meta property="%s" content="%v">`, ns, content)
	}
	b.tags = append(b.tags, tag)
	return b
}

func (b *metaBuilder) Include(mb *metaBuilder) *metaBuilder {
	b.tags = append(b.tags, mb.tags...)
	return b
}

func (b *metaBuilder) HTML() template.HTML {
	return template.HTML(b.String())
}

func (b *metaBuilder) String() string {
	return strings.Join(b.tags, "\n")
}
