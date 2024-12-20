package main

import (
	"fmt"
	"strings"
)

var activeSessions = make(map[string]bool)

var cursedWords = []string{"Fuck", "Penis", "Boobs"}

func triggers(message string, number string, name string, messageID string, groupID string) {
	if len(groupID) <= 12 && groupID != "status" {
		personalChatCommands(message, number, name, messageID)
	} else if len(groupID) > 12 {
		groupChatCommands(message, number, name, messageID, groupID)
	}
}

func allowedGroups() {

}

func groupChatCommands(message string, number string, name string, messageID string, groupID string) {
	for _, group := range SpecificCommunityGroupsList {
		if groupID == group.JID.String() {
			for _, curse := range cursedWords {
				if strings.Contains(strings.ToLower(message), strings.ToLower(curse)) {

				}
			}
		} else {
			return
		}
	}
}

func personalChatCommands(message string, number string, name string, messageID string) {
	if message == "/start" && !activeSessions[number] {
		activeSessions[number] = true
		sendmessage(number, fmt.Sprintf(startMessage, name), messageID)
	} else if activeSessions[number] {
		if message == "/start" {
			sendmessage(number, "Your session has already been started. Please Use /help to know all the commands.", messageID)
		} else if message == "/verify" {
			sendmessage(number, verifyMessage, messageID)
		} else if strings.Contains(strings.ToLower(message), "/email") {
			if strings.Contains(strings.ToLower(message), "@ds.study.iitm.ac.in") {
				sendmessage(number, `You will get an OTP on the mail shortly. Send it in the chat
	
Ex:- /otp 0000

_*Warning:*_ Don't Share it with anyone otherwise you will not be able to participate in the event.`, messageID)
				otp = otpGenerator()
				fmt.Println(otp)
			} else {
				sendmessage(number, `Please send student ID
Ex:- /email 23f0000000@ds.study.iitm.ac.in`, messageID)
			}

		} else if strings.Contains(strings.ToLower(message), "/otp") {
			otpList := strings.Split(message, " ")
			var intOTP int
			fmt.Sscanf(otpList[1], "%d", &intOTP)
			if intOTP == otp {
				sendmessage(number, "OTP Verified Successfully", messageID)
				addParticipantinGroup("Test Community", number)
			} else {
				sendmessage(number, "OTP Verification Failed", messageID)
			}
		} else if message == "/help" {
			sendmessage(number, helpMessage, messageID)
		} else if message == "/info" {
			sendmessage(number, infoMessage, messageID)
		} else if message == "/contact" {
			sendmessage(number, contactMessage, messageID)
		} else if message == "/feedback" {
			sendmessage(number, feedbackMessage, messageID)
		} else if message == "/exit" {
			delete(activeSessions, number)
			sendmessage(number, "Session Ended", messageID)
		} else {
			sendmessage(number, `Incorrect Input. Please Choose from the given commands
			/start
			/help
			/verify
			/info
			/contact
			/feedback
			/exit`, messageID)
		}
	} else {
		return
	}
}
