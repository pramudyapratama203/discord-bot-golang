package discord

import (
	"bot-discord-antitoxic/ai"
	"bot-discord-antitoxic/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(cfg *config.Config) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Abaikan pesan dari bot itu sendiri
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Kirim teks ke AI buat dicek
		isToxic, err := ai.IsToxic(m.Content, cfg.GeminiKey)
		if err != nil {
			fmt.Println("Error cek AI:", err)
			return
		}

		// Kalau AI bilang toksik, kasih peringatan!
		if isToxic {
			pesan := fmt.Sprintf("Waduh <@%s>, bahasanya dijaga ya! Jangan toksik di sini 😊", m.Author.ID)
			s.ChannelMessageSendReply(m.ChannelID, pesan, m.Reference())
		}
	}
}
