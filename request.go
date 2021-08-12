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

type Request struct {
	Method, EndPoint, Meta, Include string
	Headers, Params                 map[string]string
	Data                            map[string]interface{}
	Timeout                         time.Duration
	Filters                         map[string]interface{}
	body                            io.Reader
	Skip, Limit                     int32
}

func (lr *Request) AddHeaders(headers map[string]string) {

	for key, value := range headers {
		// will not allow users to change protected headers
		if _, ok := protectedHeaders[key]; ok {
			continue
		}

		lr.Headers[key] = value
	}
}

func (lr *Request) AddParams(params map[string]string) {

	for key, value := range params {
		lr.Params[key] = value
	}
}

func (lr *Request) AddData(data map[string]interface{}) {
	lr.Data = data
}

func (lr *Request) SetSkip(page, limit int32) {
	lr.Skip = page*limit - limit
}

func (lr *Request) SetLimit(limit int32) {
	lr.Limit = limit
}

func (lr *Request) SetInclude(include string) {
	lr.Include = include
}

func (lr *Request) SetIncludeAll() {
	lr.Include = "*"
}

func (lr *Request) prepareGetRequest() {
	queryStr := ""

	for q, v := range lr.Params {
		queryStr = queryStr + url.QueryEscape(q) + "=" + url.QueryEscape(v) + "&"
	}

	for q, v := range lr.Filters {
		queryStr = queryStr + url.QueryEscape(q) + "=" + url.QueryEscape(fmt.Sprintf("%v", v)) + "&"
	}

	if lr.Limit != 0 {
		queryStr = queryStr + url.QueryEscape("limit") + "=" + url.QueryEscape(fmt.Sprintf("%d", lr.Limit)) + "&"
	}

	if lr.Include != "" {
		queryStr = queryStr + url.QueryEscape("include") + "=" + url.QueryEscape(lr.Include) + "&"
	}

	queryStr = queryStr + url.QueryEscape("skip") + "=" + url.QueryEscape(fmt.Sprintf("%d", lr.Skip)) + "&"

	queryStr = strings.TrimRight(queryStr, "&")
	lr.EndPoint = fmt.Sprintf("%s?%s", lr.EndPoint, queryStr)

	lr.AddHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	lr.AddData(nil)
	lr.body = nil
}

func (lr *Request) preparePostRequest() error {
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

func (lr *Request) forgeRequest() (*http.Client, *http.Request, error) {

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

func (lr *Request) Get() (*LucentListResponse, error) {
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

func (lr *Request) Post() (*Response, error) {
	lr.Method = http.MethodPost
	return lr.makePostRequest()
}

func (lr *Request) Put() (*Response, error) {
	lr.Method = http.MethodPut
	return lr.makePostRequest()
}

func (lr *Request) Patch() (*Response, error) {
	lr.Method = http.MethodPatch
	return lr.makePostRequest()
}

func (lr *Request) UploadFromPath(files []string) (*UploadResponse, error) {
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

func (lr *Request) makePostRequest() (*Response, error) {
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

	var response Response
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (lr *Request) make(httpClient *http.Client, request *http.Request) ([]byte, error) {

	resp, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
