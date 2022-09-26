package main

import (
	"regexp"
	"strings"
	"syscall/js"

	"github.com/life4/gweb/web"
)

var alpha = regexp.MustCompile(`[a-zA-Z]+`)

type AutoComplete struct {
	editor web.Value
	py     Python
	window web.Window
}

func (a *AutoComplete) Register() {
	a.py.Run("import builtins")
	builtins := strings.Split(a.py.Run("' '.join(dir(builtins))"), " ")

	globals := make([]any, len(builtins))
	for i, v := range builtins {
		globals[i] = v
	}

	a.window.Get("CodeMirror").Call("registerHelper", "hintWords", "python", globals)
	a.editor.Call("on", "inputRead", js.FuncOf(a.callback))
}

func (a AutoComplete) callback(this js.Value, args []js.Value) any {
	text := args[1].Get("text").Index(0).String()
	if alpha.MatchString(text) {
		a.editor.Call("showHint")
	}
	return nil
}
