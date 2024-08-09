package component

import (
	"bin-term-chat/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Form struct {
	Border          bool
	Title           string
	BackgroundColor tcell.Color
}

type FormItem struct {
	Label, Text string
	FieldWidth  int
	Mask        rune
	ChangedFunc func(text string)
}

func CreateFormItem(formItemConfig *FormItem) (inf *tview.InputField) {
	inf = tview.NewInputField()
	inf.SetLabel(formItemConfig.Label)
	inf.SetText(formItemConfig.Text)
	inf.SetFieldWidth(formItemConfig.FieldWidth)
	inf.SetChangedFunc(formItemConfig.ChangedFunc)
	inf.SetMaskCharacter(formItemConfig.Mask)

	inf.SetBackgroundColor(model.ColorBackgroundInputField)
	inf.SetFieldBackgroundColor(model.ColorValueInputField)

	return
}

func CreateForm(formConfig *Form) (f *tview.Form) {
	f = tview.NewForm()
	f.SetTitle(formConfig.Title)
	f.SetBorder(formConfig.Border)
	f.SetBackgroundColor(formConfig.BackgroundColor)
	f.SetButtonBackgroundColor(model.ColorBackgroundButton)
	f.SetButtonTextColor(model.ColorLabelButton)

	return

}
