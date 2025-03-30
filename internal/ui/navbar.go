package ui

import (
	"net/url"

	"github.com/a-h/templ"
)

type NavLink struct {
	Name     string
	URL      url.URL
	Active   bool
	Disabled bool
}

func (n NavLink) Attributes() templ.Attributes {
	attrs := make(templ.Attributes, 3)
	if n.Active {
		attrs["class"] = "nav-item nav-link active"
		attrs["href"] = n.URL.String()
	} else if n.Disabled {
		attrs["class"] = "nav-item nav-link disabled"
		attrs["href"] = "#"
		attrs["aria-disabled"] = "true"
		attrs["tabindex"] = "-1"
	} else {
		attrs["class"] = "nav-item nav-link"
		attrs["href"] = n.URL.String()
	}
	return attrs
}
