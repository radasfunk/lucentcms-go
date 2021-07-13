package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/radasfunk/lucentcmsgo"
	"github.com/radasfunk/lucentcmsgo/dev/utils/env"
)

func init() {
	env.LoadEnv()
}

type Todo struct {
	UserId    int32  `json:"user_id"`
	Id        int32  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")

	dur := time.Duration(5 * time.Second)

	_ = lucentcmsgo.NewLucentClient(channel, token, user, dur)

	httpClient := http.Client{
		Timeout: dur,
	}

	var data = make(map[string]string, 0)

	data["filter[schema]"] = "products"
	data["include"] = "*"

	requestData, err := json.Marshal(data)

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(requestData))

	if err != nil {
		fmt.Printf("error %v", err.Error())
	}

	request.Header.Add("Accept", "application/json")
	// request.Header.Add("Lucent-Channel", channel)
	// request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Accept-Language", "el-GR")

	resp, err := httpClient.Do(request)

	if err != nil {
		fmt.Printf("error from response %v", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error reading ioutil %v", err.Error())
	}

	// body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) //

	var response Todo

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Printf("error in json response %v", err.Error())
	}
	fmt.Println(response.Id, response.Title, "response body")
}

func aTestRunThatWorked() {
	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")

	dur := time.Duration(5 * time.Second)

	_ = lucentcmsgo.NewLucentClient(channel, token, user, dur)

	httpClient := http.Client{
		Timeout: dur,
	}

	var data = make(map[string]string, 0)

	data["filter[schema]"] = "products"
	data["include"] = "*"

	requestData, err := json.Marshal(data)

	request, err := http.NewRequest("GET", "https://api.lucentcms.com/api/documents", bytes.NewBuffer(requestData))

	if err != nil {
		fmt.Printf("error %v", err.Error())
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Lucent-Channel", channel)
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Accept-Language", "el-GR")

	resp, err := httpClient.Do(request)

	if err != nil {
		fmt.Printf("error from response %v", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error reading ioutil %v", err.Error())
	}

	// body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) //

	var response interface{}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Printf("error in json response %v", err.Error())
	}
	fmt.Println(response, "response body")
}
