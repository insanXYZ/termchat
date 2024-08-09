package engine

import (
	"bin-term-chat/model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) chatBox(idHub, title string) *tview.Flex {
	chatBox := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWordWrap(true)
	chatBox.SetBorder(true)
	chatBox.SetTitle(" 💬 " + title + " ")

	inputField := tview.NewInputField()
	inputField.SetTitle(" Enter your message... ")
	inputField.SetBorder(true)
	inputField.SetTitleAlign(tview.AlignLeft)
	inputField.SetFieldBackgroundColor(tcell.ColorBlack)

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			e.sendMessage(inputField, chatBox)
		}
	})

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(chatBox, 0, 1, false).
		AddItem(inputField, 3, 0, true)
	e.setInputCapture(flex.Box, func() {
		e.setFocus(e.chatCompLayout.Sidebar)
	})

	flex.SetTitle("💬 " + title)
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
