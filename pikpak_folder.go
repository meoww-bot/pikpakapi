package pikpakapi

import (
	"bytes"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

// Get the ID of the directory, but you need to give the ID of the first directory.
func (p *PikPak) GetSubDirID(parentID string, dir Path) (string, error) {
	for _, dir := range dir.Cut() {
		id, err := p.lookupDirID(parentID, dir)
		if err != nil {
			return "", err
		}
		parentID = id
	}
	return parentID, nil
}

// Get the ID of a directory without giving the initial directory ID.
func (p *PikPak) GetDirID(dir Path) (string, error) {
	return p.GetSubDirID("", dir)
}

// Look up the ID in the parentId directory that match the dirName name
func (p *PikPak) lookupDirID(parentID string, dirName string) (string, error) {
	// Check the cache, if the folder id is in the cache, then return it
	if id, ok := p.cache.Get(newTuple(parentID, dirName)); ok {
		return id, nil
	}

	value := url.Values{}
	value.Add("parent_id", parentID)
	value.Add("page_token", "")
	value.Add("with_audit", "false")
	value.Add("thumbnail_size", "SIZE_LARGE")
	value.Add("limit", "500")
	// Avoid searching for trash
	value.Add("filters", `{"phase":{"eq":"PHASE_TYPE_COMPLETE"},"trashed":{"eq":false}}`)
	for {
		req, err := http.NewRequest("GET", "https://api-drive.mypikpak.com/drive/v1/files?"+value.Encode(), nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("Country", "CN")
		req.Header.Set("X-Peer-Id", p.DeviceId)
		req.Header.Set("X-User-Region", "1")
		req.Header.Set("X-Alt-Capability", "3")
		req.Header.Set("X-Client-Version-Code", "10083")
		req.Header.Set("X-Captcha-Token", p.CaptchaToken)
		bs, err := p.sendWithErrHandle(req, nil)
		if err != nil {
			return "", err
		}
		files := gjson.GetBytes(bs, "files").Array()

		for _, file := range files {
			kind := file.Get("kind").String()
			name := file.Get("name").String()
			trashed := file.Get("trashed").Bool()
			if kind == KIND_FOLDER && name == dirName && !trashed {
				id := file.Get("id").String()
				// Setting the cache
				p.cache.Set(newTuple(parentID, dirName), id)
				return id, nil
			}
		}
		nextToken := gjson.GetBytes(bs, "next_page_token").String()
		if nextToken == "" {
			break
		}
		value.Set("page_token", nextToken)
	}
	return "", ErrNotFoundFolder
}

func (p *PikPak) CreateSubDir(parentID string, dir Path) (string, error) {
	for _, dir := range dir.Cut() {
		id, err := p.lookupDirID(parentID, dir)
		if err != nil {
			if err == ErrNotFoundFolder {
				createID, err := p.createDir(parentID, dir)
				if err != nil {
					return "", err
				}
				parentID = createID
				continue
			}
			return "", err
		}
		parentID = id
	}
	return parentID, nil
}

func (p *PikPak) CreateDir(dir Path) (string, error) {
	return p.CreateSubDir("", dir)
}

// Create new folder in parent folder
// parentId is parent folder id
func (p *PikPak) createDir(parentID, dir string) (string, error) {
	m := map[string]interface{}{
		"kind":      KIND_FOLDER,
		"parent_id": parentID,
		"name":      dir,
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
	req.Header.Set("Product_flavor_name", "cha")
	req.Header.Set("X-Captcha-Token", p.CaptchaToken)
	req.Header.Set("X-Client-Version-Code", "10083")
	req.Header.Set("X-Peer-Id", p.DeviceId)
	req.Header.Set("X-User-Region", "1")
	req.Header.Set("X-Alt-Capability", "3")
	req.Header.Set("Country", "CN")
	bs, err = p.sendWithErrHandle(req, bs)
	if err != nil {
		return "", err
	}
	id := gjson.GetBytes(bs, "file.id").String()
	// Setting the cache
	p.cache.Set(newTuple(parentID, dir), id)
	return id, nil
}
