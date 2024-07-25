package model

import "github.com/rivo/tview"

type Type string

const (
	MessageGlobal  Type = "global"
	MessagePrivate Type = "private"
)

type WriteMessage struct {
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

type ReadMessage struct {
	Sender   *SenderMessage   `json:"sender"`
	Receiver *ReceiverMessage `json:"receiver"`
	Message  string           `json:"message"`
	Time     string           `json:"time"`
	Type     Type             `json:"type"`
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
	Comp *tview.Flex
	Chan chan any
}

type ChatComponentLayout struct {
	Banner  *tview.Flex
	Sidebar *tview.List
	ChatBox *tview.Flex
}
