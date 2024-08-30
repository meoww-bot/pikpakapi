package pikpakapi

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// Get cloud quota
func (p *PikPak) GetQuota() (Quota, error) {
	req, err := http.NewRequest("GET", "https://api-drive.mypikpak.com/drive/v1/about", nil)
	if err != nil {
		return Quota{}, err
	}
	bs, err := p.sendWithErrHandle(req)
	if err != nil {
		return Quota{}, err
	}
	var quotaRes quotaResponse
	err = jsoniter.Unmarshal(bs, &quotaRes)
	if err != nil {
		return Quota{}, err
	}
	return quotaRes.Quota, nil
}
