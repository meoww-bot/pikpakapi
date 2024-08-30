package pikpakapi_test

import (
	"testing"

	"github.com/52funny/pikpakapi"
	"github.com/stretchr/testify/assert"
)

func TestPathWithDiffentStyle(t *testing.T) {
	path := "/abc"
	p := pikpakapi.NewPath(path)
	assert.Equal(t, "abc", p.String())

	path2 := "abc"
	p2 := pikpakapi.NewPath(path2)
	assert.Equal(t, "abc", p2.String())

	path3 := "//abc/def"
	p3 := pikpakapi.NewPath(path3)
	assert.Equal(t, "abc/def", p3.String())
}

func TestPathCut(t *testing.T) {
	path := "/abc/def/ghi"
	p := pikpakapi.NewPath(path)
	assert.Equal(t, []string{"abc", "def", "ghi"}, p.Cut())
	path2 := "/"
	p2 := pikpakapi.NewPath(path2)
	assert.Equal(t, []string{}, p2.Cut())

	path3 := "//abc/def/ghi"
	p3 := pikpakapi.NewPath(path3)
	assert.Equal(t, []string{"abc", "def", "ghi"}, p3.Cut())

	path4 := "//abc//def/ghi"
	p4 := pikpakapi.NewPath(path4)
	assert.Equal(t, []string{"abc", "", "def", "ghi"}, p4.Cut())
}

func TestPathParent(t *testing.T) {
	path := "/abc/def/ghi"
	p := pikpakapi.NewPath(path)
	assert.Equal(t, "abc/def", string(p.Parent()))
	path2 := "/"
	p2 := pikpakapi.NewPath(path2)
	assert.Equal(t, "", string(p2.Parent()))
	path3 := "/abc/def//"
	p3 := pikpakapi.NewPath(path3)
	assert.Equal(t, "abc/def/", string(p3.Parent()))
}

func TestPathName(t *testing.T) {
	path := "/abc/def/ghi"
	p := pikpakapi.NewPath(path)
	assert.Equal(t, "ghi", p.Name())
	path2 := "/"
	p2 := pikpakapi.NewPath(path2)
	assert.Equal(t, "", p2.Name())
	path3 := "/abc/def//"
	p3 := pikpakapi.NewPath(path3)
	assert.Equal(t, "", p3.Name())
}
