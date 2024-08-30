package pikpakapi

import (
	"bytes"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func (p *PikPak) CreateUrlFile(parentID, url string) error {
	m := map[string]interface{}{
		"kind":        "drive#file",
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
		return err
	}
	req, err := http.NewRequest("POST", "https://api-drive.mypikpak.com/drive/v1/files", bytes.NewBuffer(bs))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Product_flavor_name", "cha")
	req.Header.Set("X-Captcha-Token", p.CaptchaToken)
	req.Header.Set("X-Client-Version-Code", "10083")
	req.Header.Set("X-Peer-Id", p.DeviceId)
	req.Header.Set("X-User-Region", "1")
	req.Header.Set("X-Alt-Capability", "3")
	req.Header.Set("Country", "CN")
	bs, err = p.sendWithErrHandle(req, bs)
	if err != nil {
		return err
	}
	task := jsoniter.Get(bs, "task")
	logger.Debug("Create url file", "task", task.ToString())
	// phase := task.Get("phase").ToString()
	// if phase == "PHASE_TYPE_COMPLETE" {
	// 	return nil
	// } else {
	// 	return fmt.Errorf("create file error: %s", phase)
	// }
	return nil
}
