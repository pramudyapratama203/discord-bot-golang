package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bot-discord-antitoxic/config"
	"bot-discord-antitoxic/discord"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// 1. Load config dari .env
	cfg := config.LoadConfig()

	// 2. Inisialisasi session Discord
	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatal("Gagal membuat session Discord : ", err)
	}

	// 3. Pasang Handler buat dengerin chat
	dg.AddHandler(discord.MessageHandler(cfg))

	// 4. Handler buat tanda kalau bot sudah login
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Bot sudah online sebagai: %v#%v\n", r.User.Username, r.User.Discriminator)
	})

	// 5. Buka koneksi
	err = dg.Open()
	if err != nil {
		log.Fatal("Gagal membuka koneksi ke Discord:", err)
	}

	fmt.Println("Bot standby. Tekan CTRL+C untuk berhenti.")

	// Menunggu sinyal tutup (CTRL+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	dg.Close()
}
