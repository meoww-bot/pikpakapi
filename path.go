package pikpakapi

import (
	"strings"
)

type Path string

func NewPath(path string) Path {
	path = strings.TrimLeft(path, "/")
	return Path(path)
}

func (p *Path) String() string {
	return string(*p)
}

func (p *Path) Cut() []string {
	path := p.String()
	if path == "" {
		return []string{}
	}
	arr := strings.Split(path, "/")
	return arr
}
