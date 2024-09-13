package pikpakapi

import (
	"fmt"
	"net/http"
)

// It's handles some simple error codes, returns true when error handling is complete, otherwise returns false.
func (p *PikPak) errorCodeHandler(code int64, req *http.Request) error {
	switch code {
	case 9:
		err := p.AuthCaptchaToken(fmt.Sprintf("%s:%s", req.Method, req.URL.Path))
		return err
	case 16:
		err := p.Login()
		return err
	case 4126:
		err := p.RefreshToken()
		return err
	default:
		return fmt.Errorf("Unknown how to handle error, error_code: %d", code)
	}
}
