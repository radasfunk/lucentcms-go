package lucentcmsgo

import (
	"testing"

	"github.com/radasfunk/lucentcmsgo/dev/utils/env"
)

var client *LucentClient

func init() {
	env.LoadEnv()

	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")

	client = NewLucentClient(channel, token, user)
}

func TestHeadersCanBeAdded(t *testing.T) {

	req, _ := client.NewRequest("GET", "documents")

	headers := map[string]string{
		"Custom-Header":   "Custom-Value",
		"Accept-Encoding": "gzip",
	}

	req.AddHeaders(headers)

	for k,v := range headers {
		if val, ok := req.Headers[k]; ! ok {
			t.Errorf("expected %v to be present",k) // empty
		} else if val != v {
			t.Errorf("expected %v got %v",v,val) // not empty
		} 
	}
}
