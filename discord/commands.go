package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func CommandHod(s *discordgo.Session, m *discordgo.MessageCreate){
	if PermissionsCheck(m.Member.Roles){
		hod, link, cas := Hodiny(0)
		if hod == "" {
			embed := discordgo.MessageEmbed{
				Title: "Momentalne nie je žiadna hodina",
				Color: 15105570, //orange
			}
			s.ChannelMessageSendEmbed(m.ChannelID,&embed)
		} else {
			embed := discordgo.MessageEmbed{
				Title: fmt.Sprintf("Najbližšia hodina je %s o %s",hod,cas),
				Description: link,
				Color: 177013, //green
			}
			s.ChannelMessageSendEmbed(m.ChannelID,&embed)
		}
	} else {
		NemasOpravnenie(s,m)
	}
}

func CommandDalsia(s *discordgo.Session, m *discordgo.MessageCreate){
	hod, link, cas := Hodiny(1)
	if PermissionsCheck(m.Member.Roles){
		if hod == "" {
			embed := discordgo.MessageEmbed{
				Title: "Už nenasleduje žiadna hodina",
				Color: 15105570, //orange
			}
			s.ChannelMessageSendEmbed(m.ChannelID,&embed)
		} else {
			embed := discordgo.MessageEmbed{
				Title: fmt.Sprintf("Ďalšia hodina je %s o %s",hod,cas),
				Description: link,
				Color: 177013,//green
			}
			s.ChannelMessageSendEmbed(m.ChannelID,&embed)
		}
	} else {
		NemasOpravnenie(s,m)
	}
}

func CommandRozvrh(s *discordgo.Session, m *discordgo.MessageCreate){
	if PermissionsCheck(m.Member.Roles){
		embed := ReturnRozvrh(0,"")
		mes,_ := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		AddReactions(s,mes.ChannelID, mes.ID)
		rozvrh := NewRozvrh(mes.ChannelID, mes.ID, mes.GuildID,time.Now().Weekday())
		RozvrhEmbedy = append(RozvrhEmbedy,rozvrh)
	} else {
		NemasOpravnenie(s,m)
	}
}
