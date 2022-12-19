package discord

import (
	"fmt"
	gogpt "github.com/StephanieAgatha/openai-discord-go/openai"
	"github.com/bwmarrin/discordgo"
)

var COMMANDS []*discordgo.ApplicationCommand = []*discordgo.ApplicationCommand{
	{
		Name:        "chat",
		Description: "Send some message to OpenAI",
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "Message to OpenAI",
				Required:    true,
			},
		},
	},
}

var commandHandlers = map[string]func(session *discordgo.Session, i *discordgo.InteractionCreate){

	"chat": func(session *discordgo.Session, i *discordgo.InteractionCreate) {

		fetchResponse := make(chan string)
		go func() {
			var userMsg string = fmt.Sprintf("%v", i.ApplicationCommandData().Options[0].Value)
			fetchResponse <- gogpt.SendToGPT(userMsg)
		}()
		session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Received message",
			},
		})
		for {
			select {
			case response := <-fetchResponse:
				session.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Content: &response,
				})
				return

			default:
				session.ChannelTyping(i.Interaction.ChannelID)
			}
		}

	},
}

func CommandInteractions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
