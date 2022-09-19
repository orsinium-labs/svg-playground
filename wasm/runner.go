package main

import (
	"syscall/js"

	"github.com/life4/gweb/web"
)

type Runner struct {
	btn    web.HTMLElement
	canvas web.HTMLElement
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
		btn:    doc.Element("py-lint"),
		canvas: doc.Element("py-canvas"),
		doc:    doc,
		win:    win,
		editor: editor,
		py:     py,
	}

}

func (fh *Runner) Register() {
	fh.btn.SetInnerHTML("draw")
	fh.btn.Set("disabled", false)

	wrapped := func(this js.Value, args []js.Value) interface{} {
		fh.btn.SetInnerHTML("running...")
		fh.btn.Set("disabled", true)
		fh.Run()
		go fh.Register()
		return true
	}
	fh.btn.Call("addEventListener", "click", js.FuncOf(wrapped))
}

func (fh *Runner) Run() {
	fh.py.Clear()
	script := fh.editor.Call("getValue").String()
	ok := fh.py.RunAndPrint(script)
	if !ok {
		return
	}

	fh.py.Clear()
	svg := fh.py.Run("str(canvas)")
	fh.canvas.SetInnerHTML(svg)
}
