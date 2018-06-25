// +build js,wasm

package web

import "syscall/js"

type HTMLSelectElement interface {
	HTMLElement

	Value() string
}

func newImplHTMLSelectElement(v js.Value) implHTMLSelectElement {
	return implHTMLSelectElement{newImplHTMLElement(v), v}
}

func newHTMLSelectElement(v js.Value) HTMLSelectElement {
	d := newImplHTMLSelectElement(v)
	return &d
}

type implHTMLSelectElement struct {
	implHTMLElement

	v js.Value
}

func (r *implHTMLSelectElement) Value() string {
	return r.v.Get("value").String()
}
