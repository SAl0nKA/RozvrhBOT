package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Token     			string
	BotPrefix 			string
	Dni		 			[][]string
	LinkKuHodine 		map[string]string
	IDs					[]string
	Casy				[]string
	DefaultChannelID	[]string
	Hodiny				[]int
	Minuty				[]int
	config    			configStruct
)

type configStruct struct {
	Token     			string
	BotPrefix 			string
	Dni		  			[][]string
	LinkKuHodine 		map[string]string
	IDs					[]string
	Casy				[]string
	Hodiny				[]int
	Minuty				[]int
	DefaultChannelID 	[]string
}

func ReadConfig() error {
	fmt.Println("Reading from config file...")
	err := godotenv.Load("config.txt")
	if err != nil && os.Getenv("DISCORD_BOT_TOKEN") == "" {
		log.Println("Can't open config file and missing variables; creating config.txt for you to use for your config")
		f, err := os.Create("config.txt")
		if err != nil {
			log.Println("Issue creating sample config.txt")
			time.Sleep(time.Second*5)
			os.Exit(3)
		}
		_, err = f.WriteString(fmt.Sprintf("DISCORD_BOT_TOKEN=\n" +
			"BOT_PREFIX=\n" +
			"#PRIKLAD=FYZ FYZ FYZ FYZ FYZ FYZ FYZ\n" +
			"PONDELOK=\n" +
			"UTOROK=\n" +
			"STREDA=\n" +
			"STVRTOK=\n" +
			"PIATOK=\n" +
			"IDS=\n" +
			"#PRIKLAD=7:50-8:35 8:40-9:25 9:35-10:20 10:40-11:25 11:35-12:20 12:30-13:10 13:20-14:00\n" +
			"CASY=\n" +
			"DEFAULT_CHANNEL=\n"))
		f.Close()
		time.Sleep(time.Second*5)
		os.Exit(3)
	}

	Token = os.Getenv("DISCORD_BOT_TOKEN")
	config.Token = Token
	if config.Token == "" {
		return errors.New("no DISCORD_BOT_TOKEN provided")
	}

	BotPrefix = os.Getenv("BOT_PREFIX")
	config.BotPrefix = BotPrefix
	if config.BotPrefix == "" {
		return errors.New("no BOT_PREFIX provided")
	}

	IDs = strings.Split(os.Getenv("IDS")," ")
	config.IDs = IDs
	if config.BotPrefix == "" {
		fmt.Println("No IDS; everyone will be able to use bot")
	}

	Casy = strings.Split(os.Getenv("CASY")," ")
	config.Casy = Casy
	if config.Casy == nil {
		return errors.New("no CASY provided")
	}

	DefaultChannelID = strings.Split(os.Getenv("DEFAULT_CHANNEL")," ")
	config.DefaultChannelID = DefaultChannelID
	if config.DefaultChannelID == nil {
		fmt.Println("No DEFAULT_CHANNEL; automatic lesson announcement will not run")
	}

	pondelok := strings.Split(os.Getenv("PONDELOK")," ")
	utorok := strings.Split(os.Getenv("UTOROK")," ")
	streda := strings.Split(os.Getenv("STREDA")," ")
	stvrtok := strings.Split(os.Getenv("STVRTOK")," ")
	piatok := strings.Split(os.Getenv("PIATOK")," ")
	Dni = [][]string{
		pondelok,
		utorok,
		streda,
		stvrtok,
		piatok,
	}
	config.Dni = Dni
	var hodiny []string
	for _,den := range config.Dni{
		hodiny = append(hodiny, den...)
	}

	var JedinecneHodiny []string
	//var LinkKuHodine map[string]string
	LinkKuHodine, err = godotenv.Read("hodiny.txt")
	config.LinkKuHodine = LinkKuHodine
	if err != nil {
		log.Println("Can't open config file or empty variables; creating hodiny.txt for you to use for your config")
		f, err := os.Create("hodiny.txt")
		if err != nil {
			log.Println("Issue creating sample hodiny.txt")
			time.Sleep(time.Second*5)
			os.Exit(3)
			return err
		}
		for _,hod := range hodiny{
			JedinecneHodiny = AppendIfMissing(JedinecneHodiny,hod)
		}
		for _,hodina := range JedinecneHodiny{
			f.WriteString(fmt.Sprint(hodina,"=\n"))
		}
		f.Close()
		time.Sleep(time.Second*5)
		os.Exit(3)
	}
	Hodiny, Minuty = SplitCasy(Casy)
	config.Hodiny = Hodiny
	config.Minuty = Minuty

	return nil
}

func SplitCasy(casy []string)([]int, []int){
	var CasySplit []string
	var cas string
	var Hodiny []int
	var Minuty []int
	//casy 7:50-8:35
	for _,cas = range casy{
		//7:50
		CasySplit = append(CasySplit,strings.Split(cas,"-")[0])
		//8:35
		CasySplit = append(CasySplit,strings.Split(cas,"-")[1])
	}
	//CasySplit 7:50
	for _, HodMin := range CasySplit{
		//7
		Hod,_ := strconv.Atoi(strings.Split(HodMin,":")[0])
		//50
		Min,_ := strconv.Atoi(strings.Split(HodMin,":")[1])
		//Hodiny 7
		Hodiny = append(Hodiny,Hod)
		//Minuty 50
		Minuty = append(Minuty,Min)
	}

	/*fmt.Println(Hodiny)
	fmt.Println(Minuty)*/
	return Hodiny,Minuty
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

/*
	logPath := os.Getenv("LOG_PATH")
	if logPath == "" {
		logPath = "./"
	}
*/

/*
	logEntry := os.Getenv("DISABLE_LOG_FILE")
	if logEntry == "" {
		file, err := os.Create(path.Join(logPath, "logs.txt"))
		if err != nil {
			return err
		}
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	}
*/