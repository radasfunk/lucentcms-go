package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type AResponse struct {
	Data           []Document
	Errors, Links  []string
	Meta, Included map[string]interface{} // TODO add paginator
}

type Content map[string]interface{}

type Document struct {
	ID            string      `json:"id"`
	RequestLocale string      `json:"requestLocale"`
	Locale        string      `json:"locale"`
	Schema        string      `json:"schema"`
	Creator       string      `json:"creator"`
	Editor        string      `json:"editor"`
	Status        string      `json:"status"`
	Version       int         `json:"version"`
	PublishedAt   time.Time   `json:"publishedAt"`
	Behind        bool        `json:"behind"`
	Content       Content     `json:"content"`
	Subdocs       interface{} `json:"subdocs"`
	Channel       string      `json:"channel"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	CreatedAt     time.Time   `json:"createdAt"`
}

func main() {
	aPostRequest()
}

func aPostRequest() {
	fmt.Print("running a post request \n")

	channel := env.Get("LUCENTV3_CHANNEL")
	secret := env.Get("LUCENTV3_SECRET")
	user := env.Get("LUCENTV3_USER")
	locale := env.Get("LUCENTV3_LOCALE")

	dur := time.Duration(5 * time.Second)

	lc := lucentcmsgo.NewLucentClient(channel, secret, user, locale, dur)

	d := make(map[string]interface{})

	requestContent := make(map[string]interface{})
	requestContent["title"] = "This should not be created long string!"
	requestContent["excerpt"] = "hello from golang a long string !"

	d["schema"] = "articles"
	d["content"] = requestContent

	q := make(map[string]interface{})
	q["data"] = d

	fmt.Printf("request body \n %v\n", d)

	request, err := lc.NewRequest("POST", "documents", nil)
	request.AddData(q)

	if err != nil {
		log.Fatalf("error %v\n", err.Error())
	}

	res, err := request.Post()

	if err != nil {
		log.Fatalf("error %v\n", err.Error())
	}

	fmt.Println(res)
}

func oldMain() {
	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")
	locale := env.Get("LUCENTV3_LOCALE")

	dur := time.Duration(5 * time.Second)

	_ = lucentcmsgo.NewLucentClient(channel, token, user, locale, dur)

	httpClient := http.Client{
		Timeout: dur,
	}

	var data = make(map[string]string, 0)

	data["filter[schema]"] = "articles"
	data["include"] = "*"

	requestData, err := json.Marshal(data)

	request, err := http.NewRequest("GET", "https://api.lucentcms.com/api/documents", bytes.NewBuffer(requestData))

	if err != nil {
		fmt.Printf("error %v", err.Error())
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
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

	var response AResponse

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Printf("error in json response %v", err.Error())
	}

	fmt.Println(response.Data[0].Channel, "response body")
	fmt.Println(response)
}

func aTestRunThatWorked() {
	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")
	locale := env.Get("LUCENTV3_LOCALE")

	dur := time.Duration(5 * time.Second)

	_ = lucentcmsgo.NewLucentClient(channel, token, user, locale, dur)

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
