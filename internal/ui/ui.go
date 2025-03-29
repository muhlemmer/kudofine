package ui

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/muhlemmer/kudofine/internal/config"
	"github.com/muhlemmer/kudofine/internal/model"
)

type Views struct {
	config config.Config
}

func NewViews(config config.Config) Views {
	return Views{config: config}
}

func (v Views) Hello(ctx context.Context, w http.ResponseWriter, resp model.HelloResponse) error {
	const pageName = "Hello"
	data := v.newPageData(pageName, hello(resp), nil)
	return page(data).Render(ctx, w)
}

type pageData struct {
	config   config.Config
	pageName string
	main     templ.Component
	footer   templ.Component
}

func (v Views) newPageData(pageName string, body, footer templ.Component) pageData {
	if footer == nil {
		footer = templ.Raw("")
	}
	return pageData{
		config:   v.config,
		pageName: pageName,
		main:     body,
		footer:   footer,
	}
}
