package lucentcmsgo

import (
	"testing"

	"github.com/radasfunk/lucentcmsgo/dev/utils/env"
)

func init() {
	env.LoadEnv()
}

// make sure a instance of lucentClient is being created
func TestNewLucentClientIsCreated(t *testing.T) {

	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")

	client := NewLucentClient(channel, token, user)

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

	expectedHeadersCount := 4
	receivedHeadersCount := 0

	for key, value := range headers {
		// check if key exists
		if v, ok := client.DefaultHeaders[key]; !ok {
			t.Errorf("channel default header %v does not exist", key)
		} else if v != value {
			t.Errorf("channel default header value does not match. got %v want %v", value, v)
		}
	}

	for _, _ = range client.DefaultHeaders {
		receivedHeadersCount++
	}

	if expectedHeadersCount != receivedHeadersCount {
		t.Errorf("channel default headers count mistammatch got %v want %v", receivedHeadersCount, expectedHeadersCount)
	}
}
