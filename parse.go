package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define a struct for the original card movement event
type Card struct {
	CardID        int    `json:"card_id"`
	CardName      string `json:"card_name"`
	StartPlayerID int    `json:"start_player_id"`
}

// Define a struct for SetCardAttr event
type CardAttr struct {
	ZoneName  string `json:"zone_name"`
	CardID    int    `json:"card_id"`
	Attribute string `json:"attribute"`
}

// Define a struct for Join event to get player name
type JoinEvent struct {
	PlayerProperties struct {
		PlayerID int `json:"player_id"`
		UserInfo struct {
			Name string `json:"name"`
		} `json:"user_info"`
	} `json:"player_properties"`
}

// Define a struct for individual events
type Event struct {
	EventList []map[string]json.RawMessage `json:"event_list"`
	Seconds   int                          `json:"seconds_elapsed"`
}

// Define a struct for the top-level JSON
type EventWrapper struct {
	ReplayID string  `json:"replay_id"`
	Events   []Event `json:"event_list"`
}

func main() {
	// Open the JSON file
	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode JSON from file
	var eventData EventWrapper
	err = json.NewDecoder(file).Decode(&eventData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Validation
	if len(eventData.Events) == 0 {
		fmt.Println("Warning: No events found in the file")
		return
	}

	// Maps to store player ID to name and card ID to name mappings
	playerNames := make(map[int]string)
	cardNames := make(map[int]string)

	// First pass: Collect player names and card names
	for _, event := range eventData.Events {
		for _, eventItem := range event.EventList {
			for eventType, rawData := range eventItem {
				switch eventType {
				case "[Event_Join.ext]":
					var join JoinEvent
					err := json.Unmarshal(rawData, &join)
					if err != nil {
						fmt.Println("Error parsing join event:", err)
						continue
					}
					playerNames[join.PlayerProperties.PlayerID] = join.PlayerProperties.UserInfo.Name

				case "[Event_MoveCard.ext]":
					var card Card
					err := json.Unmarshal(rawData, &card)
					if err != nil {
						fmt.Println("Error parsing card event:", err)
						continue
					}
					cardNames[card.CardID] = card.CardName
				}
			}
		}
	}

	// Second pass: Process card events using the player and card names maps
	for _, event := range eventData.Events {
		for _, eventItem := range event.EventList {
			for eventType, rawData := range eventItem {
				switch eventType {
				case "[Event_MoveCard.ext]":
					var card Card
					err := json.Unmarshal(rawData, &card)
					if err != nil {
						fmt.Println("Error parsing card event:", err)
						continue
					}
					playerName, ok := playerNames[card.StartPlayerID]
					if !ok {
						playerName = "Unknown"
					}
					fmt.Printf("Card Move Event at %d seconds: ID=%d, Name=%s, StartPlayer=%s\n",
						event.Seconds, card.CardID, card.CardName, playerName)

				case "[Event_SetCardAttr.ext]":
					var cardAttr CardAttr
					err := json.Unmarshal(rawData, &cardAttr)
					if err != nil {
						fmt.Println("Error parsing card attribute event:", err)
						continue
					}
					// Look up card name, use "Unknown" if not found
					cardName, ok := cardNames[cardAttr.CardID]
					if !ok {
						cardName = "Unknown"
					}
					fmt.Printf("Card Attribute Event at %d seconds: CardID=%d, Name=%s, Zone=%s, Attribute=%s\n",
						event.Seconds, cardAttr.CardID, cardName, cardAttr.ZoneName, cardAttr.Attribute)
				}
			}
		}
	}
}
