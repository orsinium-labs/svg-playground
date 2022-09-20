package main

import (
	"github.com/life4/gweb/web"
)

type Runner struct {
	doc    web.Document
	win    web.Window
	editor web.Value
	py     *Python
}

type Violation struct {
	Code        string
	Description string
	Context     string
	Line        int
	Column      int
	Plugin      string
}

func NewRunner(win web.Window, doc web.Document, editor web.Value, py *Python) Runner {
	return Runner{
		doc:    doc,
		win:    win,
		editor: editor,
		py:     py,
	}

}

func (r *Runner) Register() {
	btn := r.doc.Element("py-run-button")
	btn.SetInnerHTML("draw")
	btn.Set("disabled", false)

	handler := func(event web.Event) {
		btn.SetInnerHTML("running...")
		btn.Set("disabled", true)
		r.Run()
		btn.SetInnerHTML("draw")
		btn.Set("disabled", false)
		event.PreventDefault()
	}
	btn.EventTarget().Listen(web.EventTypeClick, handler)
}

func (r *Runner) Run() {
	r.py.Clear()
	script := r.editor.Call("getValue").String()
	ok := r.py.RunAndPrint(script)
	if !ok {
		return
	}

	r.py.Clear()
	svg := r.py.Run("str(canvas)")
	r.doc.Element("py-canvas").SetInnerHTML(svg)
}
