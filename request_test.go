package lucentcmsgo

import (
	"testing"
	"time"

	"github.com/radasfunk/lucentcmsgo/utils/env"
)

var (
	client        *LucentClient
	user, channel string
)

func init() {
	env.LoadEnv()

	channel = env.Get("LUCENTV3_CHANNEL")
	user = env.Get("LUCENTV3_USER")

	token := env.Get("LUCENTV3_TOKEN")
	locale := env.Get("LUCENTV3_LOCALE")

	dur := time.Duration(5 * time.Second)
	client = NewLucentClient(channel, token, user, locale, dur)
}

func TestHeadersCanBeAdded(t *testing.T) {

	req, _ := client.NewRequest("documents", nil)

	headers := map[string]string{
		"Custom-Header":   "Custom-Value",
		"Accept-Encoding": "gzip",
	}

	req.AddHeaders(headers)

	for k, v := range headers {
		if val, ok := req.Headers[k]; !ok {
			t.Errorf("expected %v to be present", k) // empty
		} else if val != v {
			t.Errorf("expected %v got %v", v, val) // not empty
		}
	}
}

func TestProtectedHeadersCanNotBeChanged(t *testing.T) {

	req, _ := client.NewRequest("documents", nil)

	headers := map[string]string{
		"Lucent-Channel": "UPDATED_LUCENT_CHANNEL",
		"Lucent-User":    "UPDATED_RANDOM_USER",
	}

	req.AddHeaders(headers)

	for k, v := range headers {
		if val, ok := req.Headers[k]; !ok {
			t.Errorf("expected %v to be present", k) // empty
		} else if val == v {
			t.Errorf("expected %v not to be equal %v", v, val) // not empty
		}
	}
}
