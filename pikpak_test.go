package pikpakapi_test

import (
	"os"
	"testing"

	"github.com/52funny/pikpakapi"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func ReadEnv() (string, string) {
	// dotenv := godotenv.Load()
	godotenv.Load()
	u := os.Getenv("USERNAME")
	p := os.Getenv("PASSWORD")
	return u, p
}

func BuildPikPak() pikpakapi.PikPak {
	username, password := ReadEnv()
	return pikpakapi.NewPikPak(username, password)
}

func TestPikPakLogin(t *testing.T) {
	p := BuildPikPak()
	err := p.Login()
	assert.Nil(t, err)
}

func TestPikPakFiles(t *testing.T) {
	p := BuildPikPak()
	p.Login()
	id, err := p.GetDirID(pikpakapi.NewPath("/"))
	assert.Nil(t, err)
	assert.Equal(t, "", id)
	lists, err := p.GetFolderFileStatList(id)
	assert.Nil(t, err)
	assert.NotEmpty(t, lists)
}
