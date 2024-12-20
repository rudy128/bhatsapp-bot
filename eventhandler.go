package main

import (
	"fmt"

	"go.mau.fi/whatsmeow/types/events"
)

func eventHandler(evt interface{}) {
	// Check if the event is a message
	if messageEvent, ok := evt.(*events.Message); ok {
		// Handle actual messages
		number = messageEvent.Info.Sender.User
		messageID = messageEvent.Info.ID
		name = messageEvent.Info.PushName
		groupID = messageEvent.Info.Chat.User

		if extendedText := messageEvent.Message.GetExtendedTextMessage(); extendedText != nil {
			message = extendedText.GetText()
		} else {
			message = messageEvent.Message.GetConversation()
		}

		// Check if the message is not empty
		if message != "" {
			fmt.Printf("Message from %s: %s (ID: %s)\n", number, message, messageID)
			fmt.Printf("Group ID: %s\n", groupID)
			specificCommunityGroupsList("Test Community")
			triggers(message, number, name, messageID, groupID)
		}
	} else if _, ok := evt.(*events.Receipt); ok {
		// Handle read receipts separately
		// fmt.Println("\nReceived read receipt:", receiptEvent)
	} else if _, ok := evt.(*events.Connected); ok {
		// Handle connected event
		fmt.Println("Connected to WhatsApp server.")
	} else if _, ok := evt.(*events.OfflineSyncCompleted); ok {
		// Handle offline sync completed event
		fmt.Println("Offline sync completed.")
	} else if _, ok := evt.(*events.GroupInfo); ok {
		// fmt.Println("\nGroup Info:", groupInfo)
	} else {
		// Handle other events
		fmt.Printf("\nUnhandled event: %T\n", evt)
	}
}
