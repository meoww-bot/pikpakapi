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

func TestPikPak(t *testing.T) {
	username, password := ReadEnv()
	p := pikpakapi.NewPikPak(username, password)
	err := p.Login()
	assert.Nil(t, err)
}
