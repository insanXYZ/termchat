package component

import (
	"bin-term-chat/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var styleTextView = tcell.StyleDefault.Italic(true)

func createTextView() (textview *tview.TextView) {
	textview = tview.NewTextView()
	return textview
}

func CreateTextViewNotified() (tv *tview.TextView) {
	tv = createTextView()
	tv.SetTextAlign(tview.AlignCenter)
	tv.SetTextStyle(styleTextView)
	tv.SetBackgroundColor(model.ColorBackgroundBase)
	return
}
