package discord

import (
	"github.com/SAl0nKA/RozvrhBOT/config"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type CommandType int

const (
	Help CommandType = iota
	Ping
	Pong
	Hod
	Dalsia
	Rozvrh
	Github
	Null
)

var CommandTypeStringMapping = map[string]CommandType{
	"help":  	Help,
	"ping":   	Ping,
	"pong":   	Pong,
	"hod":    	Hod,
	"dalsia":	Dalsia,
	"rozvrh": 	Rozvrh,
	"github":	Github,
	"":      	Null,
}

func GetCommandType(arg string) CommandType {
	for str, cmd := range CommandTypeStringMapping {
		if strings.ToLower(arg) == (config.BotPrefix + str) {
			return cmd
		}
	}
	return Null
}

//TODO pridať príkaz na pridanie suplovania
//TODO pridať príkaz na link na github
//TODO pridať footery správam
func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		args := strings.Fields(m.Content)
		switch GetCommandType(args[0]) {
		case Help:
			log.Printf("[RozvrhBOT] Reacting to command \"%shelp\"", config.BotPrefix)
			GetHelp(s,m)
		case Ping:
			log.Printf("[RozvrhBOT] Reacting to command \"%sping\"", config.BotPrefix)
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		case Pong:
			log.Printf("[RozvrhBOT] Reacting to command \"%spong\"", config.BotPrefix)
			if !PermissionsCheck(m.Member.Roles){
				NemasOpravnenie(s,m)
				return
			}
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		case Hod:
			log.Printf("[RozvrhBOT] Reacting to command \"%shod\"", config.BotPrefix)
			CommandHod(s, m)
		case Dalsia:
			log.Printf("[RozvrhBOT] Reacting to command \"%sdalsia\"", config.BotPrefix)
			CommandDalsia(s, m)
		case Rozvrh:
			log.Printf("[RozvrhBOT] Reacting to command \"%srozvrh\"", config.BotPrefix)
			CommandRozvrh(s, m)
		case Github:
			log.Printf("[RozvrhBOT] Reacting to command \"%sgithub\"",config.BotPrefix)
			CommandGithub(s,m)
		}
	}
}