package engine

import (
	"bin-term-chat/model"
	"encoding/json"
	"github.com/gdamore/tcell/v2"
	"github.com/gorilla/websocket"
	"github.com/rivo/tview"
	"log"
)

type setPrivateMessage struct {
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

			if readM.Type == model.MessageGlobal {
				e.compHub["global"].Chan <- readM
			} else if readM.Type == model.MessagePrivate {

				set := setPrivateMessage{}

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
	e.setInputCapture(flex.Box, func() {
		e.setFocus(e.chatCompLayout.Sidebar)
	})
	return flex
}

func (e *Engine) switchChatBox(idHub string) func() {
	return func() {
		e.receiver = idHub
		e.compHub["chat"].Chan <- e.receiver
	}
}

func (e *Engine) setInputCapture(box *tview.Box, f func()) {
	box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			f()
		}
		return event
	})
}

func (e *Engine) initChanCompChat() {

	e.chatCompLayout = &model.ChatComponentLayout{
		Banner:  e.banner(),
		Sidebar: e.listSidebar(),
		ChatBox: nil,
	}
	e.setChanHub("sidebar")
	e.setChanHub("chat")
	e.setCompHub("global", "global")
}

func (e *Engine) chat() *tview.Flex {

	e.initChanCompChat()

	sidebar := e.chatCompLayout.Sidebar
	banner := e.chatCompLayout.Banner

	flex := tview.NewFlex().
		AddItem(sidebar, 30, 1, true).
		AddItem(banner, 0, 3, false)

	go func() {
		for {
			select {
			case id := <-e.compHub["chat"].Chan:
				if c, ok := e.compHub[id.(string)]; ok {
					e.app.QueueUpdateDraw(func() {
						flex.RemoveItem(flex.GetItem(1))
						flex.AddItem(c.Comp, 0, 3, true)
						e.chatCompLayout.ChatBox = c.Comp
					})
					e.setFocus(flex.GetItem(1))
				}
			}
		}
	}()

	return flex

}
