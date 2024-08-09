package engine

import (
	"bin-term-chat/component"
	"bin-term-chat/layout"
	"bin-term-chat/model"
	"github.com/rivo/tview"
	"time"
)

func (e *Engine) login() tview.Primitive {

	req := new(model.ReqLogin)
	notify := component.CreateTextViewNotified()

	form := component.CreateForm(&component.Form{
		Border:          true,
		Title:           " 🔑 Login ",
		BackgroundColor: model.ColorBackgroundBase,
	})

	form.AddFormItem(component.CreateFormItem(&component.FormItem{
		Label:      "Email",
		FieldWidth: 40,
		ChangedFunc: func(text string) {
			req.Email = text
		},
	}))
	form.AddFormItem(component.CreateFormItem(&component.FormItem{
		Label:      "Password",
		FieldWidth: 40,
		Mask:       '*',
		ChangedFunc: func(text string) {
			req.Password = text
		},
	}))

	form.AddButton("Login", func() {

		resp, err := e.handler.Login(req)
		if err != nil {
			notify.SetText(err.Error())
			return
		}

		notify.SetText(resp.Message)

		e.queueUpdateDraw(func() {
			time.Sleep(1 * time.Second)
		})

		e.setAuthEngine(resp.Data.(map[string]any))
		err = e.connectWebsocket()
		if err != nil {
			notify.SetText(err.Error())
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

	return layout.Auth(form, notify, 9)
}
