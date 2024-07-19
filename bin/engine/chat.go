package engine

import (
	"bin-term-chat/model"
	"encoding/json"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/gorilla/websocket"
	"github.com/rivo/tview"
	"log"
)

func (e *Engine) sendMessage(field *tview.InputField, chatBox *tview.TextView) {
	message := field.GetText()
	if message != "" && e.conn != nil {

		marshal, err := json.Marshal(model.WriteMessage{
			Message:  message,
			Receiver: e.receiver,
		})
		if err != nil {
			return
		}

		err = e.conn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			return
		}
		field.SetText("")
		chatBox.ScrollToEnd()
		e.app.SetFocus(field)
	}
}

func (e *Engine) readMessage() {
	for {
		if e.conn != nil {
			_, message, err := e.conn.ReadMessage()
			if err != nil {
				e.app.Stop()
				log.Println(err.Error())
			}

			var response model.ReadMessage

			err = json.Unmarshal(message, &response)
			if err != nil {
				e.app.Stop()
				log.Println(err.Error())
			}

			e.app.Stop()
			fmt.Println(response)
		}
	}
}

func (e *Engine) banner() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewTextView().SetText(model.APPNAME).SetTextAlign(tview.AlignCenter), 1, 1, false).
		AddItem(tview.NewBox(), 0, 1, false)
	flex.SetBorder(true)
	return flex
}

func (e *Engine) listSidebar() *tview.List {
	list := tview.NewList()
	list.AddItem("🔎 Search friend", "", 0, nil)
	list.AddItem("🌎 global", "", 0, nil)
	list.SetBorder(true)

	go func() {
		for {
			select {
			case data := <-e.compHub["sidebar"].Chan:
				user := data.(model.User)

				e.app.QueueUpdateDraw(func() {
					list.AddItem(user.Name, "", 0, func() {
						e.receiver = user.ID
						e.compHub["base"].Chan <- e.receiver
					})
				})
			}
		}
	}()

	return list
}

func (e *Engine) setInputCapture(list *tview.Box, fallback tview.Primitive) {
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			e.setFocus(fallback)
		}
		return event
	})
}

func (e *Engine) initChanCompChat() {
	e.setHub("sidebar", model.CompHub{
		Chan: make(chan any), // model.User
	})
	e.setHub("chat-switch", model.CompHub{
		Chan: make(chan any), // string
	})
	e.setHub("chat-global", model.CompHub{
		Comp: e.chatBox("global"),
		Chan: make(chan any), // model.ReadMessage
	})
}

func (e *Engine) chat() tview.Primitive {

	e.initChanCompChat()

	sidebar := e.listSidebar()
	banner := e.banner()

	flex := tview.NewFlex().
		AddItem(sidebar, 30, 1, true).
		AddItem(banner, 0, 3, false)

	e.setInputCapture(sidebar.Box, flex.GetItem(1))
	e.setInputCapture(banner.Box, sidebar)

	go func() {
		for {
			select {
			case id := <-e.compHub["chat"].Chan:
				if c, ok := e.compHub[id.(string)]; ok {
					e.app.QueueUpdateDraw(func() {
						flex.RemoveItem(flex.GetItem(1))
						flex.AddItem(c.Comp, 0, 3, true)
						e.setInputCapture(sidebar.Box, flex.GetItem(1))
					})
					e.setFocus(flex.GetItem(1))
				}
			}
		}
	}()

	return flex

}

func (e *Engine) chatBox(idHub string) tview.Primitive {
	chatBox := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true)
	chatBox.SetTitle(idHub)

	inputField := tview.NewInputField().
		SetLabelColor(tcell.ColorWhite).
		SetLabel("Message: ")

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			e.sendMessage(inputField, chatBox)
		}
	})

	sendButton := tview.NewButton("⌲").SetLabelColor(tcell.ColorWhite)
	sendButton.SetBackgroundColor(tcell.ColorGreen)

	inputFlex := tview.NewFlex().
		AddItem(inputField, 0, 1, true).
		AddItem(sendButton, 5, 0, false)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(chatBox, 0, 1, false).
		AddItem(inputFlex, 1, 0, true)
	flex.SetBorder(true)

	sendButton.SetSelectedFunc(func() {
		e.sendMessage(inputField, chatBox)
	})

	go func() {
		for {
			select {
			case msg := <-e.compHub[idHub].Chan:
				message := msg.(model.ReadMessage)

				headMessage := "[green]" + message.Time

				if message.Sender.ID == e.user.ID {
					headMessage = "You " + headMessage
				} else {
					headMessage = message.Sender.Name + " [blue]#" + message.Sender.ID + " " + headMessage
				}

				chatBox.Write([]byte(headMessage + "\n" + message.Message + "\n\n"))

			}
		}
	}()

	return flex

}
