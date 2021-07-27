// testing ground
package main

import (
	"fmt"
	"log"
	"os"
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
	anUploadRequest()
}

func anUploadRequest() {
	fmt.Print("running an upload request \n")

	channel := env.Get("LUCENTV3_CHANNEL")
	secret := env.Get("LUCENTV3_SECRET")
	user := env.Get("LUCENTV3_USER")
	locale := env.Get("LUCENTV3_LOCALE")

	dur := time.Duration(5 * time.Second)

	lc := lucentcmsgo.NewLucentClient(channel, secret, user, locale, dur)

	request, err := lc.NewRequest("files", nil)

	if err != nil {
		log.Fatalf("error %v\n", err.Error())
	}
	c, _ := os.Getwd()

	res, err := request.UploadFromDisk("a_pikachu_file.png", c+"/dev/pikachu.png")

	if err != nil {
		log.Fatalf("error %v\n", err.Error())
	}

	fmt.Println(res.Errors)
}

func aGetRequest() {
	fmt.Print("running a get request \n")

	channel := env.Get("LUCENTV3_CHANNEL")
	secret := env.Get("LUCENTV3_SECRET")
	user := env.Get("LUCENTV3_USER")
	locale := env.Get("LUCENTV3_LOCALE")

	dur := time.Duration(5 * time.Second)

	lc := lucentcmsgo.NewLucentClient(channel, secret, user, locale, dur)

	request, err := lc.NewRequest("documents", nil)

	request.SetLimit(10)

	if err != nil {
		log.Fatalf("error %v\n", err.Error())
	}

	res, err := request.Get()

	if err != nil {
		log.Fatalf("error %v\n", err.Error())
	}

	fmt.Println(res.Errors)
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

	request, err := lc.NewRequest("documents", nil)
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
