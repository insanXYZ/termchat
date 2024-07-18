package engine

import (
	"bin-term-chat/handler"
	"bin-term-chat/model"
	"github.com/epiclabs-io/winman"
	"github.com/gdamore/tcell/v2"
	"github.com/gorilla/websocket"
	"github.com/rivo/tview"
	"strings"
)

var StyleTextView = tcell.StyleDefault.Italic(true)

type Engine struct {
	app      *tview.Application
	pages    *tview.Pages
	winman   *winman.Manager
	handler  *handler.Handler
	conn     *websocket.Conn
	user     *model.User
	compHub  map[string]model.CompHub
	url      string
	receiver string
	token    string
}

func NewEngine(url string) *Engine {
	app := tview.NewApplication()
	pages := tview.NewPages()
	wm := winman.NewWindowManager()
	h := handler.NewHandler(url)

	engine := &Engine{
		app:     app,
		pages:   pages,
		winman:  wm,
		handler: h,
		url:     url,
		conn:    nil,
		user:    nil,
		compHub: make(map[string]model.CompHub),
	}

	engine.initRoute()

	return engine
}

func (e *Engine) Run() error {
	return e.app.SetRoot(e.pages, true).SetFocus(e.pages).EnableMouse(true).Run()
}

func (e *Engine) initRoute() {
	e.addPage("login", e.login())
	e.addPage("register", e.register())

}

func (e *Engine) addPage(name string, comp tview.Primitive, visible ...bool) {
	if len(visible) == 0 {
		visible = append(visible, true)
	}
	e.pages.AddPage(name, comp, true, visible[0])
}

func (e *Engine) setFocus(c tview.Primitive) {
	go e.app.QueueUpdateDraw(func() {
		e.app.SetFocus(c)
	})
}

func (e *Engine) setAuthEngine(data map[string]any) {
	e.token = data["token"].(string)
	e.user = &model.User{
		Name:  data["name"].(string),
		Email: data["email"].(string),
		ID:    data["id"].(string),
	}
}

func (e *Engine) setHub(index string, value model.CompHub) {
	e.compHub[index] = value
}

func (e *Engine) setRoot(c tview.Primitive) {
	go e.app.QueueUpdateDraw(func() {
		e.app.SetRoot(c, true)
	})

	e.setFocus(c)
}

func (e *Engine) connectWebsocket() error {
	url := strings.Split(e.url, "//")[1]

	if strings.HasPrefix(e.url, "http://") {
		url = "ws://" + url
	} else {
		url = "wss://" + url
	}

	conn, _, err := websocket.DefaultDialer.Dial(url+"/api/chat?token="+e.token, nil)
	if err != nil {
		return err
	}

	e.conn = conn

	return nil
}
