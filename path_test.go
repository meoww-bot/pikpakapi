package pikpakapi_test

import (
	"testing"

	"github.com/52funny/pikpakapi"
	"github.com/stretchr/testify/assert"
)

func TestPathMustStartWithSlash(t *testing.T) {
	path := "/abc"
	p := pikpakapi.NewPath(path)
	assert.Equal(t, p.String(), path)
}
