package main

import (
	"./config"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"./bot"
)

func main() {
	fmt.Println("Version 3.0.5")
	err := config.ReadConfig()
	if err != nil {
		log.Println(err.Error())
		log.Println("Closing the window in 5 seconds")
		time.Sleep(time.Second*5)
		return
	}
	f, err := os.OpenFile("logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	bot.Start()
}
