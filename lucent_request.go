package lucentcmsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type LucentRequest struct {
	Method, EndPoint string
	Headers          map[string]string
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

func (lr *LucentRequest) AddData(data interface{}) {
	lr.Data = data
}

func (lr *LucentRequest) Send() (*LucentResponse, error) {
	// use prepare()

	lr.AddHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	fmt.Println(lr.Data)

	requestData, err := json.Marshal(lr.Data)

	if err != nil {
		fmt.Printf("1 %v\n", err.Error())
		return nil, err
	}

	httpClient := http.Client{
		Timeout: lr.Timeout,
	}
	// request
	request, err := http.NewRequest(lr.Method, lr.EndPoint, bytes.NewBuffer(requestData))

	if err != nil {
		fmt.Printf("2 %v\n", err.Error())

		return nil, err
	}

	for k, v := range lr.Headers {
		request.Header.Set(k, v)
	}

	resp, err := httpClient.Do(request)

	if err != nil {
		fmt.Printf("3 %v\n", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("4 %v\n", err.Error())

		return nil, err
	}

	// fmt.Println(string(body))
	var lucentResponse LucentResponse
	err = json.Unmarshal(body, &lucentResponse)

	if err != nil {
		fmt.Printf("5 %v\n", err.Error())
		return nil, err
	}

	fmt.Printf("%v\n", lucentResponse.Data)

	return &lucentResponse, nil
}
