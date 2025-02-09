package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
)

// Get the bot token from the .env file
func getBotToken() string {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("Bot token is missing from .env")
	}
	return token
}

func Run() {
	// Load token securely
	botToken := getBotToken()

	// Create a new Discord session
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Add message handler
	discord.AddHandler(handleMessage)

	// Open WebSocket connection
	err = discord.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer discord.Close()

	fmt.Println("Bot is running... Press Ctrl+C to exit.")

	// Graceful shutdown on Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func handleMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.HasPrefix(message.Content, "!help"):
		discord.ChannelMessageSend(message.ChannelID, "Hello world!")
	case strings.HasPrefix(message.Content, "!bye"):
		discord.ChannelMessageSend(message.ChannelID, "Goodbye!")
	case strings.HasPrefix(message.Content, "!ping"):
		discord.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}
