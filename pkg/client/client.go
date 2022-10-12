package client

import (
	"os"

	wolfram "github.com/krognol/go-wolfram"
	witai "github.com/wit-ai/wit-go/v2"
)

var Client *witai.Client
var WolframClient *wolfram.Client

func ClientInit() {
	Client = witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	WolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}
}