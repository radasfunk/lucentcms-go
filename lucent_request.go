package lucentcmsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LucentRequest struct {
	Method, EndPoint, Meta string
	Headers, Params        map[string]string
	Data                   map[string]interface{}
	Timeout                time.Duration
	Filters                map[string]interface{}
	body                   io.Reader
	Skip, Limit            int32
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

func (lr *LucentRequest) FilterWhere(key string, value interface{}) {
	key = "filter[" + key + "]"
	lr.Filters[key] = value
}

func (lr *LucentRequest) FilterOrWhere(key string, value interface{}) {
	key = "filter[or][" + key + "]"
	lr.Filters[key] = value
}

func (lr *LucentRequest) SetSkip(page, limit int32) {
	lr.Skip = page*limit - limit
}

func (lr *LucentRequest) SetLimit(limit int32) {
	lr.Limit = limit
}

func (lr *LucentRequest) SetMeta(meta string) {
	lr.Meta = meta
}

func (lr *LucentRequest) prepareGetRequest() {
	queryStr := ""

	for q, v := range lr.Params {
		queryStr = queryStr + url.QueryEscape(q) + "=" + url.QueryEscape(v) + "&"
	}

	for q, v := range lr.Filters {
		queryStr = queryStr + url.QueryEscape(q) + "=" + url.QueryEscape(fmt.Sprintf("%v", v)) + "&"
	}

	queryStr = queryStr + url.QueryEscape("skip") + "=" + url.QueryEscape(fmt.Sprintf("%d", lr.Skip)) + "&"
	queryStr = queryStr + url.QueryEscape("limit") + "=" + url.QueryEscape(fmt.Sprintf("%d", lr.Limit)) + "&"

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
	})

	formData := bytes.NewBuffer(data)

	lr.body = formData

	return nil
}

// deprecated
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

func (lr *LucentRequest) Get() (*LucentListResponse, error) {
	lr.Method = http.MethodGet

	lr.prepareGetRequest()

	httpClient, request, err := lr.forgeRequest()

	if err != nil {
		return nil, err
	}

	bytes, err := lr.make(httpClient, request)

	if err != nil {
		return nil, err
	}

	var response LucentListResponse
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (lr *LucentRequest) Post() (*LucentResponse, error) {
	lr.Method = http.MethodPost
	return lr.makePostRequest()
}

func (lr *LucentRequest) Put() (*LucentResponse, error) {
	lr.Method = http.MethodPut
	return lr.makePostRequest()
}

func (lr *LucentRequest) Patch() (*LucentResponse, error) {
	lr.Method = http.MethodPatch
	return lr.makePostRequest()
}

func (lr *LucentRequest) UploadFromDisk(files []string) (*UploadResponse, error) {
	lr.Method = http.MethodPost

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for i, path := range files {

		file, err := os.Open(path)

		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(fmt.Sprintf("files[%d]", i), filepath.Base(path))

		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
	}

	err := writer.Close()

	if err != nil {
		return nil, err
	}

	lr.body = bytes.NewBuffer([]byte(fmt.Sprintf("%v", body)))

	lr.AddHeaders(map[string]string{
		"Content-Type": writer.FormDataContentType(),
	})

	httpClient, request, err := lr.forgeRequest()

	if err != nil {
		return nil, err
	}

	bytes, err := lr.make(httpClient, request)

	if err != nil {
		return nil, err
	}

	var response UploadResponse
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (lr *LucentRequest) makePostRequest() (*LucentResponse, error) {
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

// deprecated
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
