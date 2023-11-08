package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	token := os.Getenv("DISCORD_BOT_TOKEN")
	
	discordSession, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		fmt.Println("There was a problem with creating a Discord session,", err)
		return
	}

	discordSession.AddHandler(messageCreate)

	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	err = discordSession.Open()
	if err != nil {
		fmt.Println("Error connecting to the websocket", err)
		return
	}

	fmt.Println("Bot is activated!")

	defer discordSession.Close()
	_ = discordSession.UpdateListeningStatus(("!uwu  ♡UωU♡"))
	

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-channel

	discordSession.Close()
}

func messageCreate(session *discordgo.Session, message * discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "!uwu" {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf(`
		**Hi, I'm the uwu Bot!**

		**I support these commands:**
		**!uwu:** For the Info and Commands /ᐠ - ˕ -マ🚨
		**!deriv <your uwu ID>:** For a random derivative of your uwu! ヾ( ˃ᴗ˂ )◞ • *✰

		My creators uwu Labs and Cat always open for any bug reports or suggestions!
		`))
		}

	if strings.HasPrefix(message.Content, "!deriv") {
		uwuId, err := strconv.ParseUint(strings.TrimPrefix(message.Content, "!deriv "), 10, 64) 
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		uwu, artist := SqlConnect(uwuId)

		if uwu == "no ID found" {
			session.ChannelMessageSend(message.ChannelID, "This uwu has no derivatives yet!  ( • ᴖ • ｡)")
			return
		} 

		session.ChannelMessageSend(message.ChannelID, `Showing uwu ` + strings.TrimPrefix(message.Content, "!deriv ") + ` deriv art by `+ artist +` !` )
		session.ChannelMessageSend(message.ChannelID, uwu)
	}
}
