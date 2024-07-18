package engine

import (
	"bin-term-chat/layout"
	"bin-term-chat/model"
	"github.com/rivo/tview"
	"time"
)

func (e *Engine) register() tview.Primitive {

	req := new(model.ReqRegister)

	textview := tview.NewTextView().SetTextAlign(tview.AlignCenter)

	textview.SetTextStyle(StyleTextView)

	form := tview.NewForm().
		AddInputField("Name", "", 40, nil, func(text string) {
			req.Name = text
		}).
		AddInputField("Email", "", 40, nil, func(text string) {
			req.Email = text
		}).
		AddPasswordField("Password", "", 40, '*', func(text string) {
			req.Password = text
		}).
		AddButton("Register", func() {
			resp, err := e.handler.Register(req)
			if err != nil {
				textview.SetText(err.Error())
				return
			}

			textview.SetText(resp.Message)

			go func() {
				e.app.QueueUpdateDraw(func() {
					time.Sleep(1 * time.Second)
					e.pages.SwitchToPage("login")
				})
			}()

		}).
		AddButton("Login ?", func() {
			e.pages.SwitchToPage("login")
		})

	form.SetBorder(true).SetTitle("")

	return layout.Auth(form, textview, 11)
}
