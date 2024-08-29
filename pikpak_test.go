package pikpakapi_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/52funny/pikpakapi"
	"github.com/joho/godotenv"
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
	p.Login()
	fmt.Println(p.RefreshSecond)
	fmt.Println(username, password)
}
