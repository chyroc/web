// +build js,wasm

package webapi

import "syscall/js"

var Document *document
var doc js.Value

func init() {
	Document = new(document)
	doc = js.Global.Get("document")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document
type document struct {
	Node
	EventTarget
}

func (r *document) GetElementById(id string) *Element {
	ele := doc.Call("getElementById", id)
	eventTarget := EventTarget{v: ele}
	return &Element{
		EventTarget: eventTarget,
		v:           ele,
	}
}
