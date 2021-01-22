package discord

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type RozvrhEmbed struct {
	ChannelID 		string
	MessageID		string
	GuildID			string
	EditDay 		time.Weekday
}

var RozvrhEmbedy []*RozvrhEmbed

func ReturnEmbedRozvrh(rozvrh, casy, linky []string, day time.Weekday)discordgo.MessageEmbed{
	var EmbedRozvrhPrazdny = discordgo.MessageEmbed{
		Title: "Rozvrh",
		Description: "Dnes nie sú žiadne hodiny",
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
			Title: "Rozvrh - " + GetDayName(day),
			Color:     16766976, //YELLOW
			Fields: fields,
		}
		return EmbedRozvrh
	} else {
		return EmbedRozvrhPrazdny
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
	return ReturnEmbedRozvrh(sd.Hodiny, sd.Casy, sd.Linky,Rozvrh.EditDay)
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