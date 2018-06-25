// +build js,wasm

package web

import "syscall/js"

type WindowInterface interface {
	RequestAnimationFrame(callback js.Callback)
	Alert(s string)
}

type implWindow struct {
	v js.Value
}

func newimplWindow() implWindow {
	return implWindow{js.Global.Get("window")}
}

func (r *implWindow) RequestAnimationFrame(callback js.Callback) {
	r.v.Call("requestAnimationFrame", callback)
}

func (r *implWindow) Alert(s string) {
	r.v.Call("alert", s)
}
