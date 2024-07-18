package model

type ReadMessage struct {
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

type BroadcastMessage struct {
	Sender   *SenderMessage
	Message  string
	Receiver string
}

type SendMessage struct {
	Sender   *SenderMessage   `json:"sender"`
	Receiver *ReceiverMessage `json:"receiver"`
	Message  string           `json:"message"`
	Type     string           `json:"type"`
}

type WriteMessage struct {
	Sender   *SenderMessage   `json:"sender"`
	Receiver *ReceiverMessage `json:"receiver"`
	Message  string           `json:"message"`
	Time     string           `json:"time"`
	Type     string           `json:"type"`
}

type SenderMessage struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type ReceiverMessage struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
