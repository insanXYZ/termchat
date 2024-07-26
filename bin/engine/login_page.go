package engine

import (
	"bin-term-chat/layout"
	"bin-term-chat/model"
	"github.com/rivo/tview"
	"time"
)

func (e *Engine) login() tview.Primitive {

	req := new(model.ReqLogin)
	textview := tview.NewTextView()
	textview.SetTextAlign(tview.AlignCenter)
	textview.SetTextStyle(styleTextView)

	form := tview.NewForm()
	form.AddInputField("Email", "", 40, nil, func(text string) {
		req.Email = text
	})
	form.AddPasswordField("Password", "", 40, '*', func(text string) {
		req.Password = text
	})
	form.AddButton("Login", func() {
		resp, err := e.handler.Login(req)
		if err != nil {
			textview.SetText(err.Error())
			return
		}

		textview.SetText(resp.Message)

		e.queueUpdateDraw(func() {
			time.Sleep(1 * time.Second)
		})

		e.setAuthEngine(resp.Data.(map[string]any))
		err = e.connectWebsocket()
		if err != nil {
			textview.SetText(err.Error())
			return
		}

		go e.readMessage()

		e.winman.NewWindow().
			SetRoot(e.chat()).
			Maximize().
			Show().
			SetBorder(false)

		e.setRoot(e.winman)

	})
	form.AddButton("Register ?", func() {
		e.pages.SwitchToPage("register")
	})

	form.SetBorder(true)
	form.SetTitle(" 🔑 Login ")

	return layout.Auth(form, textview, 9)
}
