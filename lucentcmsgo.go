package lucentcmsgo

import (
	"fmt"
)

const (
	lucentBaseUrl = "https://api.lucentcms.com/api/"
)

type LucentClient struct {
	Channel, Token, LucentUser, BaseUrl string
	DefaultHeaders                      map[string]string
}

// Creates a new lucent struct
// Recommend to use NewLucentClient instead of populating the fields
func NewLucentClient(channel, token, lucentUser string) *LucentClient {
	headers := make(map[string]string)

	headers["Accept"] = "application/json"
	headers["Lucent-Channel"] = channel
	headers["Authorization"] = "Bearer " + token

	if lucentUser != "" {
		headers["Lucent-User"] = lucentUser
	}

	fmt.Printf("Lucent headers %v\n", headers)

	lucentClient := &LucentClient{
		Channel:        channel,
		Token:          token,
		LucentUser:     lucentUser,
		DefaultHeaders: headers,
		BaseUrl:        lucentBaseUrl,
	}

	return lucentClient
}

func (lc *LucentClient) NewRequest(method, endpoint string, data ...interface{}) *LucentRequest {
	fmt.Printf("create a new request")

	req := &LucentRequest{
		Method:   method,
		EndPoint: endpoint,
		Data:     data,
		// add the default headers
		Headers: lc.DefaultHeaders,
	}

	return req
}
