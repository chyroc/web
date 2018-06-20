// +build js,wasm

package main

import (
	"syscall/js"
	"strconv"

	"github.com/Chyroc/webapi"
)

func main() {
	var span = webapi.Document.GetElementById("number")
	var plus = webapi.Document.GetElementById("plus")
	var minus = webapi.Document.GetElementById("minus")
	var number int

	minus.AddEventListener("click", func(args []js.Value) {
		println("press -")
		number--
		span.SetInnerHTML(strconv.Itoa(number))
	})

	plus.AddEventListener("click", func(args []js.Value) {
		println("press +")
		number++
		span.SetInnerHTML(strconv.Itoa(number))
	})

	plus.SetAttribute("disabled", false)
	minus.SetAttribute("disabled", false)

	select {}
}
