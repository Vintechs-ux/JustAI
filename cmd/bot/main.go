package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Vintechs-ux/justAI/internal/ai"
	"github.com/Vintechs-ux/justAI/internal/commands"
	"github.com/Vintechs-ux/justAI/internal/config"
	"github.com/Vintechs-ux/justAI/internal/discord"
)

func main() {
	cfg := config.Load()

	aiClient := ai.NewClient(cfg.GroqAPIKey)
	generalHandler := commands.NewGeneralHandler(aiClient)

	session := discord.NewSession(cfg.DiscordToken)
	handler := discord.NewHandler(cfg.SentinelRole, cfg.PhantomRole, generalHandler)

	session.AddHandler(handler.OnReady)
	session.AddHandler(handler.OnMessageCreate)

	err := session.Open()
	if err != nil {
		log.Fatal("Gagal membuka koneksi ke Discord:", err)
	}
	defer session.Close()

	log.Println("JustAI sedang berjalan...")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("JustAI dihentikan")
}
