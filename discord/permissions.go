package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/SAl0nKA/RozvrhBOT/config"
)

func PermissionsCheck(Roles []string)bool{
	if ContainsIDs(Roles, config.RoleIDs) || config.RoleIDSstring == "" {
		return true
	}
	return false
}

func IsDM(Member *discordgo.Member)bool{
	if Member == nil{
		//s.ChannelMessageSend(m.ChannelID,"Tento príkaz je prístupny iba pre členov serveru s prislušnou rolou")
		return true
	}
	return false
}

func ContainsIDs(roles []string, ids []string) bool {
	for _, role := range roles {
		for _,id := range ids{
			if id == role {
				return true
			}
		}
	}
	return false
}

/*
func ContainsOne(element string, slice []string) bool {
	for _, ele := range slice {
		if element == ele{
			return true
		}
	}
	return false
}*/