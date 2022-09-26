package main

import (
	"github.com/life4/gweb/web"
)

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("svg.py playground")

	// init code editor
	input := doc.Element("py-code")
	scripts := NewScripts()
	ex := scripts.ReadExample()
	input.SetInnerHTML(ex)
	editor := window.Get("CodeMirror").Call("fromTextArea",
		input.JSValue(),
		map[string]any{
			"lineNumbers": true,
			"mode":        "python",
			"hintOptions": map[string]any{
				"completeSingle": false,
			},
		},
	)

	// load python
	py := Python{doc: doc, output: doc.Element("py-output")}
	py.PrintIn("Loading Python...")
	var err web.Value
	py.pyodide, err = window.Call("loadPyodide").Promise().Get()
	if !err.IsUndefined() {
		py.PrintErr(err.String())
		return
	}
	py.PrintOut("Python is ready")
	py.RunAndPrint("'Hello world!'")
	py.Register()

	// install dependencies
	ok := py.InitMicroPip()
	if !ok {
		return
	}
	py.Install("svg.py")

	// activate autocomplete
	ac := AutoComplete{
		editor: editor,
		window: window,
		py:     py,
	}
	ac.Register()

	runner := NewRunner(window, doc, editor, &py)
	runner.Register()
	go runner.Run()

	select {}
}
