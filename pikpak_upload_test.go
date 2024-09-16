package pikpakapi_test

import (
	"crypto/rand"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPikPakFakeFileUpload(t *testing.T) {
	p := BuildPikPak()
	err := p.Login()
	assert.Nil(t, err)
	// f, err := os.CreateTemp(os.TempDir(), "*.txt")
	f, err := os.CreateTemp(".", "*.txt")
	assert.Nil(t, err)

	buf := make([]byte, 1024*1024)
	io.ReadFull(rand.Reader, buf)
	f.Write(buf)
	f.Close()

	id, err := p.UploadFile("", f.Name())
	assert.Nil(t, err)
	err = p.DeleteBatchFiles(id)
	assert.Nil(t, err)
	os.Remove(f.Name())
}
