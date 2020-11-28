package discord

import (
	"../config"
	"github.com/bwmarrin/discordgo"
	"time"
)

func GetHelp()discordgo.MessageEmbed{
	var EmbedHelp = discordgo.MessageEmbed{
		URL:   "",
		Type:  "",
		Title: "Help",
		Description: "",
		Timestamp: "",
		Color:     17407, //BLUE
		Image:    nil,
		Thumbnail: nil,
		Video:     nil,
		Provider:  nil,
		Author:    nil,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: config.BotPrefix + "help",
				Value:  "Vypíše použitelné príkazy",
				Inline: false,
			},
			{
				Name: config.BotPrefix + "ping",
				Value: "Odpíše \"Pong!\"",
				Inline: false,
			},
			{
				Name: config.BotPrefix + "pong",
				Value:  "Odpíše \"Ping!\"",
				Inline: false,
			},
			{
				Name: config.BotPrefix + "hod",
				Value:  "Vypíše najbližšiu hodinu",
				Inline: false,
			},
			{
				Name: config.BotPrefix + "dalsia",
				Value:  "Vypíše ďalšiu hodinu",
				Inline: false,
			},
			{
				Name: config.BotPrefix + "rozvrh",
				Value:  "Vypíše rozvrh na tento deň",
				Inline: false,
			},
		},
	}
	return EmbedHelp
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

func ReturnEmbedRozvrh(rozvrh, casy, linky []string, day time.Weekday)discordgo.MessageEmbed{
	var EmbedRozvrhPrazdny = discordgo.MessageEmbed{
		URL:   "",
		Type:  "",
		Title: "Rozvrh",
		Description: "Dnes nie sú žiadne hodiny, jeb na to",
		Timestamp: "",
		Color:     16711680, //RED
		Image:    nil,
		Thumbnail: nil,
		Video:     nil,
		Provider:  nil,
		Author:    nil,
		Fields: nil,
	}
	if day == 0{
		day = time.Now().Weekday()
	}

	if day <= 5 && day != 0 {
		var EmbedRozvrh = discordgo.MessageEmbed{
			URL:   "",
			Type:  "",
			Title: "Rozvrh - " + GetDay(day),
			Description: "",
			Timestamp: "",
			Color:     16766976, //YELLOW
			Image:    nil,
			Thumbnail: nil,
			Video:     nil,
			Provider:  nil,
			Author:    nil,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: rozvrh[0] + " - "+ casy[0],
					Value:  linky[0],
					Inline: false,
				},
				{
					Name: rozvrh[1] + " - "+ casy[1],
					Value:  linky[1],
					Inline: false,
				},
				{
					Name: rozvrh[2] + " - "+ casy[2],
					Value:  linky[2],
					Inline: false,
				},
				{
					Name: rozvrh[3] + " - "+ casy[3],
					Value:  linky[3],
					Inline: false,
				},
				{
					Name: rozvrh[4] + " - "+ casy[4],
					Value:  linky[4],
					Inline: false,
				},
				{
					Name: rozvrh[5] + " - "+ casy[5],
					Value:  linky[5],
					Inline: false,
				},
				{
					Name: rozvrh[6] + " - "+ casy[6],
					Value:  linky[6],
					Inline: false,
				},
			},
		}
		return EmbedRozvrh
	} else {
		return EmbedRozvrhPrazdny
	}
}

var JeKoniec = discordgo.MessageEmbed{
	URL:   "",
	Type:  "",
	Title: "Už je koniec, palte dopiče",
	Description: "Beep Boop. Táto správa je automatizovaná",
	Timestamp: "",
	Color:     16711680, //RED
	Image:    nil,
	Thumbnail: nil,
	Video:     nil,
	Provider:  nil,
	Author:    nil,
	Fields: nil,
}




