package engine

import (
	"bin-term-chat/model"
	"github.com/epiclabs-io/winman"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) modalProfile(window *winman.WindowBase) *tview.Form {

	user := new(model.UpdateUser)

	form := tview.NewForm()
	form.SetBackgroundColor(tcell.ColorGray)
	nameField := tview.NewInputField().SetLabel("Name").SetText(e.user.Name).SetFieldWidth(40).SetChangedFunc(func(text string) {
		user.Name = text
	})
	emailField := tview.NewInputField().SetLabel("Email").SetText(e.user.Email).SetFieldWidth(40).SetChangedFunc(func(text string) {
		user.Email = text
	})
	bioField := tview.NewInputField().SetLabel("Bio").SetText(e.user.Bio).SetFieldWidth(40).SetChangedFunc(func(text string) {
		user.Bio = text
	})
	passwordField := tview.NewInputField().SetLabel("New Password").SetText("").SetFieldWidth(40).SetChangedFunc(func(text string) {
		user.Password = text
	})

	form.AddFormItem(nameField)
	form.AddFormItem(emailField)
	form.AddFormItem(bioField)
	form.AddFormItem(passwordField)

	form.AddButton("Save", func() {
		updateUser, err := e.handler.UpdateUser(user, e.token)
		if err != nil {
			return
		}

		data := updateUser.Data.(map[string]interface{})

		e.user.Name = data["name"].(string)
		e.user.Email = data["email"].(string)
		e.user.Bio = data["bio"].(string)

		e.closeModal(window, e.chatCompLayout.Sidebar)
	})
	return form

}

func (e *Engine) showModalProfile() {
	modal := e.CreateModal(&modalConfig{
		title:           " 👤 Profile ",
		draggable:       false,
		border:          true,
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
	modal.SetRoot(e.modalProfile(modal))
	modal.SetBorderPadding(1, 1, 1, 1)
}
