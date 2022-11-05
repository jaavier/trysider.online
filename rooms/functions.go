package chat

type Message struct {
	Message   string `json:"message"`
	Sender    string `json:"sender"`
	RoomId    string `json:"roomId"`
	Timestamp int64  `json:"timestamp"`
}

func createRoom(roomId string) bool {
	return true
}

func joinRoom(roomId string, userId string) bool {
	return true
}

func sendMessage(roomId string, message string) bool {
	return true
}

func roomMembers() []string {
	return []string{}
}

func leaveRoom() bool {
	return true
}

func directMessagesReceived() []Message {
	return []Message{}
}

func directMessagesSent() []Message {
	return []Message{}
}
