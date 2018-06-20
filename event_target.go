// +build js,wasm

package webapi

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget struct {
	v js.Value
}

func (r *EventTarget) AddEventListener(typ string, listener func(args []js.Value)) {
	r.v.Call("addEventListener", typ, js.NewCallback(listener))
}

func (r *EventTarget) RemoveEventListener() {
	panic("RemoveEventListener")
}

func (r *EventTarget) DispatchEvent() {
	panic("DispatchEvent")
}
