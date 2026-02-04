package discord

import (
	"bot-discord-antitoxic/ai"
	"bot-discord-antitoxic/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(cfg *config.Config) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		isToxic, err := ai.IsToxic(m.Content, cfg.GeminiKey)
		if err != nil {
			fmt.Println("Error cek AI:", err)
			return
		}

		if isToxic {
			pesan := fmt.Sprintf("Waduh <@%s>, bahasanya dijaga ya! Jangan toksik di sini 😊", m.Author.ID)
			s.ChannelMessageSendReply(m.ChannelID, pesan, m.Reference())
		}
	}
}
