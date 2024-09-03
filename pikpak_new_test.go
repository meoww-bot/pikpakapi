package pikpakapi_test

import (
	"testing"

	"github.com/52funny/pikpakapi"
	"github.com/stretchr/testify/assert"
)

func TestPikPakNewUrl(t *testing.T) {
	p := BuildPikPak()
	err := p.Login()
	assert.Nil(t, err)
	id, err := p.CreateDir(pikpakapi.NewPath("/tmp"))
	assert.Nil(t, err)
	name, err := p.CreateUrlFile(id, "magnet:?xt=urn:btih:e9c98e3ed488611abc169a81d8a21487fd1d0732")
	assert.Nil(t, err)
	assert.Equal(t, "【www.gaoqing.tv】肖申克的救赎 [国 英] The.Shawshank.Redemption.1994.BluRay.1080p.DTS.2Audio.x264-CHD", name)
	err = p.DeleteBatchFiles(id)
	assert.Nil(t, err)
}
