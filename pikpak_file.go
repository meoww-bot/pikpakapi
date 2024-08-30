package pikpakapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

func (p *PikPak) GetFolderFileStatList(parentId string) ([]FileStat, error) {
	filters := `{"trashed":{"eq":false}}`
	query := url.Values{}
	query.Add("thumbnail_size", "SIZE_MEDIUM")
	query.Add("limit", "500")
	query.Add("parent_id", parentId)
	query.Add("with_audit", "false")
	query.Add("filters", filters)
	fileList := make([]FileStat, 0)

	for {
		req, err := http.NewRequest("GET", "https://api-drive.mypikpak.com/drive/v1/files?"+query.Encode(), nil)
		if err != nil {
			return fileList, err
		}
		req.Header.Set("X-Captcha-Token", p.CaptchaToken)
		req.Header.Set("Content-Type", "application/json")
		bs, err := p.sendWithErrHandle(req)
		if err != nil {
			return fileList, err
		}
		var result fileListResult
		err = json.Unmarshal(bs, &result)
		if err != nil {
			return fileList, err
		}
		fileList = append(fileList, result.Files...)
		if result.NextPageToken == "" {
			break
		}
		query.Set("page_token", result.NextPageToken)
	}
	return fileList, nil
}

// Find FileState similar to name in the parentId directory
func (p *PikPak) GetFileStat(parentId string, name string) (FileStat, error) {
	stats, err := p.GetFolderFileStatList(parentId)
	if err != nil {
		return FileStat{}, err
	}
	for _, stat := range stats {
		if stat.Name == name {
			return stat, nil
		}
	}
	return FileStat{}, errors.New("file not found")
}

func (p *PikPak) GetFile(fileId string) (File, error) {
	var fileInfo File
	query := url.Values{}
	query.Add("thumbnail_size", "SIZE_MEDIUM")
	req, err := http.NewRequest("GET", "https://api-drive.mypikpak.com/drive/v1/files/"+fileId+"?"+query.Encode(), nil)
	if err != nil {
		return fileInfo, nil
	}
	req.Header.Set("X-Captcha-Token", p.CaptchaToken)
	req.Header.Set("X-Device-Id", p.DeviceId)
	bs, err := p.sendWithErrHandle(req)
	if err != nil {
		return fileInfo, nil
	}

	err = json.Unmarshal(bs, &fileInfo)
	if err != nil {
		return fileInfo, err
	}
	return fileInfo, err
}
