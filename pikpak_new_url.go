package pikpakapi

import (
	"bytes"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

func (p *PikPak) CreateUrlFile(parentID, url string) (string, error) {
	m := map[string]interface{}{
		"kind":        KIND_FILE,
		"upload_type": "UPLOAD_TYPE_URL",
		"url": map[string]string{
			"url": url,
		},
	}
	if parentID != "" {
		m["parent_id"] = parentID
	}
	bs, err := jsoniter.Marshal(&m)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", "https://api-drive.mypikpak.com/drive/v1/files", bytes.NewBuffer(bs))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Captcha-Token", p.CaptchaToken)
	bs, err = p.sendWithErrHandle(req, bs)
	if err != nil {
		return "", err
	}
	name := gjson.GetBytes(bs, "task.name").String()
	logger.Debug("Create url file", "task", name)
	return name, nil
}
