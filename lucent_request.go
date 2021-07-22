package lucentcmsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type LucentRequest struct {
	Method, EndPoint string
	Headers, Params  map[string]string
	Data             interface{}
	Timeout          time.Duration
}

func (lr *LucentRequest) AddHeaders(headers map[string]string) {

	for key, value := range headers {
		// will not allow users to change protected headers
		if _, ok := protectedHeaders[key]; ok {
			continue
		}

		lr.Headers[key] = value
	}
}

func (lr *LucentRequest) AddParams(params map[string]string) {

	for key, value := range params {
		lr.Params[key] = value
	}
}

func (lr *LucentRequest) AddData(data interface{}) {
	lr.Data = data
}

func (lr *LucentRequest) prepareGetRequest() {
	queryStr := ""

	for q, v := range lr.Params {
		queryStr = queryStr + url.QueryEscape(q) + "=" + url.QueryEscape(v) + "&"
	}

	queryStr = strings.TrimRight(queryStr, "&")
	lr.EndPoint = fmt.Sprintf("%s?%s", lr.EndPoint, queryStr)

	lr.AddHeaders(map[string]string{
		"Content-Type": "application/json",
	})
}

func (lr *LucentRequest) Send() (*LucentResponse, error) {
	// use prepare()

	var rData interface{}

	switch lr.Method {
	case "GET", "DELETE":
		lr.prepareGetRequest()
	case "POST", "PUT", "PATCH":
		rData = "method=post"
	case "UPLOAD":
		rData = "method=upload"
	}

	fmt.Print(rData, lr.EndPoint)

	requestData, err := json.Marshal(lr.Data)

	if err != nil {
		return nil, err
	}

	httpClient := http.Client{
		Timeout: lr.Timeout,
	}
	// request
	request, err := http.NewRequest(lr.Method, lr.EndPoint, bytes.NewBuffer(requestData))

	if err != nil {
		return nil, err
	}

	for k, v := range lr.Headers {
		request.Header.Set(k, v)
	}

	resp, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var lucentResponse LucentResponse
	_ = json.Unmarshal(body, &lucentResponse)

	return &lucentResponse, nil
}
