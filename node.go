// +build js,wasm

package webapi

// https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node interface {
	BaseURI() string
	BaseURIObject() string
}
