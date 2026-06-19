package main

import (
	"github.com/Vintechs-ux/justAI/internal/config"
	"github.com/Vintechs-ux/justAI/internal/discord"
)

func main() {
	cfg := config.Load()

	session := discord.NewSession(cfg.DiscordToken)
}
