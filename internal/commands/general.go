package commands

import (
	"context"
	"log"

	"github.com/Vintechs-ux/justAI/internal/ai"
	"github.com/bwmarrin/discordgo"
)

const systemPrompt = `kamu adalah JustAI buatan
	Vintechs atai biasa di sebut Vins 
	(bilang pembuatmu jika ditanya atau user menyebut namanya) , kamu dibuat dengan bahasa golang , asisten Discord yang ramah dan membantu, kamu itu 
	representasi pemilih server discord bernama Justin yang seorang streamer game di youtube.
	jawab dengan bahasa indonesia santai, ramah ,singkat , dan jelas. jangan terlalu formal , nama 
	yotube kamu "Justin".`

type GeneralHandler struct {
	AIClient *ai.Client
}

func NewGeneralHandler(aiClient *ai.Client) *GeneralHandler {
	return &GeneralHandler{AIClient: aiClient}
}

func (h *GeneralHandler) HandleChat(s *discordgo.Session, m *discordgo.MessageCreate, userMessage string) {
	s.ChannelTyping(m.ChannelID)
	ctx := context.Background()

	answer, err := h.AIClient.Ask(ctx, systemPrompt, userMessage)
	if err != nil {
		log.Println("Error saat memanggil Groq:", err)
		s.ChannelMessageSend(m.ChannelID, "Waduh, saya lagi error! coba lagi nanti")
		return
	}

	s.ChannelMessageSend(m.ChannelID, answer)
}
