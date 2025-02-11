package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define a struct for the card-related event
type Card struct {
	CardID        int    `json:"card_id"`
	CardName      string `json:"card_name"`
	StartPlayerID int    `json:"start_player_id"`
}

// Define a struct for wrapping the event list
type EventWrapper struct {
	EventList []map[string]json.RawMessage `json:"event_list"`
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

	// Iterate over events and extract card-related data
	for _, event := range eventData.EventList {
		for eventType, rawData := range event {
			// fmt.Println("Event type Found:", eventType) // debugging #TODO debug
			// fmt.Print("raw event data:", string(rawData)) // more debugging lines
			if eventType == "[Event_MoveCard.ext]" { // Only process card move events
				var card Card
				err := json.Unmarshal(rawData, &card)
				if err != nil {
					fmt.Println("Error parsing card event:", err)
					continue
				}
				fmt.Printf("Parsed Card: %+v\n", card)
			}
		}
	}
}
