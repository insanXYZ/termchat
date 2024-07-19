package engine

import (
	"bin-term-chat/model"
	"encoding/json"
	"github.com/gdamore/tcell/v2"
	"github.com/gorilla/websocket"
	"github.com/rivo/tview"
	"log"
	"strings"
)

type SetPrivateMessage struct {
	id   string
	user model.User
}

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

			var readM model.ReadMessage

			err = json.Unmarshal(message, &readM)
			if err != nil {
				e.app.Stop()
				log.Println(err.Error())
			}

			if readM.Type == "global" {
				e.compHub["global"].Chan <- readM
			} else if readM.Type == "private" {

				set := SetPrivateMessage{}

				if readM.Sender.ID == e.user.ID {
					set.id = readM.Receiver.ID
					set.user = model.User{
						Name: readM.Receiver.Name,
						ID:   readM.Receiver.ID,
					}
				} else {
					set.id = readM.Sender.ID
					set.user = model.User{
						Name: readM.Sender.Name,
						ID:   readM.Sender.ID,
					}
				}

				if _, ok := e.compHub[set.id]; !ok {
					e.compHub[set.id] = model.CompHub{
						Comp: e.chatBox(set.id, set.user.Name),
						Chan: make(chan any),
					}
					e.compHub["sidebar"].Chan <- set.user
				}

				e.compHub[set.id].Chan <- readM

			}

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
	list.AddItem(strings.Repeat(string(tcell.RuneHLine), 30), "", 0, nil)
	list.AddItem("🌎 global", "", 0, func() {
		e.receiver = "global"
		e.compHub["chat"].Chan <- "global"
	})
	list.SetTitle("👥 Chat Menu")
	list.SetBorder(true)

	go func() {
		for {
			select {
			case data := <-e.compHub["sidebar"].Chan:
				user := data.(model.User)

				e.app.QueueUpdateDraw(func() {
					list.AddItem(user.Name, "", 0, func() {
						e.receiver = user.ID
						e.compHub["chat"].Chan <- e.receiver
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
	e.setHub("chat", model.CompHub{
		Chan: make(chan any), // string
	})
	e.setHub("global", model.CompHub{
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

func (e *Engine) chatBox(idHub string, title ...string) tview.Primitive {
	chatBox := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true)

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

	t := idHub

	if len(title) != 0 {
		t = title[0]
	}

	flex.SetTitle("💬 " + t)

	sendButton.SetSelectedFunc(func() {
		e.sendMessage(inputField, chatBox)
	})

	go func() {
		for {
			select {
			case msg := <-e.compHub[idHub].Chan:
				message := msg.(model.ReadMessage)

				headMessage := "[green]" + message.Time + ":"

				if message.Sender.ID == e.user.ID {
					headMessage = "You " + headMessage
				} else {
					headMessage = message.Sender.Name + " [blue]#" + message.Sender.ID + " " + headMessage
				}

				e.app.QueueUpdateDraw(func() {
					chatBox.Write([]byte(headMessage + "\n[white]" + message.Message + "\n\n"))
				})

			}
		}
	}()

	return flex

}
