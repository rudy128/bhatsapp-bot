package main

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func sendmessage(number string, message string, messageID string) {

	targetJID, err := types.ParseJID(number + "@s.whatsapp.net")
	if err != nil {
		fmt.Println("Failed to parse JID:", err)
		return
	}
	msg := waE2E.Message{
		Conversation: proto.String(message),
	}
	fmt.Println(messageID)
	// if messageID != "" {
	// 	fmt.Println("Replying to message with ID:", messageID)

	// 	participant := number
	// 	quotedMsg := &waE2E.Message{
	// 		Conversation: proto.String(message),
	// 	}
	// 	msg = waE2E.Message{
	// 		ExtendedTextMessage: &waE2E.ExtendedTextMessage{
	// 			ContextInfo: &waE2E.ContextInfo{
	// 				StanzaID:      proto.String(messageID),
	// 				Participant:   proto.String(participant),
	// 				QuotedMessage: quotedMsg,
	// 			},

	// 			Text: proto.String(message),
	// 		},
	// 	}
	// } else {
	// 	fmt.Println("Sending normal message without reply context.")
	// }

	// fmt.Printf("Sending message: %s\n", message)
	// fmt.Printf("To: %s\n", targetJID)
	// fmt.Println(msg)
	if client != nil {
		// fmt.Printf("\nSending message to %s \nMessage: %s\n", targetJID, message)
		sendMessage, err := client.SendMessage(context.Background(), targetJID, &msg)
		if err != nil {
			fmt.Println("Failed to send message:", err)
		} else {
			fmt.Println("Message sent successfully:", sendMessage)
		}
	} else {
		fmt.Println("Client is not connected")
	}

}
