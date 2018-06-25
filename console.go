// +build js,wasm

package web

import (
	"fmt"
	"syscall/js"
)

var Console *console

func init() {
	Console = &console{
		v: js.Global.Get("console"),
	}
}

type console struct {
	v js.Value
}

func (r *console) Log(msg string, a ...interface{}) {
	r.v.Call("log", fmt.Sprintf(msg+"\n", a...))
}
