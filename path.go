package pikpakapi

import (
	"strings"
)

type Path struct {
	inner string
}

func NewPath(path string) Path {
	path = strings.TrimLeft(path, "/")
	return Path{inner: path}
}

func (p Path) String() string {
	return p.inner
}

func (p *Path) Cut() []string {
	path := p.String()
	if path == "" {
		return []string{}
	}
	arr := strings.Split(path, "/")
	return arr
}

func (p *Path) Parent() Path {
	path := p.String()
	if path == "" {
		return NewPath("")
	}
	arr := strings.Split(path, "/")
	if len(arr) == 1 {
		return NewPath("")
	}
	return Path{inner: strings.Join(arr[:len(arr)-1], "/")}
}

func (p *Path) Name() string {
	path := p.String()
	if path == "" {
		return ""
	}
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
