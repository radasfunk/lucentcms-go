package main

import (
	"fmt"
	"time"

	"github.com/radasfunk/lucentcmsgo"
	"github.com/radasfunk/lucentcmsgo/dev/utils/env"
)

func init() {
	env.LoadEnv()
}

func main() {
	channel := env.Get("LUCENTV3_CHANNEL")
	token := env.Get("LUCENTV3_TOKEN")
	user := env.Get("LUCENTV3_USER")

	dur := time.Duration(5 * time.Second)

	client := lucentcmsgo.NewLucentClient(channel, token, user, dur)

	fmt.Printf("got client %v\n", client)
}
