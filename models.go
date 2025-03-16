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
		MediaID   string `json:"media_id"`
		MediaName string `json:"media_name"`
		Video     Video  `json:"video"`
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

type Video struct {
	AudioCodec string `json:"audio_codec"`
	BitRate    int    `json:"bit_rate"`
	Duration   int    `json:"duration"`
	FrameRate  int    `json:"frame_rate"`
	HdrType    string `json:"hdr_type"`
	Height     int    `json:"height"`
	VideoCodec string `json:"video_codec"`
	VideoType  string `json:"video_type"`
	Width      int    `json:"width"`
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

type UrlFileResponse struct {
	UploadType string `json:"upload_type"`
	URL        URL    `json:"url"`
	Task       Task   `json:"task"`
}

type URL struct {
	Kind string `json:"kind"`
}

type Task struct {
	Kind        string    `json:"kind"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	UserID      string    `json:"user_id"`
	Statuses    []string  `json:"statuses"`
	StatusSize  int       `json:"status_size"`
	Params      Params    `json:"params"`
	FileID      string    `json:"file_id"`
	FileName    string    `json:"file_name"`
	FileSize    string    `json:"file_size"`
	Message     string    `json:"message"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	ThirdTaskID string    `json:"third_task_id"`
	Phase       string    `json:"phase"`
	Progress    int       `json:"progress"`
	IconLink    string    `json:"icon_link"`
	Callback    string    `json:"callback"`
	Space       string    `json:"space"`
}

type Params struct {
	PredictSpeed  string `json:"predict_speed"`
	PredictType   string `json:"predict_type"`
	ThumbnailLink string `json:"thumbnail_link"`
}

type QuantityResponse struct {
	Transfer Transfer `json:"transfer"`
	Data     any      `json:"data"`
	HasMore  bool     `json:"has_more"`
	Base     Base     `json:"base"`
}

type Transfer struct {
	Offline  TransferDetails `json:"offline"`
	Download TransferDetails `json:"download"`
	Upload   TransferDetails `json:"upload"`
}

type TransferDetails struct {
	Info        string `json:"info"`
	TotalAssets int    `json:"total_assets"`
	Assets      int    `json:"assets"`
}

type Base struct {
	UserID        string       `json:"user_id"`
	Info          string       `json:"info"`
	SubStatus     bool         `json:"sub_status"`
	VipStatus     string       `json:"vip_status"`
	ExpireTime    time.Time    `json:"expire_time"`
	Assets        string       `json:"assets"`
	Size          int64        `json:"size"`
	Offline       AssetDetails `json:"offline"`
	Download      AssetDetails `json:"download"`
	Upload        AssetDetails `json:"upload"`
	DownloadDaily AssetDetails `json:"download_daily"`
}

type AssetDetails struct {
	TotalAssets int64 `json:"total_assets"`
	Assets      int64 `json:"assets"`
	Size        int64 `json:"size"`
}

type TaskResponse struct {
	Tasks         []Task `json:"tasks"`
	NextPageToken string `json:"next_page_token"`
	ExpiresIn     int    `json:"expires_in"`
}

type TaskParams struct {
	Age          string `json:"age"`
	MimeType     string `json:"mime_type"`
	PredictSpeed string `json:"predict_speed"`
	PredictType  string `json:"predict_type"`
	URL          string `json:"url"`
}

type ReferenceFile struct {
	Type          string                 `json:"@type"`
	Kind          string                 `json:"kind"`
	ID            string                 `json:"id"`
	ParentID      string                 `json:"parent_id"`
	Name          string                 `json:"name"`
	Size          string                 `json:"size"`
	MimeType      string                 `json:"mime_type"`
	IconLink      string                 `json:"icon_link"`
	Hash          string                 `json:"hash"`
	Phase         string                 `json:"phase"`
	ThumbnailLink string                 `json:"thumbnail_link"`
	Params        map[string]interface{} `json:"params"`
	Space         string                 `json:"space"`
	Medias        []interface{}          `json:"medias"`
	Starred       bool                   `json:"starred"`
	Tags          []string               `json:"tags"`
}
