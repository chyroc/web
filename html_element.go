// +build js,wasm

package webapi

type HTMLElement interface {
	Element
}

type implHTMLElement struct {
	implElement
}