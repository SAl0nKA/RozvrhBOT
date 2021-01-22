package bot

import (
	"../config"
	"../discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Println("[RozvrhBOT] error creating Discord session,", err)
		return
	}

	goBot.AddHandler(Innit)
	goBot.AddHandler(discord.HandleCommand)
	goBot.AddHandler(discord.HandleReaction)

	goBot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuildMessageReactions)
	log.Println("[RozvrhBOT] Opening connection")
	err = goBot.Open()
	if err != nil {
		log.Println("[RozvrhBOT] error opening connection,", err)
		return
	}


	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	goBot.Close()
}

func Innit(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("[RozvrhBOT] Updating status")
	s.UpdateStatus(0, fmt.Sprintf("%shelp",config.BotPrefix))

	//Lesson announcing
	if config.DefaultChannelsID != nil {
		log.Println("[RozvrhBOT] Checking for current day")
		t := time.Now().Weekday()
		if t <= 5 && t != 0 {
			log.Println("[RozvrhBOT] Running automatic lesson announcing")
			go discord.HodAnnounce(s)
		} else {
			log.Println("[RozvrhBOT] Not running automatic lesson announcing")
		}
	} else {
		log.Println("[RozvrhBOT] Not running automatic lesson announcing")
	}
}
