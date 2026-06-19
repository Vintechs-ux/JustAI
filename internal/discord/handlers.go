package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	SentinelRole string
	PantheonRole string
}

func NewHandler(SentinelRole, PantheonRole string) *Handler {
	return &Handler{
		SentinelRole: SentinelRole,
		PantheonRole: PantheonRole,
	}
}

func (h *Handler) OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Bot sukses login sebagai: %s#%s\n", s.State.User.Username, s.State.User.Discriminator)
	log.Printf("Terhubung ke %d server\n", len(r.Guilds))
}

func (h *Handler) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "!ping" {
		level := GetPermissionLevel(s, m.GuildID, m.Author.ID, h.SentinelRole, h.PantheonRole)
		reply := fmt.Sprintf("Pong!\nHalo **%s**, role kamu: **%s**", m.Author.Username, level.String())
		s.ChannelMessageSend(m.ChannelID, reply)
	}
}
