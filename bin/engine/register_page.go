package engine

import (
	"bin-term-chat/component"
	"bin-term-chat/layout"
	"bin-term-chat/model"
	"github.com/rivo/tview"
	"time"
)

func (e *Engine) register() tview.Primitive {

	req := new(model.ReqRegister)

	notify := component.CreateTextViewNotified()
	form := component.CreateForm(&component.Form{
		Border:          true,
		Title:           " 📝 Register ",
		BackgroundColor: model.ColorBackgroundBase,
	})
	form.AddFormItem(component.CreateFormItem(&component.FormItem{
		Label:      "Name",
		FieldWidth: 40,
		ChangedFunc: func(text string) {
			req.Name = text
		},
	}))
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
	form.AddButton("Register", func() {
		resp, err := e.handler.Register(req)
		if err != nil {
			notify.SetText(err.Error())
			return
		}

		notify.SetText(resp.Message)

		e.queueUpdateDraw(func() {
			time.Sleep(1 * time.Second)
			e.pages.SwitchToPage("login")
		})

	})
	form.AddButton("Login ?", func() {
		e.pages.SwitchToPage("login")
	})

	return layout.Auth(form, notify, 11)
}
