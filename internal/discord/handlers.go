package discord

import (
	"fmt"
	"log"
	"strings"

	"github.com/Vintechs-ux/justAI/internal/commands"
	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	SentinelRole string
	PantheonRole string
	General      *commands.GeneralHandler
}

func NewHandler(SentinelRole, PantheonRole string, general *commands.GeneralHandler) *Handler {
	return &Handler{
		SentinelRole: SentinelRole,
		PantheonRole: PantheonRole,
		General:      general,
	}
}

func (h *Handler) OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Bot sukses login sebagai: %s#%s\n", s.State.User.Username, s.State.User.Discriminator)
	log.Printf("Terhubung ke %d server %v\n", len(r.Guilds), r.Guilds)
}

func (h *Handler) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "!ping" {
		level := GetPermissionLevel(s, m.GuildID, m.Author.ID, h.SentinelRole, h.PantheonRole)
		reply := fmt.Sprintf("Pong!\nHalo **%s**, role kamu: **%s**", m.Author.Username, level.String())
		s.ChannelMessageSend(m.ChannelID, reply)
		return
	}

	botMention := "<@" + s.State.User.ID + ">"
	if strings.Contains(m.Content, botMention) {
		userMessage := strings.TrimSpace(strings.ReplaceAll(m.Content, botMention, ""))
		if userMessage == "" {
			s.ChannelMessageSend(m.ChannelID, "Ada yang bisa dibantu blay? Mention gw yaak!")
			return
		}
		h.General.HandleChat(s, m, userMessage)
		return
	}
}
