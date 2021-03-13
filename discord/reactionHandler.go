package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

var Emojis = []string{"◀️", "▶️", "🔄", "❌"}

//krajšie to nebude
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

	if PermissionsCheck(member.Roles){
		channel, err := s.Channel(r.ChannelID)
		if err != nil {
			log.Println("Couldn't get the channel name: ", err)
		}

		switch r.Emoji.Name {
		case "◀️":
			log.Printf("[RozvrhBOT] User %s reacted with \"Previous day\" to rozvrh in channel %s",member.User.Username, channel.Name)
			s.MessageReactionRemove(r.ChannelID,r.MessageID, "◀️",r.UserID)

			embed := ReturnRozvrh(-1,r.MessageID)
			s.ChannelMessageEditEmbed(r.ChannelID,r.MessageID,&embed)
		case "▶️":
			log.Printf("[RozvrhBOT] User %s reacted with \"Next day\" to rozvrh in channel %s", member.User.Username, channel.Name)
			s.MessageReactionRemove(r.ChannelID,r.MessageID, "▶️",r.UserID)

			embed := ReturnRozvrh(1,r.MessageID)
			s.ChannelMessageEditEmbed(r.ChannelID,r.MessageID,&embed)
		case "🔄":
			log.Printf("[RozvrhBOT] User %s reacted with \"Refresh\" to rozvrh in channel %s", member.User.Username, channel.Name)
			s.ChannelMessageDelete(r.ChannelID,r.MessageID)

			embed := ReturnRozvrh(0,"")
			m, _ := s.ChannelMessageSendEmbed(r.ChannelID,&embed)
			AddReactions(s,m.ChannelID, m.ID)
			rozvrh := NewRozvrh(m.ChannelID, m.ID, m.GuildID,time.Now().Weekday())
			RozvrhEmbedy = append(RozvrhEmbedy,rozvrh)
		case "❌":
			log.Printf("[RozvrhBOT] User %s reacted with \"Delete\" to rozvrh in channel %s", member.User.Username, channel.Name)
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
			log.Println("[RozvrhBOT] Error adding reactions:",err)
		}
	}
}