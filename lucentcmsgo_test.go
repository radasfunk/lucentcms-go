package lucentcmsgo

import (
	"testing"
	"time"

	"github.com/radasfunk/lucentcmsgo/utils/env"
)

func init() {
	env.LoadEnv()
}

func TestNewLucentClientIsCreatedWithExpectedValue(t *testing.T) {

	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")
	locale := env.Get("LUCENTV3_LOCALE")

	dur := time.Duration(5 * time.Second)

	client := NewLucentClient(channel, token, user, locale, dur)

	if client.Channel != channel {
		t.Errorf("channel name got %v, want %v", client.Channel, channel)
	}

	if client.Token != token {
		t.Errorf("channel token got %v, want %v", client.Token, token)
	}

	if client.LucentUser != user {
		t.Errorf("channel lucent user got %v, want %v", client.LucentUser, user)
	}

	// make sure headers exists
	headers := make(map[string]string)

	headers["Accept"] = "application/json"
	headers["Lucent-Channel"] = channel
	headers["Authorization"] = "Bearer " + token
	headers["Lucent-User"] = user

	expectedHeadersCount := 5
	receivedHeadersCount := 0

	for key, value := range headers {
		// check if key exists
		if v, ok := client.DefaultHeaders[key]; !ok {
			t.Errorf("channel default header %v does not exist", key)
		} else if v != value {
			t.Errorf("channel default header value does not match. got %v want %v", value, v)
		}
	}

	for _ = range client.DefaultHeaders {
		receivedHeadersCount++
	}

	if expectedHeadersCount != receivedHeadersCount {
		t.Errorf("channel default headers count mistammatch got %v want %v", receivedHeadersCount, expectedHeadersCount)
	}
}

func TestValidURLMethodAreBeingCreated(t *testing.T) {
	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")

	dur := time.Duration(5 * time.Second)
	locale := env.Get("LUCENTV3_LOCALE")

	client := NewLucentClient(channel, token, user, locale, dur)

	// TODO need to update
	checklist := map[string]bool{
		"documents/":                     true,
		"files":                          true,
		"channels":                       true,
		"":                               false,
		"asd":                            false,
		"123":                            false,
		"https://api.lucentcms.com/api/": false,
	}

	for url, expected := range checklist {
		_, err := client.NewRequest(url, nil)

		if expected == false && err == nil {
			t.Errorf("expected %v for url %v got %v", expected, url, nil)
		}

		if expected == true && err != nil {
			t.Errorf("expected %v for url %v got %v", expected, url, nil)
		}

	}
}
