package handler

import (
	"strings"
)

func endpoint(method, host, path string) string {
	var b strings.Builder
	if method != "" {
		b.WriteString(method)
		b.WriteString(" ")
	}
	if host != "" {
		b.WriteString(host)
	}
	b.WriteString("/")
	b.WriteString(path)
	return b.String()
}
