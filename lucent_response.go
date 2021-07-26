package lucentcmsgo

import (
	"fmt"
	"time"
)

type LucentResponse struct {
	Data           []Document
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
	Channel       string      `json:"channel"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	CreatedAt     time.Time   `json:"createdAt"`
}

func makeResponse() {
	fmt.Printf("Get response")
}

func (lrr *LucentResponse) HasErrors() bool {
	return len(lrr.Errors) > 0
}

func (lrr *LucentResponse) GetData() []Document {
	return lrr.Data
}

func (lrr *LucentResponse) First() (Document, bool) {
	if len(lrr.Data) > 0 {
		return lrr.Data[0], true
	}
	return Document{}, false
}

func (lrr *LucentResponse) GetIncluded() map[string]interface{} {
	return lrr.Included
}

// returns all the errors
func (lrr *LucentResponse) GetErrors() []string {
	return lrr.Errors
}

// returns the first error
func (lrr *LucentResponse) Error() string {
	if lrr.HasErrors() {
		return lrr.Errors[0]
	}

	return ""
}
