package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	generateObject()
	select{}
}

func generateObject() {
	var generateObject js.Func
	generateObject = js.FuncOf(func(this js.Value, args[]js.Value) interface{} {
		goObj := map[string]interface{}{}
		jsObj := js.ValueOf(goObj)
		jsObj.Set("do", js.FuncOf(func(this js.Value, args[]js.Value) interface{} {
			onMessage := jsObj.Get("onMessage") 
			if onMessage != js.Undefined() {
				fmt.Println("We have on message")
				onMessage.Invoke(2)
			}
			fmt.Println("After on message check")
			return 2
		}))
		generateObject.Release()
		js.Global().Set("generateObject", js.Undefined())
		return jsObj
	})
	js.Global().Set("generateObject", generateObject)
}