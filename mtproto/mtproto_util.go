/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mtproto

////////////////////////////////////////////////////////////////////////////////
func ToBool(b bool) *Bool {
	if b {
		return MakeBool(&TLBoolTrue{})
	} else {
		return MakeBool(&TLBoolFalse{})
	}
}

func FromBool(b *Bool) bool {
	switch b.Payload.(type) {
	case *Bool_BoolTrue:
		return true
	default:
		return false
	}
}

//////////////////////////////////////////////////////////////////////////////////
// 太麻烦了
func GetUserIdListByChatParticipants(participants *TLChatParticipants) []int32 {
	chatUserIdList := []int32{}

	// TODO(@benqi):  nil check
	for _, participant := range participants.GetParticipants() {
		switch participant.Payload.(type) {
		case *ChatParticipant_ChatParticipant:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipant().GetUserId())
		case *ChatParticipant_ChatParticipantAdmin:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipantAdmin().GetUserId())
		case *ChatParticipant_ChatParticipantCreator:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipantCreator().GetUserId())
		}
	}
	return chatUserIdList
}
