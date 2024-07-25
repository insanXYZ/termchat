package engine

import (
	"bin-term-chat/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) chatBox(idHub, title string) *tview.Flex {
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
	e.setInputCapture(flex.Box, func() {
		e.setFocus(e.chatCompLayout.Sidebar)
	})

	flex.SetTitle("💬 " + title)

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
