package main

import (
	"fmt"
	"github.com/SAl0nKA/RozvrhBOT/bot"
	"github.com/SAl0nKA/RozvrhBOT/config"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Version",config.Version)
	err := config.ReadConfig()
	if err != nil {
		log.Println(err.Error())
		log.Println("Closing the window in 5 seconds")
		time.Sleep(time.Second*5)
		return
	}
	f, err := os.OpenFile("logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening logs file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	bot.Start()
}
