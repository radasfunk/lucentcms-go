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
