package pikpakapi_test

import (
	"os"
	"testing"

	"github.com/52funny/pikpakapi"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func ReadEnv() (string, string, string) {
	// dotenv := godotenv.Load()
	godotenv.Load()
	u := os.Getenv("USERNAME")
	p := os.Getenv("PASSWORD")
	proxy := os.Getenv("PROXY")
	return u, p, proxy
}

func BuildPikPak() pikpakapi.PikPak {
	username, password, proxy := ReadEnv()
	p := pikpakapi.NewPikPak(username, password)
	if proxy != "" {
		p.SetProxy(proxy)
	}
	return p
}

func TestPikPakLogin(t *testing.T) {
	p := BuildPikPak()
	err := p.Login()
	assert.Nil(t, err)
}

func TestPikPakChangeJwtToken(t *testing.T) {
	p := BuildPikPak()
	err := p.Login()
	assert.Nil(t, err)
	p.JwtToken = "123456"
	stats, err := p.GetDirFilesStatByPath(pikpakapi.NewPath("/"))
	assert.Nil(t, err)
	assert.NotEmpty(t, stats)
}

func TestPikPakFiles(t *testing.T) {
	p := BuildPikPak()
	p.Login()
	id, err := p.GetDirID(pikpakapi.NewPath("/"))
	assert.Nil(t, err)
	assert.Equal(t, "", id)
	lists, err := p.GetDirFilesStat(id)
	assert.Nil(t, err)
	assert.NotEmpty(t, lists)
}
