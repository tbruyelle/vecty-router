package router

import (
	"syscall/js"
)

func init() {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		refreshRoutes()
		return nil
	})
	js.Global().Set("onpopstate", f)
}
