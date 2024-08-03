package engine

import (
	"bin-term-chat/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) modalProfile() *tview.Form {

	user := new(model.UpdateUser)

	form := tview.NewForm()
	form.SetBackgroundColor(tcell.ColorGray)
	form.AddInputField("Name", e.user.Name, 40, nil, func(text string) {
		user.Name = text
	})
	form.AddInputField("Email", e.user.Email, 40, nil, func(text string) {
		user.Email = text
	})
	form.AddInputField("Bio", e.user.Bio, 40, nil, func(text string) {
		user.Bio = text
	})
	form.AddInputField("New Password", "", 40, nil, func(text string) {
		user.Password = text
	})
	form.AddButton("Save", func() {
		updateUser, err := e.handler.UpdateUser(user, e.token)
	})
	return form

}

func (e *Engine) showModalProfile() {
	modal := e.CreateModal(&modalConfig{
		title:           " 👤 Profile ",
		draggable:       false,
		border:          true,
		root:            e.modalProfile(),
		resizeable:      false,
		fallback:        e.chatCompLayout.Sidebar,
		backgroundColor: tcell.ColorGrey,
		size: size{
			x:      0,
			y:      0,
			width:  59,
			height: 15,
		},
	})
	modal.SetBorderPadding(1, 1, 1, 1)
}
