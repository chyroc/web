// +build js,wasm

package webapi

import (
	"strings"
	"syscall/js"
)

var Document *document

func init() {
	Document = new(document)
	v := js.Global.Get("document")
	Document.v = v
	Document.implEventTarget.v = v
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document
type document struct {
	implEventTarget

	v js.Value
}

func (r *document) Location() Location {
	return &implLocation{
		v: r.v.Get("location"),
	}
}

func (r *document) GetElementById(id string) HTMLElement {
	ele := r.v.Call("getElementById", id)

	tagName := strings.ToUpper(ele.Get("tagName").String())

	implHTMLElement := implHTMLElement{
		implElement: implElement{
			implEventTarget: implEventTarget{
				v: ele,
			},
			v: ele,
		},
	}

	switch tagName {
	case "CANVAS":
		return &implHTMLCanvasElement{
			implHTMLElement: implHTMLElement,
		}
	default:
		return &implHTMLElement
	}
}
