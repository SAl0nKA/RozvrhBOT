package discord

import (
	"../config"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

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
				Name: config.BotPrefix + "hod",
				Value:  "Vypíše najbližšiu hodinu",
			},
			{
				Name: config.BotPrefix + "dalsia",
				Value:  "Vypíše ďalšiu hodinu",
			},
			{
				Name: config.BotPrefix + "rozvrh",
				Value:  "Vypíše rozvrh na tento deň",
			},
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &EmbedHelp)
}




