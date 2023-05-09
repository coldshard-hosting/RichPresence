package main

import (
	"github.com/hugolgst/rich-go/client"
	"time"
)

func main() {
	err := client.Login("1013542996572110898")

	if err != nil {
		panic(err)
	}

	err = client.SetActivity(client.Activity{
		Details:    "Hosting on a budget!",
		LargeImage: "coldshard",
		Buttons: []*client.Button{
			{
				Label: "Visit our Store!",
				Url:   "https://coldshard.com/billing",
			},
			{
				Label: "Join our Discord!",
				Url:   "https://discord.com/f6yGPAxsBE",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	for true {
		time.Sleep(time.Minute * 10)
	}
}