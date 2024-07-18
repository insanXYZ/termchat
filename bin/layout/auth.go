package layout

import "github.com/rivo/tview"

func Auth(comp, textview tview.Primitive, size int) tview.Primitive {
	flex := tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(comp, size, 1, true).
			AddItem(textview, 0, 1, false), 0, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)

	return flex
}
