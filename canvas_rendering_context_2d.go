// +build js,wasm

package webapi

import "syscall/js"

type CanvasRenderingContext2D interface {
	// fillStyle
	FillStyle() string
	SetFillStyle(string)
	// font
	SetFont(string)

	BeginPath()
	Rect(x, y float64, width, height int)
	ClearRect(x, y float64, width, height int)
	Fill()
	ClosePath()
	Arc(x, y, radius, startAngle, endAngle float64)
	MoveTo(x, y float64)
	LineTo(x, y float64)
	FillText(s string, x, y float64)
}

type implCanvasRenderingContext2D struct {
	v js.Value
}

func (r *implCanvasRenderingContext2D) FillStyle() string {
	return r.v.Get("fillStyle").String()
}

func (r *implCanvasRenderingContext2D) SetFillStyle(s string) {
	r.v.Set("fillStyle", s)
}

func (r *implCanvasRenderingContext2D) SetFont(s string) {
	r.v.Set("front", s)
}

func (r *implCanvasRenderingContext2D) BeginPath() {
	r.v.Call("beginPath")
}

func (r *implCanvasRenderingContext2D) Rect(x, y float64, width, height int) {
	r.v.Call("rect", x, y, width, height)
}

func (r *implCanvasRenderingContext2D) ClearRect(x, y float64, width, height int) {
	r.v.Call("clearRect", x, y, width, height)
}

func (r *implCanvasRenderingContext2D) Fill() {
	r.v.Call("fill")
}

func (r *implCanvasRenderingContext2D) ClosePath() {
	r.v.Call("closePath")
}

func (r *implCanvasRenderingContext2D) Arc(x, y, radius, startAngle, endAngle float64) {
	r.v.Call("arc", x, y, radius, startAngle, endAngle)
}

func (r *implCanvasRenderingContext2D) MoveTo(x, y float64) {
	r.v.Call("moveTo", x, y)
}

func (r *implCanvasRenderingContext2D) LineTo(x, y float64) {
	r.v.Call("lineTo", x, y)
}

func (r *implCanvasRenderingContext2D) FillText(s string, x, y float64) {
	r.v.Call("fillText", s, x, y)
}
