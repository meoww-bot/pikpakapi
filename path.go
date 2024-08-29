package pikpakapi

import "strings"

type Path string

func NewPath(path string) Path {
	if path == "" || path[0] != '/' {
		panic("Path must start with /")
	}
	return Path(path)
}

func (p *Path) String() string {
	return string(*p)
}

func (p *Path) Cut() []string {
	arr := strings.Split(p.String()[1:], "/")
	return arr
}
