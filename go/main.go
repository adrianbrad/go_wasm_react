package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("startChat", js.FuncOf(func(this js.Value, args[]js.Value) interface{} {
		// return js.FuncOf(func(this js.Value, args[]js.Value) interface{} {
			chatObj := map[string]interface{}{}
			jsObj := js.ValueOf(chatObj)
			jsObj.Set("do", js.FuncOf(func(this js.Value, args[]js.Value) interface{} {
				return 2
			}))
			return jsObj
		// })
	}))
	select{}
}