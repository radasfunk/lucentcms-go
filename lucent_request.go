package lucentcmsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type LucentRequest struct {
	Method, EndPoint string
	Headers, Params  map[string]string
	Data             map[string]interface{}
	Timeout          time.Duration
	body             io.Reader
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

func (lr *LucentRequest) AddData(data map[string]interface{}) {
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

	lr.AddData(nil)
	lr.body = nil
}

func (lr *LucentRequest) preparePostRequest() error {
	data, err := json.Marshal(lr.Data)

	if err != nil {
		return err
	}

	lr.AddHeaders(map[string]string{
		"Content-Type": "application/json",
		// "Content-Length": strconv.Itoa(len(data)),
	})

	formData := bytes.NewBuffer(data)

	lr.body = formData

	return nil
}

func (lr *LucentRequest) prepareRequest() (*http.Client, *http.Request, error) {

	// var rData interface{}

	switch lr.Method {
	case "GET", "DELETE":
		lr.prepareGetRequest()
	case "POST", "PUT", "PATCH":
		// rData = "method=post"
		lr.preparePostRequest()
	case "UPLOAD":
		fmt.Printf("handle upload data")
		// rData = "method=upload"
	}

	return lr.forgeRequest()
}

func (lr *LucentRequest) forgeRequest() (*http.Client, *http.Request, error) {

	httpClient := http.Client{
		Timeout: lr.Timeout,
	}

	request, err := http.NewRequest(lr.Method, lr.EndPoint, lr.body)

	if err != nil {
		return nil, nil, err
	}

	for k, v := range lr.Headers {
		request.Header.Set(k, v)
	}

	return &httpClient, request, nil
}

func (lr *LucentRequest) Post() (*LucentResponse, error) {
	err := lr.preparePostRequest()

	if err != nil {
		return nil, err
	}

	httpClient, request, err := lr.forgeRequest()

	if err != nil {
		return nil, err
	}

	bytes, err := lr.make(httpClient, request)

	if err != nil {
		return nil, err
	}

	var response LucentResponse
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		return nil, err
	}

	fmt.Println(string(bytes))

	return &response, nil
}

func (lr *LucentRequest) make(httpClient *http.Client, request *http.Request) ([]byte, error) {

	resp, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (lr *LucentRequest) Send() (*LucentListResponse, error) {

	httpClient, request, err := lr.prepareRequest()

	if err != nil {
		return nil, err
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

	var LucentListResponse LucentListResponse
	err = json.Unmarshal(body, &LucentListResponse)

	if err != nil {
		return nil, err
	}

	return &LucentListResponse, nil
}
