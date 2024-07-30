package engine

import (
	"bin-term-chat/model"
	"github.com/rivo/tview"
)

func (e *Engine) listSidebar() *tview.List {
	list := tview.NewList()
	list.AddItem(" 👤 Profile ", "", 0, nil)
	list.AddItem(" 🔎 Search friend ", "", 0, e.showModalSearchFriend)
	list.AddItem("", "", 0, nil)
	list.AddItem(" 🌎 global ", "", 0, e.switchChatBox("global"))
	list.SetTitle(" 🍔 Menu ")
	list.SetBorder(true)
	list.SetBorderPadding(1, 0, 0, 0)

	e.setInputCapture(list.Box, func() {
		if e.chatCompLayout.ChatBox == nil {
			e.setFocus(e.chatCompLayout.Banner)
			return
		}

		e.setFocus(e.chatCompLayout.ChatBox)
	})

	go func() {
		for {
			select {
			case data := <-e.compHub["sidebar"].Chan:
				user := data.(model.User)

				e.app.QueueUpdateDraw(func() {
					list.AddItem(" 🗿 "+user.Name, "", 0, e.switchChatBox(user.ID))
				})
			}
		}
	}()

	return list
}
