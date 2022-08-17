package main

import (
	"net/http"
	"syscall/js"
)

func getInput() {
	doc := js.Global().Get("document")

}

//
//// function definition
//func add(this js.Value, i []js.Value) interface{} {
//	return js.ValueOf(i[0].Int() + i[1].Int())
//}
//
//func registerCallbacks() {
//	js.Global().Set("add", js.FuncOf(add))
//
//}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./assets")))
	if err != nil {
		return
	}

}
