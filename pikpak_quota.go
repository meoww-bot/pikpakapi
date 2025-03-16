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
	bs, err := p.sendWithErrHandle(req, nil)
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

func (p *PikPak) GetVipQuantity() (QuantityResponse, error) {
	req, err := http.NewRequest("GET", "https://api-drive.mypikpak.com/vip/v1/quantity/list?type=transfer&limit=200", nil)
	if err != nil {
		return QuantityResponse{}, err
	}
	bs, err := p.sendWithErrHandle(req, nil)
	if err != nil {
		return QuantityResponse{}, err
	}
	var quantityRes QuantityResponse
	err = jsoniter.Unmarshal(bs, &quantityRes)
	if err != nil {
		return QuantityResponse{}, err
	}
	return quantityRes, nil
}
