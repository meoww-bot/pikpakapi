package pikpakapi

import (
	"bytes"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

func (p *PikPak) DeleteBatchFiles(ids ...string) error {
	url := "https://api-drive.mypikpak.com/drive/v1/files:batchTrash"
	m := map[string]any{
		"ids": ids,
	}
	bs, err := jsoniter.Marshal(&m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bs))
	if err != nil {
		return err
	}
	bs, err = p.sendRequest(req)
	if err != nil {
		return err
	}
	error_code := gjson.GetBytes(bs, "error_code").Int()
	if error_code != 0 {
		// refresh token failed
		if error_code == 4126 {
			// 重新登录
			return p.Login()
		}
		return fmt.Errorf("delete file error: %v", gjson.GetBytes(bs, "error").Int())
	}
	return nil
}
