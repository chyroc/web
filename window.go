// +build js,wasm

package web

import "syscall/js"

var Window *window

func init() {
	Window = new(window)
	Window.v = js.Global.Get("window")
}

type window struct {
	v js.Value
}

func (r *window) RequestAnimationFrame(callback js.Callback) {
	r.v.Call("requestAnimationFrame", callback)
}

func (r *window) Alert(s string) {
	r.v.Call("alert", s)
}
