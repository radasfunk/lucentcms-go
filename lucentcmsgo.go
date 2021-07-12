package lucentcmsgo

import (
	"fmt"
	"net/url"
)

const (
	lucentBaseUrl = "https://api.lucentcms.com/api/"
)

var (
	validMethods = map[string]bool{
		"GET":    true,
		"POST":   true,
		"PUT":    true,
		"PATCH":  true,
		"DELETE": true,
		"UPLOAD": true,
	}

	validEndpoints = map[string]bool{
		"documents":  true,
		"documents/": true,
		"channels":   true,
		"channels/":  true,
		"files":      true,
		"files/":     true,
	}

	protectedHeaders = map[string]bool {
		"Lucent-Channel" : true,
		"Lucent-User": true,
	}
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

	lucentClient := &LucentClient{
		Channel:        channel,
		Token:          token,
		LucentUser:     lucentUser,
		DefaultHeaders: headers,
		BaseUrl:        lucentBaseUrl,
	}

	return lucentClient
}

func (lc *LucentClient) NewRequest(method, endpoint string, data ...interface{}) (*LucentRequest, error) {

	if _, ok := validMethods[method]; !ok {
		return nil, fmt.Errorf("unsupported method. can not create request %v", method)
	}

	if _, ok := validEndpoints[endpoint]; !ok {
		return nil, fmt.Errorf("unsupported out of scope. can not create request endpoint %v", endpoint)
	}

	endpoint = lc.BaseUrl + endpoint

	_, err := url.ParseRequestURI(endpoint)

	if err != nil {
		return nil, err
	}

	req := &LucentRequest{
		Method:   method,
		EndPoint: endpoint,
		Data:     data,
		Headers:  lc.DefaultHeaders,
	}

	return req, nil
}
