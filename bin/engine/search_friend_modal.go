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
			id := inputField.GetText()

			if id == e.user.ID {
				return
			}

			httpresp, err := e.handler.GetUserWithId(id, e.token)
			if flex.GetItemCount() > 2 {
				flex.RemoveItem(flex.GetItem(2))
			}

			if err != nil {
				flex.AddItem(textviewError.SetText(err.Error()), 0, 1, false)
				return
			}

			data := httpresp.Data.(map[string]interface{})

			lbio := len(data["bio"].(string))
			endBio := ""
			if lbio > 20 {
				lbio = 20
				endBio = "..."
			}

			button := tview.NewButton(data["name"].(string) + " - " + data["bio"].(string)[:lbio] + endBio)
			button.SetBackgroundColor(tcell.ColorBlue)
			button.SetSelectedFunc(func() {
				e.setCompHub(id, data["name"].(string))
				if _, ok := e.compHub[id]; !ok {
					e.compHub["sidebar"].Chan <- model.User{
						Name: data["name"].(string),
						ID:   data["id"].(string),
					}
				}
				e.switchChatBox(id)()
			})
			e.setInputCapture(button.Box, func() {
				e.setFocus(inputField)
			})

			flex.AddItem(button, 0, 1, true)

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
			height: 7,
		},
	})
	modal.SetBorderPadding(1, 1, 1, 1)
}
