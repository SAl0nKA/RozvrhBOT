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
	IDstring			string
	config    			configStruct
	SchoolDays			[]*SchoolDay
)

type configStruct struct {
	Token     			string
	BotPrefix 			string
	Dni		  			[][]string
	LinkKuHodine 		map[string]string
	IDs					[]string
	Casy				[]string
	Linky				[]string
	Hodiny				[]int
	Minuty				[]int
	DefaultChannelID 	[]string
	IDstring			string
}

type SchoolDay struct {
	Hodiny	[]string
	Linky	[]string
	Casy 	[]string
	day 	time.Weekday
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
			"#PRIKLAD=FYZ,FYZ,FYZ,FYZ,FYZ,FYZ,FYZ\n" +
			"PONDELOK=\n" +
			"UTOROK=\n" +
			"STREDA=\n" +
			"STVRTOK=\n" +
			"PIATOK=\n" +
			"IDS=\n" +
			"#PRIKLAD=7:50-8:35,8:40-9:25,9:35-10:20,10:40-11:25,11:35-12:20,12:30-13:10,13:20-14:00\n" +
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

	IDstring = os.Getenv("IDS")
	config.IDstring = IDstring
	IDs = strings.Split(os.Getenv("IDS"),",")
	config.IDs = IDs
	if config.IDstring == "" {
		fmt.Println("No IDS; everyone will be able to use the bot")
	}

	DefaultChannelID = strings.Split(os.Getenv("DEFAULT_CHANNEL"),",")
	config.DefaultChannelID = DefaultChannelID
	if config.DefaultChannelID == nil {
		fmt.Println("No DEFAULT_CHANNEL; automatic lesson announcement will not run")
	}
	//####################################################################
	Casy = strings.Split(os.Getenv("CASY"),",")
	config.Casy = Casy
	if config.Casy == nil {
		log.Println("no CASY provided")
		time.Sleep(time.Second*5)
		os.Exit(3)
	}

	pondelok := strings.Split(os.Getenv("PONDELOK"),",")
	utorok := strings.Split(os.Getenv("UTOROK"),",")
	streda := strings.Split(os.Getenv("STREDA"),",")
	stvrtok := strings.Split(os.Getenv("STVRTOK"),",")
	piatok := strings.Split(os.Getenv("PIATOK"),",")
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
	LinkKuHodine, err = godotenv.Read("linky.txt")
	config.LinkKuHodine = LinkKuHodine
	LinkKuHodine["-"]="Momentalne nie je žiadna hodina"

	if err != nil {
		log.Println("Can't open config file or empty variables; creating linky.txt for you to use for your config")
		f, err := os.Create("linky.txt")
		if err != nil {
			log.Println("Issue creating sample linky.txt")
			time.Sleep(time.Second*5)
			os.Exit(3)
			return err
		}
		for _,hod := range hodiny{
			JedinecneHodiny = AppendIfMissing(JedinecneHodiny,hod)
		}
		for _,hodina := range JedinecneHodiny{
			if hodina != "-"{
				f.WriteString(fmt.Sprint(hodina,"=\n"))
			}
		}
		f.Close()
		time.Sleep(time.Second*5)
		os.Exit(3)
	}

	Hodiny, Minuty = SplitCasy(Casy)
	config.Hodiny = Hodiny
	config.Minuty = Minuty


	SchoolDays = append(SchoolDays,NewSchoolDay(nil,nil,nil,0))
	for index, den := range Dni{
		if len(den) > 8{
			log.Printf("There are more than 8 lessons in the config file.")
			time.Sleep(time.Second*5)
			os.Exit(3)
		}
		var (
			hodiny	[]string
			linky   []string
			casy 	[]string
		)
		for _,hodina := range den{
			if hodina == "-"{
				hodiny = append(hodiny,"")
			} else {
				hodiny = append(hodiny,hodina)
			}
		}
		i := 0
		for _,hod := range den{
			for hodina,link := range LinkKuHodine{
				if hod == hodina{
					linky = append(linky, link)
					if hod == "-"{
						casy = append(casy,"")
					} else {
						casy = append(casy,Casy[i])
					}
					i++
					break
				}
			}
		}
	/*	for i := 0; i<len(den);i++{
			casy = append(casy,Casy[i])
		}*/

		SchoolDays = append(SchoolDays,NewSchoolDay(hodiny,linky,casy, time.Weekday(index+1)))
	}
	/*for _,sd := range SchoolDays{
		fmt.Println(sd.Hodiny)
		fmt.Println(sd.Casy)
		fmt.Println(sd.Linky)
	}*/
	fmt.Println("Reading config file successful")
	return nil
}

func SplitCasy(casy []string)([]int, []int){
	var CasySplit []string
	var Hodiny []int
	var Minuty []int
	//casy 7:50-8:35
	for _,cas := range casy{
		//7:50
		CasySplit = append(CasySplit,strings.Split(cas,"-")[0])
		//8:35
		CasySplit = append(CasySplit,strings.Split(cas,"-")[1])
	}
	var Hod int
	var	Min int

	//CasySplit 7:50
	for _, HodMin := range CasySplit{
		if HodMin == "0" {
			Hod = 0
			Min = 0
		} else {
			//7
			Hod,_ = strconv.Atoi(strings.Split(HodMin,":")[0])
			//50
			Min,_ = strconv.Atoi(strings.Split(HodMin,":")[1])
		}
		//Hodiny 7
		Hodiny = append(Hodiny,Hod)
		//Minuty 50
		Minuty = append(Minuty,Min)
	}
	for i:=0;i<(15-len(Hodiny));i++{
		Hodiny = append(Hodiny,0)
		Minuty = append(Minuty,0)
	}
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

func NewSchoolDay(hodiny, linky, casy []string, day time.Weekday)*SchoolDay{
	s := SchoolDay{
		Hodiny: hodiny,
		Linky:  linky,
		Casy:   casy,
		day: 	day,
	}
	return &s
}
