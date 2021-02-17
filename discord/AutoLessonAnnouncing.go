package discord

import (
	"../config"
	//"../discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func HodAnnounce(s *discordgo.Session) {
	hodiny := config.Hodiny
	minuty := config.Minuty
	for {
		t := time.Now()
		h := t.Hour()
		m := t.Minute()
		//TODO prerobiť tento retardovany switch
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
			for _,channelID := range config.DefaultChannelsID{
				JeKoniec := discordgo.MessageEmbed{
					Title: config.EndMessage,
					Description: "Beep Boop. Táto správa je automatizovaná",
					Color:     16711680, //RED
				}
				s.ChannelMessageSendEmbed(channelID, &JeKoniec)
			}
		}
		if h >= hodiny[len(config.Casy)*2-1] {
			log.Printf("Turning off the automatic lesson announcing")
			break
		}
		time.Sleep(time.Second * 35)
	}
}

func HodAnnounceHelp(s *discordgo.Session, BaseHod int) {
	log.Println("[RozvrhBOT] Announcing lesson")
	sd := GetSChoolday(0)
	if len(sd.Hodiny)-1<BaseHod{
		return
	}
	hod := sd.Hodiny[BaseHod]
	cas := sd.Casy[BaseHod]
	link :=  sd.Linky[BaseHod]
	if hod == "" {
		for _,channelID := range config.DefaultChannelsID{
			embed := discordgo.MessageEmbed{
				Title: link,
				Description: "*Beep Boop. Táto správa je automatizovaná*",
				Color: 16711680,//red
			}
			s.ChannelMessageSendEmbed(channelID,&embed)
		}
		time.Sleep(time.Minute)
	} else {
		for _,channelID := range config.DefaultChannelsID{
			embed := discordgo.MessageEmbed{
				Title: fmt.Sprintf("Najbližšia hodina je %s o %s",hod,cas),
				Description: fmt.Sprintf("%s\n*Beep Boop. Táto správa je automatizovaná*",link),
				Color: 177013,//green
			}
			s.ChannelMessageSendEmbed(channelID,&embed)
		}
		time.Sleep(time.Minute)
	}
}