package main

import (
	"regexp"
	"syscall/js"

	"github.com/life4/gweb/web"
)

var alpha = regexp.MustCompile(`[a-zA-Z]+`)

type AutoComplete struct {
	editor web.Value
}

func (a AutoComplete) Register() {
	a.editor.Call("on", "inputRead", js.FuncOf(a.callback))
}

func (a AutoComplete) callback(this js.Value, args []js.Value) any {
	text := args[1].Get("text").Index(0).String()
	if alpha.MatchString(text) {
		a.editor.Call("showHint")
	}
	return nil
}
