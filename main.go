package main

import (
	"./config"
	"log"
	"time"
	"./bot"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Println(err.Error())
		log.Println("Closing the window in 5 seconds")
		time.Sleep(time.Second*5)
		return
	}

	bot.Start()
}
