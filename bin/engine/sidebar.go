package engine

import (
	"bin-term-chat/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

func (e *Engine) listSidebar() *tview.List {
	list := tview.NewList()
	list.AddItem("🔎 Search friend", "", 0, e.showModalSearchFriend)
	list.AddItem(strings.Repeat(string(tcell.RuneHLine), 30), "", 0, nil)
	list.AddItem("🌎 global", "", 0, e.switchChatBox("global"))
	list.SetTitle("👥 Chat Menu")
	list.SetBorder(true)

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
					list.AddItem(user.Name, "", 0, e.switchChatBox(user.ID))
				})
			}
		}
	}()

	return list
}
