package pikpakapi

import (
	"bytes"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

func (p *PikPak) RefreshToken() error {
	url := "https://user.mypikpak.com/v1/auth/token"
	m := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"grant_type":    "refresh_token",
		"refresh_token": p.refreshToken,
	}
	bs, err := jsoniter.Marshal(&m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bs))
	if err != nil {
		return err
	}
	bs, err = p.send(req)
	if err != nil {
		return err
	}
	p.JwtToken = gjson.GetBytes(bs, "access_token").String()
	p.refreshToken = gjson.GetBytes(bs, "refresh_token").String()
	p.RefreshSecond = gjson.GetBytes(bs, "expires_in").Int()
	logger.Debug("RefreshToken", "access_token", p.JwtToken, "refresh_token", p.refreshToken)
	return nil
}
