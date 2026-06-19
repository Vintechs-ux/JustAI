package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func NewSession(token string) *discordgo.Session {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Gagal membuat session Discord:", err)
	}

	session.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentMessageContent | discordgo.IntentGuilds | discordgo.IntentGuildMembers

	return session
}
