package path

import (
	"path"
)

type Path struct{}

func NewPath() *Path {
	return &Path{}
}

func (*Path) PathIsValid(uri string) bool {
	cleanPath := path.Clean(uri)
	return uri == cleanPath
}

/*
pathIsValid := h.path.PathIsValid(r.RequestURI)

If a client requests http://example.com/path?query=value, the RequestURI would be /path?query=value.

Inputing a path alternation such as: http://localhost:8080/pathIsValid?path=../ 

returns false.
*/