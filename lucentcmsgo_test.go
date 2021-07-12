package lucentcmsgo

import (
	"fmt"
	"testing"

	"github.com/radasfunk/lucentcmsgo/dev/utils/env"
)

func init() {
	env.LoadEnv()
}

func TestNewLucentClientIsCreatedWithExpectedValue(t *testing.T) {

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

func TestValidMethodsAreAcceptedWhileCreatingRequest(t *testing.T) {
	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")

	client := NewLucentClient(channel, token, user)

	method := "INVALID"
	_, err := client.NewRequest(method, "documents/")

	expected := fmt.Sprintf("unsupported method. can not create request %v", method)

	if err == nil {
		t.Errorf("allows invalid request method. expected %v got %v", expected, err.Error())
	}

	methods := []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "UPLOAD",
	}

	for _, m := range methods {
		_, err = client.NewRequest(m, "documents/")

		if err != nil {
			t.Errorf("expected error to be %v got %v", nil, err.Error())
		}
	}
}

func TestValidURLMethodAreBeingCreated(t *testing.T) {
	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")

	client := NewLucentClient(channel, token, user)

	// TODO need to update
	checklist := map[string]bool{
		"documents/":                 true,
		"files":                      true,
		"channels":                   true,
		"":                           false,
		"asd":                        false,
		"123":                        false,
		"https://api.lucentcms.com/": false,
	}

	for url, expected := range checklist {
		_, err := client.NewRequest("GET", url)

		if expected == false && err == nil {
			t.Errorf("expected %v for url %v got %v", expected, url, nil)
		}

		if expected == true && err != nil {
			t.Errorf("expected %v for url %v got %v", expected, url, nil)
		}

	}
}
