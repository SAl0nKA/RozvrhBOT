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
<<<<<<< HEAD
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
		Title: "Rozvrh",
<<<<<<< HEAD
		Description: "Dnes nie sú žiadne hodiny, jeb na to",
		Timestamp: "",
=======
		Description: "Dnes nie sú žiadne hodiny",
>>>>>>> 8754588 (Save)
		Color: 16711680, //RED
	}
	if day == 0{
		day = time.Now().Weekday()
	}
	if day <= 5 && day != 0 {
		fields := []*discordgo.MessageEmbedField{}
		for i := 0;i<len(rozvrh);i++{
			if rozvrh[i] == ""{
				continue
			}
			f := discordgo.MessageEmbedField{
				Name:   rozvrh[i] + " - " + casy[i],
				Value:  linky[i],
				Inline: false,
			}
			fields = append(fields,&f)
		}
		var EmbedRozvrh = discordgo.MessageEmbed{
			Title: "Rozvrh - " + GetDay(day),
			Color:     16766976, //YELLOW
			Fields: fields,
		}
		return EmbedRozvrh
	} else {
		return EmbedRozvrhPrazdny
	}
}

var JeKoniec = discordgo.MessageEmbed{
<<<<<<< HEAD
	URL:   "",
	Type:  "",
	Title: "Už je koniec, palte dopiče",
=======
	Title: "Je koniec. Ste voľní!",
>>>>>>> 8754588 (Save)
	Description: "Beep Boop. Táto správa je automatizovaná",
	Color:     16711680, //RED
=======
	s.ChannelMessageSendEmbed(m.ChannelID, &EmbedHelp)
>>>>>>> a833281 (Upratanie kodu)
}




