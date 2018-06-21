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

func (r *document) GetElementById(id string) HTMLElement {
	ele := doc.Call("getElementById", id)

	return &implHTMLElement{
		implElement: implElement{
			implEventTarget: implEventTarget{
				v: ele,
			},
			v: ele,
		},
	}
}
