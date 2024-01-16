package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"strconv"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
// var (
// 	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
// 	BotToken       = flag.String("token", "", "Bot access token")
// 	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
// )

var discordSession *discordgo.Session

func init() {
	var err error
	token := os.Getenv("DISCORD_BOT_TOKEN")
	
	discordSession, err = discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		fmt.Println("There was a problem with creating a Discord session,", err)
		return
	}
}

var (
	// integerOptionMinValue          = 1.0
	// dmPermission                   = false
	// defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand {
		{
			Name: "warning",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Warning from the uwu Bot",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type: discordgo.ApplicationCommandOptionString,
					Name: "message",
					Description: "The message to echo back",
					Required: true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"warning": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			message := i.ApplicationCommandData().Options[0].StringValue()
			embed := &discordgo.MessageEmbed{
				Description: message,
				Color:       0x00ff00, // Set the color to green
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})
		},
	}
)

func init() {
	discordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {

	discordSession.AddHandler(messageCreate)

	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	err := discordSession.Open()
	if err != nil {
		fmt.Println("Error connecting to the websocket", err)
		return
	}

	log.Println("Adding commands...")

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := discordSession.ApplicationCommandCreate(discordSession.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
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

	if strings.HasPrefix(message.Content, "!warn") {
		echoMessage := strings.TrimSpace(strings.TrimPrefix(message.Content, "!warn"))
		session.ChannelMessageSend(message.ChannelID, echoMessage)
		session.ChannelMessageDelete(message.ChannelID, message.ID)
	}

	if message.Content == "!uwu" {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf(`
		**Hi, I'm the uwu Bot!**

		**I support these commands:**
		**!uwu:** For the Info and Commands /ᐠ - ˕ -マ🚨
		**!deriv <your uwu ID> :** For a random derivative of your uwu! ヾ( ˃ᴗ˂ )◞ • *✰

		My creators uwu Labs and Cat always open for any bug reports or suggestions!
		`))
	}

	if strings.HasPrefix(message.Content, "!all") {
		_ = PostAll()

	 }

	if strings.HasPrefix(message.Content, "!deriv") {
        uwuIdStr := strings.TrimSpace(strings.TrimPrefix(message.Content, "!deriv"))
        var uwuId uint64
        var err error
        if uwuIdStr == "" {
            uwuId = 0
        } else {
            uwuId, err = strconv.ParseUint(uwuIdStr, 10, 64)
            if err != nil {
                fmt.Println("Error:", err)
                return
            }

			if uwuId <= 0 || uwuId > 9669 {
				session.ChannelMessageSend(message.ChannelID, "Incorrect uwu ID!  ( ^ ᴖ ^ )")
				return
			}
        }

		uwu, artist, rId := SqlConnect(uwuId)

		if uwu == "no ID found" {
			session.ChannelMessageSend(message.ChannelID, "This uwu has no derivatives yet!  ( • ᴖ • ｡)")
			return
		} else if uwuId == 0 {
			session.ChannelMessageSend(message.ChannelID, `Showing uwu ` + rId + ` deriv art by `+ artist +` !` )
			session.ChannelMessageSend(message.ChannelID, uwu)
			return
		}

		session.ChannelMessageSend(message.ChannelID, `Showing uwu ` + rId + ` deriv art by `+ artist +` !` )
		session.ChannelMessageSend(message.ChannelID, uwu)
	}


}
