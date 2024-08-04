package engine

import (
	"bin-term-chat/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) modalSearchFriend() *tview.Flex {
	inputField := tview.NewInputField()
	flex := tview.NewFlex()

	e.setInputCapture(inputField.Box, func() {
		e.setFocus(flex.GetItem(2))
	})

	textview := tview.NewTextView().SetText("-|-").SetTextAlign(tview.AlignCenter)
	textview.SetBackgroundColor(tcell.ColorGray)

	textviewError := tview.NewTextView().SetTextAlign(tview.AlignCenter)
	textviewError.SetTextStyle(styleTextView)
	textviewError.SetBackgroundColor(tcell.ColorGray)

	inputField.SetFieldWidth(50)
	inputField.SetFieldBackgroundColor(tcell.ColorDarkGrey)
	inputField.SetPlaceholder("id...")

	inputField.SetPlaceholderTextColor(tcell.ColorDarkGreen)
	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			input := inputField.GetText()

			if input == "#"+e.user.ID || input == "" {
				return
			}

			httpresp, err := e.handler.GetUserWithId(input, e.token)
			if flex.GetItemCount() > 2 {
				flex.RemoveItem(flex.GetItem(2))
			}

			if err != nil {
				flex.AddItem(textviewError.SetText(err.Error()), 0, 1, false)
				return
			}

			data := httpresp.Data.([]any)
			list := tview.NewList()
			for _, item := range data {
				i := item.(map[string]any)
				list.AddItem(i["name"].(string)+" - "+i["id"].(string), i["bio"].(string)+"\n", 0, func() {
					if _, ok := e.compHub[i["id"].(string)]; !ok {
						e.compHub["sidebar"].Chan <- model.User{
							Name: i["name"].(string),
							ID:   i["id"].(string),
						}
					}
					e.setCompHub(i["id"].(string), i["name"].(string))
					e.switchChatBox(i["id"].(string))()
				})
			}

			flex.AddItem(list, 0, 1, false)

			e.setInputCapture(list.Box, func() {
				e.setFocus(inputField)
			})

		}
	})

	flex.SetDirection(tview.FlexRow)
	flex.AddItem(inputField, 1, 1, true)
	flex.AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorGrey), 1, 0, false)
	flex.AddItem(textview, 0, 1, false)

	return flex

}

func (e *Engine) showModalSearchFriend() {
	modal := e.CreateModal(&modalConfig{
		title:           " 🔎 search friend ",
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
			height: 10,
		},
	})
	modal.SetBorderPadding(1, 1, 1, 1)
}
