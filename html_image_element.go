// +build js,wasm

package web

import "syscall/js"

type HTMLImageElement interface {
	HTMLElement
}

type implHTMLImageElement struct {
	implHTMLElement
	v js.Value
}

func newHTMLImageElement(v js.Value) HTMLImageElement {
	return &implHTMLImageElement{
		newImplHTMLElement(v),
		v,
	}
}
