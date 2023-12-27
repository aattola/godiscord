package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMessages | discordgo.IntentDirectMessages)

	err = discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	fmt.Println("Bot is running. Press CTRL-C to exit.")
	defer discord.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if b := strings.Contains(m.Content, "homo"); !b {
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, "eisaatana sun kanssas nyt")
	if err != nil {
		return
	}
}
