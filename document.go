// +build js,wasm

package web

import (
	"strings"
	"syscall/js"
)

var Document *document

func init() {
	v := js.Global.Get("document")
	Document = &document{
		implEventTarget: newImplEventTarget(v),
		v:               v,
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document
type document struct {
	implEventTarget

	v js.Value
}

func (r *document) Location() Location {
	t := newImplLocation(r.v.Get("location"))
	return &t
}

func (r *document) GetElementById(id string) HTMLElement {
	v := r.v.Call("getElementById", id)

	switch tagName := strings.ToUpper(v.Get("tagName").String()); tagName {
	case "CANVAS":
		return newHTMLCanvasElement(v)
	default:
		t := newImplHTMLElement(v)
		return &t
	}
}
