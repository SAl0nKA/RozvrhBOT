package discord

import (
	"github.com/SAl0nKA/RozvrhBOT/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

//TODO pridať nazov bota do helpu
func GetHelp(s *discordgo.Session,m *discordgo.MessageCreate){
	var EmbedHelp = discordgo.MessageEmbed{
		Title: fmt.Sprintf("Help - Verzia %s",config.Version),
		Color:     17407, //BLUE
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: config.BotPrefix + "help",
				Value:  "Vypíše použitelné príkazy",
			},
			{
				Name: config.BotPrefix + "ping",
				Value: "Odpíše \"Pong!\"",
			},
			{
				Name: config.BotPrefix + "pong",
				Value:  "Odpíše \"Ping!\"",
			},
			{
				Name: config.BotPrefix + "github",
				Value:  "Link na GitHub",
			},
			{
				Name: config.BotPrefix + "hod",
				Value:  "Vypíše najbližšiu hodinu",
			},
			{
				Name: config.BotPrefix + "dalsia",
				Value:  "Vypíše hodinu za najbližšou hodinou",
			},
			{
				Name: config.BotPrefix + "rozvrh",
				Value:  "Vypíše rozvrh na tento deň",
			},
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID,&EmbedHelp)
}

func GetDay(day time.Weekday)string {
	switch day {
	case 1:
		return "Pondelok"
	case 2:
		return "Utorok"
	case 3:
		return "Streda"
	case 4:
		return "Štvrtok"
	case 5:
		return "Piatok"
	default:
		return ""
	}
}

func NemasOpravnenie(s *discordgo.Session, m *discordgo.MessageCreate){
	embed := discordgo.MessageEmbed{
		Title: "Na tento príkaz nemáš opravnenie",
		Color: 16711680,
	}
	s.ChannelMessageSendEmbed(m.ChannelID,&embed)
}
