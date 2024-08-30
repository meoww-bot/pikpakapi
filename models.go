package pikpakapi

import (
	"time"
)

type FileStat struct {
	Kind          string    `json:"kind"`
	ID            string    `json:"id"`
	ParentID      string    `json:"parent_id"`
	Name          string    `json:"name"`
	UserID        string    `json:"user_id"`
	Size          string    `json:"size"`
	FileExtension string    `json:"file_extension"`
	MimeType      string    `json:"mime_type"`
	CreatedTime   time.Time `json:"created_time"`
	ModifiedTime  time.Time `json:"modified_time"`
	IconLink      string    `json:"icon_link"`
	ThumbnailLink string    `json:"thumbnail_link"`
	Md5Checksum   string    `json:"md5_checksum"`
	Hash          string    `json:"hash"`
	Phase         string    `json:"phase"`
}

type File struct {
	FileStat
	Revision       string `json:"revision"`
	Starred        bool   `json:"starred"`
	WebContentLink string `json:"web_content_link"`
	Links          struct {
		ApplicationOctetStream struct {
			URL    string    `json:"url"`
			Token  string    `json:"token"`
			Expire time.Time `json:"expire"`
		} `json:"application/octet-stream"`
	} `json:"links"`
	Audit struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Title   string `json:"title"`
	} `json:"audit"`
	Medias []struct {
		MediaID   string      `json:"media_id"`
		MediaName string      `json:"media_name"`
		Video     interface{} `json:"video"`
		Link      struct {
			URL    string    `json:"url"`
			Token  string    `json:"token"`
			Expire time.Time `json:"expire"`
		} `json:"link"`
		NeedMoreQuota  bool          `json:"need_more_quota"`
		VipTypes       []interface{} `json:"vip_types"`
		RedirectLink   string        `json:"redirect_link"`
		IconLink       string        `json:"icon_link"`
		IsDefault      bool          `json:"is_default"`
		Priority       int           `json:"priority"`
		IsOrigin       bool          `json:"is_origin"`
		ResolutionName string        `json:"resolution_name"`
		IsVisible      bool          `json:"is_visible"`
		Category       string        `json:"category"`
	} `json:"medias"`
	Trashed     bool   `json:"trashed"`
	DeleteTime  string `json:"delete_time"`
	OriginalURL string `json:"original_url"`
	Params      struct {
		Platform     string `json:"platform"`
		PlatformIcon string `json:"platform_icon"`
	} `json:"params"`
	OriginalFileIndex int           `json:"original_file_index"`
	Space             string        `json:"space"`
	Apps              []interface{} `json:"apps"`
	Writable          bool          `json:"writable"`
	FolderType        string        `json:"folder_type"`
	Collection        interface{}   `json:"collection"`
}

type fileListResult struct {
	NextPageToken string     `json:"next_page_token"`
	Files         []FileStat `json:"files"`
}

// Quota response
type quotaResponse struct {
	Kind      string `json:"kind"`
	Quota     Quota  `json:"quota"`
	ExpiresAt string `json:"expires_at"`
	Quotas    Quotas `json:"quotas"`
}

type Quota struct {
	Kind           string `json:"kind"`
	Limit          string `json:"limit"`
	Usage          string `json:"usage"`
	UsageInTrash   string `json:"usage_in_trash"`
	PlayTimesLimit string `json:"play_times_limit"`
	PlayTimesUsage string `json:"play_times_usage"`
}

type Quotas struct{}
