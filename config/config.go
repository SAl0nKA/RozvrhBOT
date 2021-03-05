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

type SchoolDay struct {
	Hodiny	[]string
	Linky	[]string
	Casy 	[]string
	day 	time.Weekday
}
//TODO vytvoriť lepši spôsob ukladania časov
type cas struct {
	hodina 		int
	minuta		int
}

var (
	Token     			string
	BotPrefix 			string
	Dni		 			[][]string
	LinkKuHodine 		map[string]string
	RoleIDs				[]string
	Casy				[]string
	DefaultChannelsID	[]string
	Hodiny				[]int
	Minuty				[]int
	RoleIDSstring		string
	SchoolDays			[]*SchoolDay
	EndMessage			string
	Version 			string = "v3.1.0"
)

func ReadConfig() error {
	log.Println("[RozvrhBOT] Innitializing RozvrhBOT")
	err := godotenv.Load("config.txt")
	if err != nil {
		log.Println("[RozvrhBOT] Can't open config file; creating config.txt for you to use for your config")
		f, err := os.Create("config.txt")
		if err != nil {
			log.Println("[RozvrhBOT] Issue creating sample config.txt")
			time.Sleep(time.Second*5)
			os.Exit(3)
		}

		_, err = f.WriteString(fmt.Sprintf("DISCORD_BOT_TOKEN=\n" +
			"BOT_PREFIX=\n" +
			"PONDELOK=\n" +
			"UTOROK=\n" +
			"STREDA=\n" +
			"STVRTOK=\n" +
			"PIATOK=\n" +
			"ROLES_IDS=\n" +
			"#PRIKLAD=7:50-8:35,8:40-9:25,9:35-10:20,10:40-11:25,11:35-12:20,12:30-13:10,13:20-14:00\n" +
			"CASY=\n" +
			"DEFAULT_CHANNELS=\n" +
			"END_MESSAGE="))
		f.Close()
		time.Sleep(time.Second*5)
		os.Exit(3)
	}

	Token = os.Getenv("DISCORD_BOT_TOKEN")
	//config.Token = Token
	if Token == "" {
		return errors.New("No DISCORD_BOT_TOKEN provided; shutting down")
	}

	Casy = strings.Split(os.Getenv("CASY"),",")
	if os.Getenv("CASY") == "" {
		return errors.New("No CASY provided; shutting down")
	}

	BotPrefix = os.Getenv("BOT_PREFIX")
<<<<<<< HEAD
<<<<<<< HEAD
	config.BotPrefix = BotPrefix
	/*if config.BotPrefix == "" {
		return errors.New("no BOT_PREFIX provided")
	}*/
=======
	//config.BotPrefix = BotPrefix
>>>>>>> 8754588 (Save)
=======
	if BotPrefix == "" {
		log.Println("[RozvrhBOT][WARNING] BOT_PREFIX not configured")
	}
>>>>>>> a833281 (Upratanie kodu)

	RoleIDSstring = os.Getenv("ROLES_IDS")

	RoleIDs = strings.Split(os.Getenv("ROLES_IDS"),",")
	if RoleIDSstring == ""{
		log.Println("[RozvrhBOT][WARNING] No IDS; everyone will be able to use the bot")
	}

	DefaultChannelsID = strings.Split(os.Getenv("DEFAULT_CHANNELS"),",")
	if DefaultChannelsID == nil {
		log.Println("[RozvrhBOT][WARNING] No DEFAULT_CHANNEL; automatic lesson announcement will not run")
	}

	EndMessage = os.Getenv("END_MESSAGE")
	if EndMessage == ""{
		EndMessage = "Konečne je koniec."
	}

	//hodiny pre jednotlive dni
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

	var hodiny []string
	for _,den := range Dni{
		hodiny = append(hodiny, den...)
	}
	var JedinecneHodiny []string
	LinkKuHodine, err := godotenv.Read("linky.txt")
	LinkKuHodine["-"]="Momentalne nie je žiadna hodina"

	if err != nil {
		log.Println("[RozvrhBOT] Can't open config file; creating linky.txt for you to use for your config")
		f, err := os.Create("linky.txt")
		defer f.Close()
		if err != nil {
			return errors.New("Issue creating sample linky.txt")
		}
		for _,hod := range hodiny{
			JedinecneHodiny = AppendIfMissing(JedinecneHodiny,hod)
		}
		for _,hodina := range JedinecneHodiny{
			if hodina != "-"{
				f.WriteString(fmt.Sprint(hodina,"=\n"))
			}
		}
		time.Sleep(time.Second*5)
		os.Exit(3)
	}

	Hodiny, Minuty = SplitCasy(Casy)

	SchoolDays = append(SchoolDays,NewSchoolDay(nil,nil,nil,0))
	for index, den := range Dni{
		if len(den) > 8{
			return errors.New("There are more than 8 lessons in the config file")
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
		SchoolDays = append(SchoolDays,NewSchoolDay(hodiny,linky,casy, time.Weekday(index+1)))
	}
<<<<<<< HEAD
	/*for _,sd := range SchoolDays{
		fmt.Println(sd.Hodiny)
		fmt.Println(sd.Casy)
		fmt.Println(sd.Linky)
	}*/
=======

	log.Println("[RozvrhBOT] Innitialization succsessful")
>>>>>>> a833281 (Upratanie kodu)
	return nil
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

func SplitCasy(casy []string)([]int, []int){
	var (
		casySplit []string
		hodiny []int
		minuty []int
	)

	//casy 7:50-8:35
	for _,cas := range casy{
		//7:50
		casySplit = append(casySplit,strings.Split(cas,"-")[0])
		//8:35
		casySplit = append(casySplit,strings.Split(cas,"-")[1])
	}
	var hod int
	var	min int

	//CasySplit 7:50
	for _, HodMin := range casySplit{
		if HodMin == "0" {
			hod = 0
			min = 0
		} else {
			//7
			hod,_ = strconv.Atoi(strings.Split(HodMin,":")[0])
			//50
			min,_ = strconv.Atoi(strings.Split(HodMin,":")[1])
		}
		//Hodiny 7
		hodiny = append(hodiny,hod)
		//Minuty 50
		minuty = append(minuty,min)
	}
	for i:=0;i<(15-len(Hodiny));i++{
		hodiny = append(hodiny,0)
		minuty = append(minuty,0)
	}
	return hodiny,minuty
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
<<<<<<< HEAD
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

=======
}
>>>>>>> a833281 (Upratanie kodu)
