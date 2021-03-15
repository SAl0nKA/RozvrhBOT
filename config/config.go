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
	PingRoleID		  string

	EndMessageEnable  bool
	PingRoleEnable	  bool

	Version           string = "v3.1.0"
)

func ReadConfig() error {
	log.Println("[RozvrhBOT] Innitializing RozvrhBOT")
	err := godotenv.Load("config.txt")
	if err != nil {
		log.Println("[RozvrhBOT] Can't open config.txt; creating file for you to use")
		CreateConfigFile()
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
		log.Println("[RozvrhBOT][WARNING] No ROLES_IDS; everyone will be able to use the bot")
	}

	DefaultChannelsID = strings.Split(os.Getenv("DEFAULT_CHANNELS"),",")
	if DefaultChannelsID[0] == "" {
		log.Println("[RozvrhBOT][WARNING] No DEFAULT_CHANNEL; automatic lesson announcement will not run")
	}

	if os.Getenv("END_MESSAGE_ENABLE") == "true"{
		log.Println("[RozvrhBOT] Enabling sending end messages")
		EndMessageEnable = true

		EndMessage = os.Getenv("END_MESSAGE")
		if EndMessage == ""{
			log.Println("[RozvrhBOT][Warning] No END_MESSAGE; using default message")
			EndMessage = "Konečne je koniec."
		}
	} else {
		log.Println("[RozvrhBOT][Warning] End messages not enabled")
	}

	if os.Getenv("PING_ROLE_ENABLE") == "true"{
		PingRoleID = os.Getenv("PING_ROLE_ID")
		if PingRoleID == ""{
			log.Println("[RozvrhBOT][Warning] No PING_ROLE_ID; role pinging will not be enabled")
		}
		log.Println("[RozvrhBOT] Enabling role pinging")
		PingRoleEnable = true
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

	//vytvorenie linky.txt
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
		if len(den) > len(CasyStrings){
			return errors.New("There are more lessons than set times")
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
<<<<<<< HEAD
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
=======
}

func CreateConfigFile(){
	f, err := os.Create("config.txt")
	defer f.Close()
	if err != nil {
		log.Println("[RozvrhBOT] Issue creating sample config.txt")
		time.Sleep(time.Second*5)
		os.Exit(3)
	}

	_, err = f.WriteString(fmt.Sprintf("#Token ktorým sa bot prihlasuje\nDISCORD_BOT_TOKEN=\n\n" +
		"#Prefix pred príkazy pre bota\nBOT_PREFIX=\n\n" +
		"#Miesto na hodiny v jednotlivé dni, zadavajte vo formáte FYZ,FYZ,FYZ\nPONDELOK=\n" +
		"UTOROK=\n" +
		"STREDA=\n" +
		"STVRTOK=\n" +
		"PIATOK=\n\n" +
		"#ID rolí ktoré majú mať prístup k príkazom\n#Ak nechané prázdne, každy bude mať prístup k príkazo\nROLES_IDS=\n\n" +
		"#Časy v ktorých prebiehajú hodiny, zadajte všetky časy od prvej po poslednú hodinu, príklad:\n" +
		"#CASY=7:50-8:35,8:40-9:25,9:35-10:20,10:40-11:25,11:35-12:20,12:30-13:10,13:20-14:00\n" +
		"CASY=\n\n" +
		"#ID kanálov do ktorých sa automatický budú oznamovať hodiny\nDEFAULT_CHANNELS=\n" +
		"#Povolenie koncových správ\nEND_MESSAGE_ENABLE=true\n\n" +
		"#Vlastná koncová správa, ak nechané prázdné, použije sa prednastavená správa\nEND_MESSAGE=\n\n" +
		"#Povolenie pingovania role v automatických správach\nPING_ROLE_ENABLE=\n\n" +
		"#ID role ktorá sa ma pingnuť\nPING_ROLE_ID="))
	if err != nil {
		log.Println("[RozvrhBOT] Issue writing to config.txt")
		time.Sleep(time.Second*5)
		os.Exit(3)
	}

	time.Sleep(time.Second*5)
	os.Exit(3)
}
>>>>>>> e28bd25 (Pridané komentáre k nastaveniam)
