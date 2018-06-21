// +build js,wasm

package webapi

type HTMLCanvasElement interface {
	HTMLElement

	GetContext(string) CanvasRenderingContext2D
	Width() int
	Height() int
}

type implHTMLCanvasElement struct {
	implHTMLElement
}

func (r *implHTMLCanvasElement) GetContext(contextID string) CanvasRenderingContext2D {
	v := r.v.Call("getContext", contextID)
	return &implCanvasRenderingContext2D{
		v: v,
	}
}

func (r *implHTMLCanvasElement) Width() int {
	return r.v.Get("width").Int()
}

func (r *implHTMLCanvasElement) Height() int {
	return r.v.Get("height").Int()
}
