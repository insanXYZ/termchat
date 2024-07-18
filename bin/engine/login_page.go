package engine

import (
	"bin-term-chat/layout"
	"bin-term-chat/model"
	"github.com/rivo/tview"
)

func (e *Engine) login() tview.Primitive {

	req := new(model.ReqLogin)
	textview := tview.NewTextView()

	form := tview.NewForm().
		AddInputField("Email", "", 40, nil, func(text string) {
			req.Email = text
		}).
		AddPasswordField("Password", "", 40, '*', func(text string) {
			req.Password = text
		}).
		AddButton("Login", func() {
			resp, err := e.handler.Login(req)
			if err != nil {
				textview.SetText(err.Error())
				return
			}

			textview.SetText(resp.Message)

			e.setAuthEngine(resp.Data.(map[string]any))
			err = e.connectWebsocket()
			if err != nil {
				textview.SetText(err.Error())
				return
			}
			e.setHub("base", model.CompHub{
				Chan: make(chan any),
			})

			e.winman.NewWindow().
				SetRoot(e.chat()).
				Maximize().
				Show().
				SetBorder(false)

			e.setRoot(e.winman)

		}).
		AddButton("Register ?", func() {
			e.pages.SwitchToPage("register")
		})

	form.SetBorder(true)

	return layout.Auth(form, textview, 9)
}
