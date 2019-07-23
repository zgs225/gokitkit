package gokitkit

import (
	"net/url"
)

// CopyURL 复制 url 并替换其 path
func CopyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
