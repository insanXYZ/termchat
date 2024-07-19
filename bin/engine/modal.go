package engine

import "github.com/rivo/tview"

type Modal struct {
	root     tview.Primitive
	title    string
	fallback tview.Primitive
}

func (e *Engine) CreateModal() {

}
