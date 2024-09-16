package pikpakapi

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

const userAgent = `ANDROID-com.pikcloud.pikpak/1.21.0`
const clientID = `YNxT9w7GMdWvEOKa`
const clientSecret = `dbw2OtmVEeuUvIptb1Coyg`

type PikPak struct {
	Account       string `json:"account"`
	Password      string `json:"password"`
	JwtToken      string `json:"token"`
	refreshToken  string
	CaptchaToken  string `json:"captchaToken"`
	Sub           string `json:"userId"`
	DeviceId      string `json:"deviceId"`
	RefreshSecond int64  `json:"refreshSecond"`
	client        *http.Client
	cache         *cache
}

func NewPikPak(account, password string) PikPak {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	seed := make([]byte, 16)
	io.ReadFull(rand.Reader, seed)
	n := md5.Sum(append([]byte(account+clientID), seed...))
	return PikPak{
		Account:  account,
		Password: password,
		DeviceId: hex.EncodeToString(n[:]),
		client:   client,
		cache:    newCache(),
	}
}

// Default proxy from environment
func (p *PikPak) SetDefaultProxy() {
	p.client.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}
}

// Seting the cumstom proxy
func (p *PikPak) SetProxy(proxy string) error {
	u, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	p.client.Transport = &http.Transport{
		Proxy: http.ProxyURL(u),
	}
	return nil
}

// Login method
func (p *PikPak) Login() error {
	err := p.authSigninCaptchaToken()
	if err != nil {
		return err
	}
	m := make(map[string]string)
	m["client_id"] = clientID
	m["username"] = p.Account
	m["password"] = p.Password
	bs, err := jsoniter.Marshal(&m)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://user.mypikpak.com/v1/auth/signin", bytes.NewBuffer(bs))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Captcha-Token", p.CaptchaToken)

	bs, err = p.send(req)
	if err != nil {
		return err
	}
	errorCode := gjson.GetBytes(bs, "error_code").Int()
	if errorCode != 0 {
		errorMessage := gjson.GetBytes(bs, "error").String()
		return fmt.Errorf("url: %s error_code: %d, error: %s", req.URL.String(), errorCode, errorMessage)
	}

	p.JwtToken = gjson.GetBytes(bs, "access_token").String()
	p.refreshToken = gjson.GetBytes(bs, "refresh_token").String()
	p.Sub = gjson.GetBytes(bs, "sub").String()
	p.RefreshSecond = gjson.GetBytes(bs, "expires_in").Int()

	logger.Debug("Login params", "access_token", p.JwtToken, "refresh_token", p.refreshToken, "sub", p.Sub, "expires_in", p.RefreshSecond)

	return nil
}

// Get the captcha token
func (p *PikPak) authSigninCaptchaToken() error {
	m := make(map[string]any)
	m["client_id"] = clientID
	m["device_id"] = p.DeviceId
	m["captcha_token"] = p.CaptchaToken
	m["action"] = "POST:/v1/auth/signin"
	m["meta"] = map[string]string{
		"username": p.Account,
	}
	body, err := jsoniter.Marshal(&m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://user.mypikpak.com/v1/shield/captcha/init", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	bs, err := p.send(req)
	if err != nil {
		return err
	}
	errorCode := gjson.GetBytes(bs, "error_code").Int()
	if errorCode != 0 {
		errorMessage := gjson.GetBytes(bs, "error").String()
		return fmt.Errorf("url: %s error_code: %d, error: %s", req.URL.String(), errorCode, errorMessage)
	}

	captchaToken := gjson.GetBytes(bs, "captcha_token").String()
	logger.Debug("Login captcha token", "captcha_token", captchaToken)
	p.CaptchaToken = captchaToken
	return nil
}

// Send the request
func (p *PikPak) send(req *http.Request) ([]byte, error) {
	// Setting the header
	if p.JwtToken != "" {
		req.Header.Set("Authorization", "Bearer "+p.JwtToken)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Device-Id", p.DeviceId)

	// Send the request
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the response
	bs, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// Send the request with error handling
func (p *PikPak) sendWithErrHandle(req *http.Request, body []byte) ([]byte, error) {
	// Add the header
	p.addHeader(req)
START:
	// Send the request
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the response
	bs, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	errorCode := gjson.GetBytes(bs, "error_code").Int()
	if errorCode != 0 {
		// Handle the error code, if successful handling, then repeat the request
		err := p.errorCodeHandler(errorCode, req)
		// Repeat the request
		if err == nil {
			req, err = http.NewRequest(req.Method, req.URL.String(), bytes.NewBuffer(body))
			if err != nil {
				return nil, err
			}
			p.addHeader(req)
			goto START
		}
		errorMessage := gjson.GetBytes(bs, "error").String()
		return nil, fmt.Errorf("url: %s error_code: %d, error: %s", req.URL.String(), errorCode, errorMessage)
	}
	return bs, nil
}

// Setting the header
func (p *PikPak) addHeader(req *http.Request) {
	if p.JwtToken != "" {
		req.Header.Set("Authorization", "Bearer "+p.JwtToken)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Device-Id", p.DeviceId)
}
