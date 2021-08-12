package lucentcmsgo

import (
	"time"
)

type Response struct {
	Data Document `json:"data"`
	baseResponse
}

type LucentListResponse struct {
	Data []Document `json:"data"`
	baseResponse
}

type UploadResponse struct {
	Data []File `json:"data"`
	baseResponse
}

type baseResponse struct {
	Errors, Links  []string
	Meta, Included map[string]interface{} // TODO add paginator
}

type Content map[string]interface{}

type Document struct {
	ID            string      `json:"id"`
	RequestLocale string      `json:"requestLocale"`
	Locale        string      `json:"locale"`
	Schema        string      `json:"schema"`
	Creator       string      `json:"creator"`
	Editor        string      `json:"editor"`
	Status        string      `json:"status"`
	Version       int         `json:"version"`
	PublishedAt   time.Time   `json:"publishedAt"`
	Behind        bool        `json:"behind"`
	Content       Content     `json:"content"`
	Subdocs       interface{} `json:"subdocs"`
	Relationships interface{} `json:"relationships"`
	Channel       string      `json:"channel"`
	Resource      string      `json:"resource"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	CreatedAt     time.Time   `json:"createdAt"`
}

type File struct {
	ID           string      `json:"id"`
	OriginalName string      `json:"originalName"`
	Filename     string      `json:"filename"`
	Path         string      `json:"path"`
	Mime         string      `json:"mime"`
	URL          string      `json:"url"`
	Image        string      `json:"image"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Size         int         `json:"size"`
	Width        int         `json:"width"`
	Height       int         `json:"height"`
	Alt          string      `json:"alt"`
	Credits      interface{} `json:"credits"`
	Checksum     string      `json:"checksum"`
	Copyright    string      `json:"copyright"`
	Tags         []string    `json:"tags"`
	UploaderID   string      `json:"uploaderId"`
	Channel      string      `json:"channel"`
	UpdatedAt    time.Time   `json:"updatedAt"`
	CreatedAt    time.Time   `json:"createdAt"`
	Resource     string      `json:"resource"`
}

func (lrr *baseResponse) HasErrors() bool {
	return len(lrr.Errors) > 0
}

func (lrr *LucentListResponse) GetData() []Document {
	return lrr.Data
}

func (lrr *Response) GetData() Document {
	return lrr.Data
}

func (lrr *LucentListResponse) First() (Document, bool) {
	if len(lrr.Data) > 0 {
		return lrr.Data[0], true
	}
	return Document{}, false
}

func (lrr *baseResponse) GetIncluded() map[string]interface{} {
	return lrr.Included
}

// returns all the errors
func (lrr *baseResponse) GetErrors() []string {
	return lrr.Errors
}

// returns the first error
func (lrr *baseResponse) Error() string {
	if lrr.HasErrors() {
		return lrr.Errors[0]
	}

	return ""
}
