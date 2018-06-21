// +build js,wasm

package webapi

import "syscall/js"

type Location interface {
	Reload()
}

type implLocation struct {
	v js.Value
}

func (r *implLocation) Reload() {
	r.v.Call("reload")
}
