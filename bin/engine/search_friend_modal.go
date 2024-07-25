package engine

import (
	"bin-term-chat/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) modalSearchFriend() *tview.Flex {
	inputField := tview.NewInputField()
	result := tview.NewFlex()

	inputField.SetFieldWidth(50)
	inputField.SetFieldBackgroundColor(tcell.ColorDarkGrey)
	inputField.SetPlaceholder("id...")

	inputField.SetPlaceholderTextColor(tcell.ColorDarkGreen)
	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			id := inputField.GetText()

			if id == e.user.ID {
				return
			}

			httpresp, err := e.handler.GetUserWithId(id, e.token)
			if err != nil {
				if result.GetItemCount() > 0 {
					result.RemoveItem(result.GetItem(0))
				}
				result.AddItem(tview.NewTextView().SetText(err.Error()), 0, 1, true)
				return
			}

			data := httpresp.Data.(map[string]interface{})
			if result.GetItemCount() > 0 {
				result.RemoveItem(result.GetItem(0))
			}
			result.AddItem(tview.NewButton(data["name"].(string)).SetSelectedFunc(func() {
				e.setCompHub(id, data["name"].(string))
				e.compHub["sidebar"].Chan <- model.User{
					Name: data["name"].(string),
					ID:   data["id"].(string),
				}
				e.switchChatBox(id)()
			}), 0, 1, true)

		}
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(inputField, 1, 1, true)
	flex.AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorGrey), 1, 0, false) // Add this line for the gap
	flex.AddItem(result, 0, 1, false)

	return flex

}

func (e *Engine) showModalSearchFriend() {
	modal := e.CreateModal(&modalConfig{
		title:           "🔎 search friend",
		draggable:       false,
		border:          true,
		root:            e.modalSearchFriend(),
		resizeable:      false,
		fallback:        e.chatCompLayout.Sidebar,
		backgroundColor: tcell.ColorGrey,
		size: size{
			x:      0,
			y:      0,
			width:  50,
			height: 7,
		},
	})
	modal.SetBorderPadding(1, 1, 1, 1)
}
