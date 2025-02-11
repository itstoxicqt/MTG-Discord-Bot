ackage main

// import bot "example.com/hello_world_bot/Bot"
import (
	"example.com/hello_world_bot/Bot"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Starting Discord bot...")
	bot.Run() // Calls Run() from bot/bot.go
}
