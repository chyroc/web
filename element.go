// +build js,wasm

package webapi

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/element
type Element struct {
	EventTarget

	v js.Value
}

func (r *Element) SetInnerHTML(s string) {
	r.v.Set("innerHTML", s)
}

func (r *Element) SetAttribute(name string, value interface{}) {
	r.v.Set(name, value)
}
