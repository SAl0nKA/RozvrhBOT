package bot

import (
	"../config"
	"../discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type CommandType int

type RozvrhEmbed struct {
	ChannelID 		string
	MessageID		string
	GuildID			string
	EditDay 		time.Weekday
}

const (
	Help CommandType = iota
	Ping
	Pong
	Hod
	Dalsia
	Rozvrh
	Null
)

var CommandTypeStringMapping = map[string]CommandType{
	"help":    Help,
	"ping":    Ping,
	"pong":    Pong,
	"hod":     Hod,
	"dalsia":  Dalsia,
	"rozvrh":  Rozvrh,
	"":        Null,
}
var RozvrhEmbedy []*RozvrhEmbed
var goBot *discordgo.Session
var Emojis = []string{"◀️", "▶️", "🔄", "❌"}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	goBot.AddHandler(ready)
	goBot.AddHandler(HandleCommand)
	//goBot.AddHandler(messageCreate)
	goBot.AddHandler(HandleReaction)

	goBot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuildMessageReactions)

	err = goBot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}


	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	goBot.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("Updating status")
	s.UpdateStatus(0, fmt.Sprintf("%shelp",config.BotPrefix))
	t := time.Now().Weekday()
	if config.DefaultChannelID != nil {
		log.Println("Checking for current day")
		if t <= 5 && t != 0 {
			log.Println("Running HodAnnounce function in a separate proccess")
			go HodAnnounce(s)
		} else {
			log.Println("Not runnning HodAnnounce function")
		}
	} else {
		log.Println("Not runnning HodAnnounce function")
	}
}

/*
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID{
		return
	}
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Println("toto je v piči: ", err)
	}
	log.Printf("User %s wrote \"%s\" in channel %s", m.Author, m.Content, channel.Name)
}*/

func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		args := ContentSplit(m.Content)
		switch GetCommandType(args[0]) {
		case Help:
			log.Printf("Reacting to command \"%shelp\"", config.BotPrefix)
			embed := discord.GetHelp()
			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		case Ping:
			log.Printf("Reacting to command \"%sping\"", config.BotPrefix)
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		case Pong:
			log.Printf("Reacting to command \"%spong\"", config.BotPrefix)
			if ContainsIDs(m.Member.Roles, config.IDs) || config.IDstring == ""{
				s.ChannelMessageSend(m.ChannelID, "Ping!")
			} else {
				NemasOpravnenie(s,m)
			}
		case Hod:
			log.Printf("Reacting to command \"%shod\"", config.BotPrefix)
			if ContainsIDs(m.Member.Roles, config.IDs) || config.IDstring == ""{
				hod, link, cas := Hodiny(0)
				if hod == "" {
					s.ChannelMessageSend(m.ChannelID, link)
					return
				} else {
<<<<<<< HEAD
					s.ChannelMessageSend(m.ChannelID, "Najbližšia hodina je "+hod+" o: "+cas+" a link na ňu je: "+link)
					s.ChannelMessageSend(m.ChannelID, "Ďakujeme že využívate nás a nie nejakého relasBOTa")
=======
					embed := discordgo.MessageEmbed{
						Title: fmt.Sprintf("Najbližšia hodina je %s o: %s a link na ňu je:",hod,cas),
						Description: link,
						Color: 177013,
					}
					s.ChannelMessageSendEmbed(m.ChannelID,&embed)
>>>>>>> b43d75b (Pridane embedy)
				}
			} else {
				NemasOpravnenie(s,m)
			}
		case Dalsia:
			log.Printf("Reacting to command \"%sdalsia\"", config.BotPrefix)
			hod, link, cas := Hodiny(1)
			if ContainsIDs(m.Member.Roles, config.IDs) || config.IDstring == ""{
				if cas == "" {
					s.ChannelMessageSend(m.ChannelID, "Už nie je žiadna hodina")
				} else {
<<<<<<< HEAD
					s.ChannelMessageSend(m.ChannelID, "Ďalšia hodina je "+hod+" o: "+cas+" a link na ňu je: "+link)
					s.ChannelMessageSend(m.ChannelID, "Ďakujeme že využívate nás a nie nejakého relasBOTa")
=======
					embed := discordgo.MessageEmbed{
						Title: fmt.Sprintf("Ďalšia hodina je %s o: %s a link na ňu je:",hod,cas),
						Description: link,
						Color: 177013,
					}
					s.ChannelMessageSendEmbed(m.ChannelID,&embed)
					//s.ChannelMessageSend(m.ChannelID, "Ďalšia hodina je "+hod+" o: "+cas+" a link na ňu je: "+link)
>>>>>>> b43d75b (Pridane embedy)
				}
			} else {
				NemasOpravnenie(s,m)
			}
		case Rozvrh:
			log.Printf("Reacting to command \"%srozvrh\"", config.BotPrefix)
			if ContainsIDs(m.Member.Roles, config.IDs) || config.IDstring == ""{
				embed := ReturnRozvrh(0,"")
				mes,_ := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
				//time.Sleep(time.Second*3)
				AddReactions(s,mes.ChannelID, mes.ID)
				rozvrh := NewRozvrh(mes.ChannelID, mes.ID, mes.GuildID,time.Now().Weekday())
				RozvrhEmbedy = append(RozvrhEmbedy,rozvrh)
			} else {
				NemasOpravnenie(s,m)
			}
		}
	}
}

func ContentSplit(sprava string) []string {
	obsah := strings.Fields(sprava)
	return obsah
}

func GetCommandType(arg string) CommandType {
	for str, cmd := range CommandTypeStringMapping {
		if strings.ToLower(arg) == (config.BotPrefix + str) {
			return cmd
		}
	}
	return Null
}

func NewRozvrh(ChannelID,MessageID,GuildID string, EditDay time.Weekday) *RozvrhEmbed {
	r := RozvrhEmbed{
		ChannelID:      ChannelID,
		MessageID:      MessageID,
		GuildID:		GuildID,
		EditDay: 		EditDay,
	}
	return &r
}

func GetSChoolday(day time.Weekday) *config.SchoolDay {
	if day == 0 {
		day = time.Now().Weekday()
	}
	switch day {
	case 1:
		sd := config.SchoolDays[1]
		return sd
	case 2:
		sd := config.SchoolDays[2]
		return sd
	case 3:
		sd := config.SchoolDays[3]
		return sd
	case 4:
		sd := config.SchoolDays[4]
		return sd
	case 5:
		sd := config.SchoolDays[5]
		return sd
	default:
		sd := config.SchoolDays[0]
		return sd
	}
}

func Hodiny(dalsia int) (string, string, string) {
	hodiny := config.Hodiny
	minuty := config.Minuty
	t := time.Now()
	h := t.Hour()
	m := t.Minute()
	if t.Weekday() > 0 && t.Weekday() < 6{
		switch {
		case h <= hodiny[1] ||(h==hodiny[1] && m < minuty[1]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[0+dalsia]
			cas := sd.Casy[0+dalsia]
			link :=  sd.Linky[0+dalsia]
			return hod, link, cas
		case (h == hodiny[2] && m >= minuty[1]) || (h == hodiny[3] && m < minuty[3]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[1+dalsia]
			cas := sd.Casy[1+dalsia]
			link :=  sd.Linky[1+dalsia]
			return hod, link, cas
		case (h == hodiny[4] && m >= minuty[3]) || (h == hodiny[5] && m < minuty[5]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[2+dalsia]
			cas := sd.Casy[2+dalsia]
			link :=  sd.Linky[2+dalsia]
			return hod, link, cas
		case (h == hodiny[6] && m >= minuty[5]) || (h == hodiny[7] && m < minuty[7]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[3+dalsia]
			cas := sd.Casy[3+dalsia]
			link :=  sd.Linky[3+dalsia]
			return hod, link, cas
		case (h == hodiny[8] && m >= minuty[7]) || (h == hodiny[9] && m < minuty[9]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[4+dalsia]
			cas := sd.Casy[4+dalsia]
			link :=  sd.Linky[4+dalsia]
			return hod, link, cas
		case (h == hodiny[10] && m >= minuty[9]) || (h == hodiny[11] && m < minuty[11]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[5+dalsia]
			cas := sd.Casy[5+dalsia]
			link :=  sd.Linky[5+dalsia]
			return hod, link, cas
		case h == hodiny[12] && m >= minuty[11]:
			sd := GetSChoolday(0)
			hod := sd.Hodiny[6+dalsia]
			cas := sd.Casy[6+dalsia]
			link :=  sd.Linky[6+dalsia]
			return hod, link, cas
		default:
			link := "Momentalne nie je žiadna hodina"
			hod := ""
			cas := ""
			return hod, link, cas
		}
	} else {
		link := "Momentalne nie je žiadna hodina"
		hod := ""
		cas := ""
		return hod, link, cas
	}
}

func ReturnRozvrh(day time.Weekday, MessageID string) discordgo.MessageEmbed {
	Rozvrh := NewRozvrh("","","",0)
	if MessageID != "" {
		for _,rozvrh := range RozvrhEmbedy{
			if rozvrh.MessageID == MessageID{
				Rozvrh = rozvrh
				Rozvrh.EditDay = Rozvrh.EditDay + day
				if Rozvrh.EditDay < 1{
					Rozvrh.EditDay = 5
				} else if Rozvrh.EditDay > 5{
					Rozvrh.EditDay = 1
				}
			}
		}
	}
	sd := GetSChoolday(Rozvrh.EditDay)
	return discord.ReturnEmbedRozvrh(sd.Hodiny, sd.Casy, sd.Linky,Rozvrh.EditDay)
}

func HodAnnounce(s *discordgo.Session) {
	hodiny := config.Hodiny
	minuty := config.Minuty
	for {
		time.Sleep(time.Second * 35)
		t := time.Now()
		h := t.Hour()
		m := t.Minute()
		switch {
		case h == hodiny[0] && m == (minuty[0]-5):
			HodAnnounceHelp(s, 0)
		case h == hodiny[2] && m == (minuty[2]-5):
			HodAnnounceHelp(s, 1)
		case h == hodiny[4] && m == (minuty[4]-5):
			HodAnnounceHelp(s, 2)
		case h == hodiny[6] && m == (minuty[6]-5):
			HodAnnounceHelp(s, 3)
		case h == hodiny[8] && m == (minuty[8]-5):
			HodAnnounceHelp(s, 4)
		case h == hodiny[10] && m == (minuty[10]-5):
			HodAnnounceHelp(s, 5)
		case h == hodiny[12] && m == (minuty[12]-5):
			HodAnnounceHelp(s, 6)
		case h == hodiny[14] && m == (minuty[14]-5):
			HodAnnounceHelp(s, 7)
		case h == hodiny[len(config.Casy)*2-1]:
			for _,channelID := range config.DefaultChannelID{
				s.ChannelMessageSendEmbed(channelID, &discord.JeKoniec)
			}
		}
		if h >= hodiny[len(config.Casy)*2-1] {
			log.Printf("Turning off the automatic lesson announcing")
			break
		}
	}
}

func HodAnnounceHelp(s *discordgo.Session, BaseHod int) {
	sd := GetSChoolday(0)
	hod := sd.Hodiny[BaseHod]
	cas := sd.Casy[BaseHod]
	link :=  sd.Linky[BaseHod]
	if hod == "" {
		for _,channelID := range config.DefaultChannelID{
			embed := discordgo.MessageEmbed{
				Title: link,
				Description: "*Beep Boop. Táto správa je automatizovaná*",
				Color: 16711680,//red
			}
			s.ChannelMessageSendEmbed(channelID,&embed)
			/*s.ChannelMessageSend(channelID, link)
			s.ChannelMessageSend(channelID, "*Beep Boop. Táto správa je automatizovaná*")*/
		}
		time.Sleep(time.Minute)
	} else {
		for _,channelID := range config.DefaultChannelID{
			embed := discordgo.MessageEmbed{
				Title: fmt.Sprintf("Najbližšia hodina je %s o: %s",hod,cas),
				Description: fmt.Sprintf("%s\n*Beep Boop. Táto správa je automatizovaná*",link),
				Color: 177013,//green
			}
			s.ChannelMessageSendEmbed(channelID,&embed)
			/*s.ChannelMessageSend(channelID, "Najbližšia hodina je "+hod+" o: "+cas+" a link na ňu je: "+link)
			s.ChannelMessageSend(channelID, "*Beep Boop. Táto správa je automatizovaná*")*/
		}
		time.Sleep(time.Minute)
	}
}

func NemasOpravnenie(s *discordgo.Session, m *discordgo.MessageCreate){
	embed := discordgo.MessageEmbed{
		Title: "Na tento príkaz nemáš opravnenie",
		Color: 16711680,
	}
	s.ChannelMessageSendEmbed(m.ChannelID,&embed)
}

func HandleReaction(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}
	var Rozvrh *RozvrhEmbed
	member,_ := s.GuildMember(r.GuildID,r.UserID)
	for _,rozvrh := range RozvrhEmbedy{
		if rozvrh.MessageID == r.MessageID{
			Rozvrh = rozvrh
		}
	}
	if Rozvrh == nil {
		return
	}
	if Rozvrh.MessageID == r.MessageID && Rozvrh.ChannelID == r.ChannelID && (ContainsIDs(member.Roles, config.IDs) || config.IDs == nil){
		channel, err := s.Channel(r.ChannelID)
		if err != nil {
			log.Println("toto je v piči: ", err)
		}
		switch r.Emoji.Name {
		case "◀️":
			log.Printf("User %s reacted with \"Previous day\" to rozvrh in channel %s", member.Nick, channel.Name)
			s.MessageReactionRemove(r.ChannelID,r.MessageID, "◀️",r.UserID)
			embed := ReturnRozvrh(-1,r.MessageID)
			s.ChannelMessageEditEmbed(r.ChannelID,r.MessageID,&embed)
		case "▶️":
			log.Printf("User %s reacted with \"Next day\" to rozvrh in channel %s", member.Nick, channel.Name)
			s.MessageReactionRemove(r.ChannelID,r.MessageID, "▶️",r.UserID)
			embed := ReturnRozvrh(1,r.MessageID)
			s.ChannelMessageEditEmbed(r.ChannelID,r.MessageID,&embed)
		case "🔄":
			//s.MessageReactionRemove(RozvrhChannelID,RozvrhMessageID, "🔄",r.UserID)
			log.Printf("User %s reacted with \"Refresh\" to rozvrh in channel %s", member.Nick, channel.Name)
			s.ChannelMessageDelete(r.ChannelID,r.MessageID)
			embed := ReturnRozvrh(0,"")
			m, _ := s.ChannelMessageSendEmbed(r.ChannelID,&embed)
			AddReactions(s,m.ChannelID, m.ID)
			rozvrh := NewRozvrh(m.ChannelID, m.ID, m.GuildID,time.Now().Weekday())
			RozvrhEmbedy = append(RozvrhEmbedy,rozvrh)
		case "❌":
			//s.MessageReactionRemove(r.ChannelID,r.MessageID, "❌",r.UserID)
			log.Printf("User %s reacted with \"Delete\" to rozvrh in channel %s", member.Nick, channel.Name)
			s.ChannelMessageDelete(r.ChannelID,r.MessageID)
		}
	} else {
		s.MessageReactionRemove(r.ChannelID,r.MessageID, r.Emoji.Name,r.UserID)
	}
}

func AddReactions(s *discordgo.Session, channelid, messageid string){
	for _,emoji := range Emojis{
		err := s.MessageReactionAdd(channelid,messageid,emoji)
		if err != nil {
			log.Println(err)
		}
	}
}

func ContainsIDs(roles []string, ids []string) bool {
	if config.IDs == nil {
		return true
	}
	for _, role := range roles {
		for _,id := range ids{
			if id == role {
				return true
			}
		}
	}
	return false
}