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
	Hodiny			[]string
	Linky			[]string
	Casy 			[]string
	KoniecVyuc		cas
}

//TODO vytvoriť lepši spôsob ukladania časov
type cas struct {
	Hodina 		int
	Minuta		int
}

var (
	Token             string
	BotPrefix         string
	Dni               [][]string
	HodinaKuLinku     map[string]string
	RoleIDs           []string
	CasyStrings       []string
	DefaultChannelsID []string
	ZaciatokHodin 	  []*cas
	KoniecHodin		  []*cas
	RoleIDSstring     string
	SchoolDays        []*SchoolDay
	EndMessage        string
	Version           string = "v3.1.0"
)

func ReadConfig() error {
	log.Println("[RozvrhBOT] Innitializing RozvrhBOT")
	err := godotenv.Load("config.txt")
	if err != nil {
		log.Println("[RozvrhBOT] Can't open config.txt; creating file for you to use")
		f, err := os.Create("config.txt")
		if err != nil {
			log.Println("[RozvrhBOT] Issue creating sample config.txt")
			time.Sleep(time.Second*5)
			os.Exit(3)
		}
		//TODO pridať komentare k jednotlivym premennym
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

	CasyStrings = strings.Split(os.Getenv("CASY"),",")
	if os.Getenv("CASY") == "" {
		return errors.New("No CASY provided; shutting down")
	}

	BotPrefix = os.Getenv("BOT_PREFIX")
	if BotPrefix == "" {
		log.Println("[RozvrhBOT][WARNING] BOT_PREFIX not configured")
	}

	RoleIDSstring = os.Getenv("ROLES_IDS")

	RoleIDs = strings.Split(os.Getenv("ROLES_IDS"),",")
	if RoleIDSstring == ""{
		log.Println("[RozvrhBOT][WARNING] No ROLES_IDS; everyone will be able to use the bot")
	}

	DefaultChannelsID = strings.Split(os.Getenv("DEFAULT_CHANNELS"),",")
	if DefaultChannelsID[0] == "" {
		log.Println("[RozvrhBOT][WARNING] No DEFAULT_CHANNEL; automatic lesson announcement will not run")
	}

	//TODO pridať možnosť zapnutia koncovych sprav
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
	HodinaKuLinku, err := godotenv.Read("linky.txt")
	HodinaKuLinku["-"]=""

	if err != nil {
		log.Println("[RozvrhBOT] Can't open linky.txt; creating sample file for you to use")
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

	SplitCasy(CasyStrings)

	SchoolDays = append(SchoolDays,NewSchoolDay(nil,nil,nil,cas{Hodina: 0, Minuta: 0,}))
	//Vytvorenie SchoolDay pre každy jeden deň
	for _, den := range Dni{
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
			for hodina,link := range HodinaKuLinku{
				if hod == hodina{
					linky = append(linky, link)
					if hod == "-"{
						casy = append(casy,"")
					} else {
						casy = append(casy,CasyStrings[i])
					}
					i++
					break
				}
			}
		}

		var koniecdna cas
		if len(linky) > 0{
			koniecdna = *KoniecHodin[len(hodiny)-1]
		}

		//doratanie do plnej dlžky
		for i:=0;i<(9-len(hodiny));i++{
			hodiny = AppendIfMissing(hodiny,"")
			linky = AppendIfMissing(linky,"")
			casy = AppendIfMissing(casy,"")
		}
		SchoolDays = append(SchoolDays,NewSchoolDay(hodiny,linky,casy,koniecdna))
	}

	log.Println("[RozvrhBOT] Innitialization succsessful")
	return nil
}

func NewSchoolDay(hodiny, linky, casy []string,koniecvyuc cas)*SchoolDay{
	s := SchoolDay{
		Hodiny: hodiny,
		Linky:  linky,
		Casy:   casy,
		KoniecVyuc: koniecvyuc,
	}
	return &s
}

func SplitCasy(casy []string){
	var casySplit []string

	//cas 7:50-8:35
	for _,cas := range casy{
		//7:50,8:35
		casySplit = append(casySplit,strings.Split(cas,"-")...)
	}

	var hod int
	var	min int

	//casySplit 7:50
	for i := 0; i < len(casySplit);i++{
		hod,min = 0,0
		hod,_ = strconv.Atoi(strings.Split(casySplit[i],":")[0])
		min,_ = strconv.Atoi(strings.Split(casySplit[i],":")[1])

		if i % 2 == 0{
			c := cas{
				Hodina: hod,
				Minuta: min,
			}
			ZaciatokHodin = append(ZaciatokHodin,&c)
		} else {
			c := cas{
				Hodina: hod,
				Minuta: min,
			}
			KoniecHodin = append(KoniecHodin,&c)
		}
	}

	//doratanie do plnej dlžky
	c := cas{
		Hodina: -1,
		Minuta: -1,
	}
	for i:=0;i<(8-len(ZaciatokHodin));i++{
		ZaciatokHodin = append(ZaciatokHodin,&c)
		KoniecHodin = append(KoniecHodin,&c)
	}
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}