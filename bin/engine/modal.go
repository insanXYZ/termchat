package engine

import (
	"github.com/epiclabs-io/winman"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type size struct {
	x, y, width, height int
}

type modalConfig struct {
	root                          tview.Primitive
	title                         string
	border, draggable, resizeable bool
	fallback                      tview.Primitive
	backgroundColor               tcell.Color
	size                          size
}

func (e *Engine) CreateModal(config *modalConfig) *winman.WindowBase {
	wnd := winman.NewWindow().Show()
	wnd.SetTitle(config.title)
	wnd.SetRoot(config.root)
	wnd.SetDraggable(config.draggable)
	wnd.SetResizable(config.resizeable)
	wnd.SetModal(true)
	wnd.SetBackgroundColor(config.backgroundColor)
	wnd.SetBorder(config.border)

	wnd.SetRect(config.size.x, config.size.y, config.size.width, config.size.height)
	wnd.AddButton(&winman.Button{
		Symbol: 'X',
		OnClick: func() {
			e.closeModal(wnd, config.fallback)
		},
	})

	e.winman.AddWindow(wnd)
	e.winman.Center(wnd)
	e.setFocus(wnd)

	return wnd

}

func (e *Engine) closeModal(wnd *winman.WindowBase, focus tview.Primitive) {
	e.winman.RemoveWindow(wnd)
	e.setFocus(focus)
}
