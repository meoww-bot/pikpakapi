package pikpakapi_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/52funny/pikpakapi"
	"github.com/stretchr/testify/assert"
)

func RandStr(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := ""
	for i := 0; i < n; i++ {
		s += fmt.Sprintf("%c", r.Intn(26)+97)
	}
	return s
}

func TestPikPakFolderCreate(t *testing.T) {
	p := BuildPikPak()
	err := p.Login()
	assert.Nil(t, err)
	str := RandStr(10)
	randDir := pikpakapi.NewPath(str)

	id, err := p.CreateDir(randDir)
	assert.Nil(t, err)
	id2, err := p.GetDirID(randDir)
	assert.Nil(t, err)
	assert.Equal(t, id, id2)
	err = p.DeleteBatchFiles(id)
	assert.Nil(t, err)
}

func TestPikPakDeepFolderCreate(t *testing.T) {
	p := BuildPikPak()
	err := p.Login()
	assert.Nil(t, err)

	strs := []string{RandStr(10), RandStr(10), RandStr(10)}
	dir := pikpakapi.NewPath(strings.Join(strs, "/"))

	id, err := p.CreateDir(dir)
	assert.Nil(t, err)
	id2, err := p.GetDirID(dir)
	assert.Nil(t, err)
	assert.Equal(t, id, id2)

	parentID, err := p.GetDirID(pikpakapi.NewPath(strs[0]))
	assert.Nil(t, err)

	err = p.DeleteBatchFiles(parentID)
	assert.Nil(t, err)
}
