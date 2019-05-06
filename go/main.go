package main

import (
	"fmt"
	"syscall/js"
)

func main() {
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
		return jsObj
	})
	js.Global().Set("generateObject", generateObject)
	select{}
}