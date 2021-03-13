package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)
//TODO pridať možnosť pinguť potrebnú rolu
func CommandHod(s *discordgo.Session, m *discordgo.MessageCreate){
	if PermissionsCheck(m.Member.Roles){
		hod, link, cas := Hodiny(0)
		if hod == "" {
			s.ChannelMessageSend(m.ChannelID, link)
			return
		} else {
			embed := discordgo.MessageEmbed{
				Title: fmt.Sprintf("Najbližšia hodina je %s o %s",hod,cas),
				Description: link,
				Color: 177013,
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
		if cas == "" {
			s.ChannelMessageSend(m.ChannelID, "Už nie je žiadna hodina")
		} else {
			embed := discordgo.MessageEmbed{
				Title: fmt.Sprintf("Ďalšia hodina je %s o %s",hod,cas),
				Description: link,
				Color: 177013,
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
