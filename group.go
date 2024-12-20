package main

import (
	"fmt"
	"strings"

	"go.mau.fi/whatsmeow/types"
)

var AllGroupList []*types.GroupInfo

var SpecificCommunityGroupsList []*types.GroupLinkTarget

func groupList() {

	if client != nil {

		var err error
		AllGroupList, err = client.GetJoinedGroups()

		if err != nil {

			fmt.Println("Failed to get group")

		} else {
			fmt.Printf("Gathered %v Groups\n", len(AllGroupList))
		}

	} else {

		fmt.Println("Devive must be scanned first")

	}

}

func specificCommunityGroupsList(groupName string) {
	groupID := group(groupName)
	var err error
	SpecificCommunityGroupsList, err = client.GetSubGroups(groupID)
	if err != nil {
		fmt.Println("Couldn't Find List of SubGroups")
	}
}

func group(groupName string) types.JID {

	if AllGroupList == nil {
		groupList()
	}
	for _, group := range AllGroupList {
		if strings.Contains(group.GroupName.Name, groupName) {
			return group.JID
		} else {
		}
	}

	return types.JID{}
}

func addParticipantinGroup(groupName string, number string) {
	groupID := group(groupName)
	participantNumber, err := types.ParseJID(number + "@s.whatsapp.net")
	if err != nil {
		fmt.Println("Failed to parse JID:", err)
		return
	}
	if client != nil {
		participantAdded, err := client.UpdateGroupParticipants(groupID, []types.JID{participantNumber}, "add")
		if err != nil {
			fmt.Println("Failed to add participant:", err)
		} else {
			fmt.Println("Participant added successfully:", participantAdded)
		}
	} else {
		fmt.Println("Client is not connected")
	}
}
