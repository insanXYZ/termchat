package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) modalProfile() *tview.Form {
	form := tview.NewForm()
	form.AddInputField("Name", "", 40, nil, nil)
	form.AddInputField("Email", "", 40, nil, nil)
	form.AddInputField("New Password", "", 40, nil, nil)

	return form

}

func (e *Engine) showModalProfile() {
	modal := e.CreateModal(&modalConfig{
		title:           "🔎 search friend",
		draggable:       false,
		border:          true,
		root:            e.modalProfile(),
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
