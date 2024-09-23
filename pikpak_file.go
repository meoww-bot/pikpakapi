package pikpakapi

import (
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
)

// Get file information based on directory ID
func (p *PikPak) GetDirFilesStat(parentID string) ([]FileStat, error) {
	return p.lookupDirFilesStat(parentID)
}

// Get file information based on directory path
func (p *PikPak) GetDirFilesStatByPath(path Path) ([]FileStat, error) {
	id, err := p.GetSubDirID("", path)
	if err != nil {
		return nil, err
	}
	return p.GetDirFilesStat(id)
}

func (p *PikPak) lookupDirFilesStat(parentID string) ([]FileStat, error) {
	query := url.Values{}
	// Avoid searching for trash
	filters := `{"phase":{"eq":"PHASE_TYPE_COMPLETE"},"trashed":{"eq":false}}`
	query.Add("thumbnail_size", "SIZE_MEDIUM")
	query.Add("limit", "500")
	query.Add("parent_id", parentID)
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
		bs, err := p.sendWithErrHandle(req, nil)
		if err != nil {
			return fileList, err
		}
		var result fileListResult
		err = jsoniter.Unmarshal(bs, &result)
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

func (p *PikPak) GetFile(fileID string) (info File, err error) {
	query := url.Values{}
	query.Add("thumbnail_size", "SIZE_MEDIUM")
	req, err := http.NewRequest("GET", "https://api-drive.mypikpak.com/drive/v1/files/"+fileID+"?"+query.Encode(), nil)
	if err != nil {
		return info, err
	}
	req.Header.Set("X-Captcha-Token", p.CaptchaToken)
	bs, err := p.sendWithErrHandle(req, nil)
	if err != nil {
		return info, nil
	}

	err = jsoniter.Unmarshal(bs, &info)
	if err != nil {
		return info, err
	}
	return
}

// Find FileState similar to name in the parentId directory
func (p *PikPak) GetFileStatByPath(path Path) (stat FileStat, err error) {
	parent, name := path.Parent(), path.Name()
	parentID, err := p.GetDirID(parent)
	if err != nil {
		return stat, err
	}
	stats, err := p.GetDirFilesStat(parentID)
	if err != nil {
		return FileStat{}, err
	}
	for _, stat := range stats {
		if stat.Name == name {
			return stat, nil
		}
	}
	return stat, ErrNotFoundFile
}
