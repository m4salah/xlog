package main

import (
	"fmt"
	"html/template"
	"net/url"
	"strings"
)

const STARRED_PAGES = "starred"

func init() {
	WIDGET(META_WIDGET, starMeta)
	WIDGET(SIDEBAR_WIDGET, starredPages)
	GET(`/\+/star/{page:.*}`, starHandler)
	DELETE(`/\+/star/{page:.*}`, unstarHandler)
}

func starredPages(p *Page, r Request) template.HTML {
	pages := NewPage(STARRED_PAGES)
	content := strings.TrimSpace(pages.Content())
	if content == "" {
		return template.HTML("")
	}

	list := strings.Split(content, "\n")
	return template.HTML(partial("extension/starred", Locals{
		"pages": list,
	}))
}

func starMeta(p *Page, r Request) template.HTML {
	starred := isStarred(p)

	return template.HTML(partial("extension/star-meta", Locals{
		"csrf":    CSRF(r),
		"starred": starred,
		"action":  fmt.Sprintf("/+/star/%s", url.PathEscape(p.Name)),
	}))
}

func starHandler(w Response, r Request) Output {
	vars := VARS(r)
	page := NewPage(vars["page"])
	if !page.Exists() {
		return Redirect("/")
	}

	starred_pages := NewPage(STARRED_PAGES)
	starred_pages.Write(strings.TrimSpace(starred_pages.Content()) + "\n" + page.Name)
	return Redirect("/" + page.Name)
}

func unstarHandler(w Response, r Request) Output {
	vars := VARS(r)
	page := NewPage(vars["page"])
	if !page.Exists() {
		return Redirect("/")
	}

	starred_pages := NewPage(STARRED_PAGES)
	content := strings.Split(strings.TrimSpace(starred_pages.Content()), "\n")
	new_content := ""
	for _, v := range content {
		if v != page.Name {
			new_content += "\n" + v
		}
	}
	starred_pages.Write(new_content)

	return Redirect("/" + page.Name)
}

func isStarred(p *Page) bool {
	starred_page := NewPage(STARRED_PAGES)
	for _, k := range strings.Split(starred_page.Content(), "\n") {
		if strings.TrimSpace(k) == p.Name {
			return true
		}
	}

	return false
}
