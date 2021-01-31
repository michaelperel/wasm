package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	value1, _ := strconv.Atoi(i[0].String())
	value2, _ := strconv.Atoi(i[1].String())
	return value1 + value2
}

func subtract(this js.Value, i []js.Value) interface{} {
	value1, _ := strconv.Atoi(i[0].String())
	value2, _ := strconv.Atoi(i[1].String())
	return value1 - value2
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
