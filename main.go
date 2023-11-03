package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"os/signal"
	"syscall"
	"strings"
	"strconv"
	"math/rand"

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
	_ = discordSession.UpdateListeningStatus(("!help  ♡UωU♡"))
	

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-channel

	discordSession.Close()
}

func messageCreate(session *discordgo.Session, message * discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "!help" {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf(`
		**Hi, I'm the uwu Bot!**

		**I support these commands:**
		**!help:** For the Info and Commands /ᐠ - ˕ -マ🚨
		**!pic <your uwu ID>:** For a random derivative of your uwu! ヾ( ˃ᴗ˂ )◞ • *✰

		My creators uwu Labs and Cat always open for any bug reports or suggestions!
		`))
		}

	if strings.HasPrefix(message.Content, "!pic") {
		uwuId, err := strconv.ParseUint(strings.TrimPrefix(message.Content, "!pic "), 10, 64) 
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		results := SqlConnect(uwuId)

		rand.Seed(time.Now().UnixNano())

		if results[0] == "no ID found" {
			session.ChannelMessageSend(message.ChannelID, "This uwu has no derivatives yet!  ( • ᴖ • ｡)")
			return
		} 

		randomIndex := rand.Intn(len(results))
		randomUwuPic := results[randomIndex]

		messages := []string{"Omg! look at da booba!  ( ๏ 人 ๏ )", "Dis uwu is so beautifuwu!  >ᴗ<", "Wish i had an uwu wike dis hehe!  ✪ ω ✪"}
		randomMessage := rand.Intn(len(messages))
		randomMessageId := messages[randomMessage]

		session.ChannelMessageSend(message.ChannelID, randomMessageId)
		session.ChannelMessageSend(message.ChannelID, randomUwuPic)
	}
}
