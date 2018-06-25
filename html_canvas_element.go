// +build js,wasm

package web

import (
	"fmt"
	"syscall/js"
)

type HTMLCanvasElement interface {
	HTMLElement

	GetContext(string) CanvasRenderingContext2D
	Width() int
	Height() int
}

type implHTMLCanvasElement struct {
	implHTMLElement
}

func newHTMLCanvasElement(v js.Value) HTMLCanvasElement {
	return &implHTMLCanvasElement{
		implHTMLElement: newImplHTMLElement(v),
	}
}

func (r *implHTMLCanvasElement) GetContext(contextID string) CanvasRenderingContext2D {
	t := newCanvasRenderingContext2D(r.v, r.v.Call("getContext", contextID))

	switch contextID {
	case "2d":
		return &t
	default:
		panic(fmt.Sprintf("unsupport %s contextId", contextID))
	}
}

func (r *implHTMLCanvasElement) Width() int {
	return r.v.Get("width").Int()
}

func (r *implHTMLCanvasElement) Height() int {
	return r.v.Get("height").Int()
}
