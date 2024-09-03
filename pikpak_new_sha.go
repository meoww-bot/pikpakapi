package pikpakapi

import (
	"bytes"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func (p *PikPak) CreateShaFile(parentID, fileName, size, sha string) error {
	m := map[string]interface{}{
		"body": map[string]string{
			"duration": "",
			"width":    "",
			"height":   "",
		},
		"kind":        KIND_FILE,
		"name":        fileName,
		"size":        size,
		"hash":        sha,
		"upload_type": "UPLOAD_TYPE_RESUMABLE",
		"objProvider": map[string]string{
			"provider": "UPLOAD_TYPE_UNKNOWN",
		},
	}
	if parentID != "" {
		m["parent_id"] = parentID
	}
	bs, err := jsoniter.Marshal(&m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://api-drive.mypikpak.com/drive/v1/files", bytes.NewBuffer(bs))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Captcha-Token", p.CaptchaToken)
	bs, err = p.sendWithErrHandle(req, bs)
	if err != nil {
		return err
	}
	file := jsoniter.Get(bs, "file")
	phase := file.Get("phase").ToString()
	if phase == "PHASE_TYPE_COMPLETE" {
		return nil
	} else {
		return fmt.Errorf("create file error: %s", phase)
	}
}
