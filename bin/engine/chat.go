package engine

import (
	"bin-term-chat/model"
	"encoding/json"
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
			case data := <-e.compHub["base"].Chan:
				user := data.(model.User)

				e.app.QueueUpdateDraw(func() {
					list.AddItem(user.Name, "", 0, func() {
						e.receiver = user.ID
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

func (e *Engine) chat() tview.Primitive {

	sidebar := e.listSidebar()
	banner := e.banner()

	flex := tview.NewFlex().
		AddItem(sidebar, 30, 1, true).
		AddItem(banner, 0, 3, false)

	e.setInputCapture(sidebar.Box, flex.GetItem(1))
	e.setInputCapture(banner.Box, sidebar)

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

	return flex

}
