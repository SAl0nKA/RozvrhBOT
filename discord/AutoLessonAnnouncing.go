package discord

import (
	"github.com/SAl0nKA/RozvrhBOT/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func HodAnnounce(s *discordgo.Session) {
	zaciatky := config.ZaciatokHodin
	for {
		t := time.Now()
		w := t.Weekday()
		if w == 0 || w == 6{
			time.Sleep(time.Hour*2)
			return
		}
		h := t.Hour()
		m := t.Minute()

		switch {
		case zaciatky[0].Hodina == h && (zaciatky[0].Minuta-5) == m:
			HodAnnounceHelp(s, 0)
		case zaciatky[1].Hodina == h && (zaciatky[1].Minuta-5) == m:
			HodAnnounceHelp(s, 1)
		case zaciatky[2].Hodina == h && (zaciatky[2].Minuta-5) == m:
			HodAnnounceHelp(s, 2)
		case zaciatky[3].Hodina == h && (zaciatky[3].Minuta-5) == m:
			HodAnnounceHelp(s, 3)
		case zaciatky[4].Hodina == h && (zaciatky[4].Minuta-5) == m:
			HodAnnounceHelp(s, 4)
		case zaciatky[5].Hodina == h && (zaciatky[5].Minuta-5) == m:
			HodAnnounceHelp(s, 5)
		case zaciatky[6].Hodina == h && (zaciatky[6].Minuta-5) == m:
			HodAnnounceHelp(s, 6)
		case zaciatky[7].Hodina == h && (zaciatky[7].Minuta-5) == m:
			HodAnnounceHelp(s, 7)
		}

		KoniecDna := config.SchoolDays[int(w)-1].KoniecVyuc
		if KoniecDna.Hodina == h && KoniecDna.Hodina == m && config.EndMessageEnable{
			log.Println("[RozvrhBOT] Sending end message")
			for _,channelID := range config.DefaultChannelsID{
				JeKoniec := discordgo.MessageEmbed{
					Title: config.EndMessage,
					Description: "",
					Footer: &discordgo.MessageEmbedFooter{
						Text:         "*Beep, Boop. Táto správa je automatizovaná*",
						IconURL:      s.State.User.AvatarURL("128"),
					},
					Color:     16711680, //RED
				}
				s.ChannelMessageSendEmbed(channelID, &JeKoniec)
			}
			time.Sleep(time.Hour)
		}

		time.Sleep(time.Second * 20)
	}
}

func HodAnnounceHelp(s *discordgo.Session, BaseHod int) {
	sd := GetSChoolday(0)
	hod := sd.Hodiny[BaseHod]
	cas := sd.Casy[BaseHod]
	link :=  sd.Linky[BaseHod]
	if hod == "" {
		return
	} else {
		log.Println("[RozvrhBOT] Announcing lesson")
		var ping string
		if config.PingRoleEnable{
			ping = fmt.Sprintf("\n<@&%s>",config.PingRoleID)
		}
		for _,channelID := range config.DefaultChannelsID{
			embed := discordgo.MessageEmbed{
				Title: fmt.Sprintf("Najbližšia hodina je %s o %s",hod,cas),
				Description: fmt.Sprintf("%s%s",link,ping),
				Footer: &discordgo.MessageEmbedFooter{
					Text:         "*Beep, Boop. Táto správa je automatizovaná*",
					IconURL:      s.State.User.AvatarURL("128"),
				},
				Color: 177013,//green
			}
			s.ChannelMessageSendEmbed(channelID,&embed)
		}
		time.Sleep(time.Minute)
	}
}