package pikpakapi

import (
	"bytes"
	"net/http"

	jsoniter "github.com/json-iterator/go"
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
	bs, err = p.sendWithErrHandle(req, bs)
	if err != nil {
		return err
	}
	return nil
}
