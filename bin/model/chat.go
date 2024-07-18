package model

import "github.com/rivo/tview"

type WriteMessage struct {
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

type ReadMessage struct {
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

type CompHub struct {
	Comp tview.Primitive
	Chan chan any
}
