package main

import (
	"syscall/js"
)

func main() {
	var generateObject js.Func
	generateObject = js.FuncOf(func(this js.Value, args[]js.Value) interface{} {
		chatObj := map[string]interface{}{}
		jsObj := js.ValueOf(chatObj)
		jsObj.Set("do", js.FuncOf(func(this js.Value, args[]js.Value) interface{} {
			return 2
		}))
		generateObject.Release()
		return jsObj
	})
	
	js.Global().Set("generateObject", generateObject)

	select{}
}