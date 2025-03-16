package pikpakapi

import (
	"fmt"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
)

func (p *PikPak) GetTask(taskID string) (TaskResponse, error) {
	query := url.Values{}
	query.Add("with", "reference_resource")
	query.Add("type", "offline")
	query.Add("thumbnail_size", "SIZE_SMALL")
	query.Add("limit", "100")
	filters := fmt.Sprintf(`{"id":{"in":"%s"}}`, taskID)

	query.Add("filters", filters)

	req, err := http.NewRequest("GET", "https://api-drive.mypikpak.com/drive/v1/tasks?"+query.Encode(), nil)
	if err != nil {
		return TaskResponse{}, err
	}
	req.Header.Set("X-Captcha-Token", p.CaptchaToken)
	req.Header.Set("Content-Type", "application/json")
	bs, err := p.sendWithErrHandle(req, nil)
	if err != nil {
		return TaskResponse{}, err
	}
	var result TaskResponse
	err = jsoniter.Unmarshal(bs, &result)
	if err != nil {
		return TaskResponse{}, err
	}

	return result, nil
}
