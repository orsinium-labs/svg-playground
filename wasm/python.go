package main

import (
	"fmt"
	"syscall/js"

	"github.com/life4/gweb/web"
)

type Python struct {
	pyodide web.Value
	doc     web.Document
	output  web.HTMLElement
}

func (py *Python) Register() {
	wrapped := func(this js.Value, args []js.Value) interface{} {
		cmd := py.doc.Element("py-input-text").Get("value").String()
		py.RunAndPrint(cmd)
		return true
	}
	py.doc.Element("py-input-form").Call("addEventListener", "submit", js.FuncOf(wrapped))
}

func (py Python) print(text string, cls string) {
	el := py.doc.CreateElement("pre")
	el.Attribute("class").Set("alert alert-" + cls)
	el.SetText(text)
	py.output.Node().AppendChild(el.Node())
}

func (py Python) PrintIn(text string) {
	py.print(text, "secondary")
}

func (py Python) PrintOut(text string) {
	py.print(text, "success")
}

func (py Python) PrintErr(text string) {
	py.print(text, "danger")
}

func (py Python) Run(cmd string) string {
	result := py.pyodide.Call("runPython", cmd)
	if result.Type() == js.TypeObject {
		return result.Call("toString").String()
	}
	return result.String()
}

func (py Python) RunAndPrint(cmd string) (ok bool) {
	ok = true
	defer func() {
		err := recover()
		if err != nil {
			py.PrintErr(fmt.Sprintf("%v", err))
			ok = false
		}
	}()
	py.PrintIn(cmd)
	result := py.Run(cmd)
	if result != "<undefined>" {
		py.PrintOut(result)
	}
	return ok
}

func (py Python) Install(pkg string) bool {
	cmd := fmt.Sprintf("micropip.install('%s', deps=False)", pkg)
	py.PrintIn(cmd)
	_, fail := py.pyodide.Call("runPython", cmd).Promise().Get()
	if fail.Truthy() {
		py.PrintErr(fail.String())
		return false
	}
	py.PrintOut(fmt.Sprint(pkg, " installed"))
	return true
}

func (py Python) Set(name string, text string) {
	py.pyodide.Get("globals").Call("set", name, text)
}

func (py Python) Clear() {
	py.output.SetText("")
}

func (py Python) InitMicroPip() bool {
	py.PrintIn("import micropip")
	_, fail := py.pyodide.Call("loadPackage", "micropip").Promise().Get()
	if fail.Truthy() {
		py.PrintErr(fail.String())
		return false
	}
	py.Run("import micropip")
	py.PrintOut("True")
	return true
}
