package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"syscall/js"
)

func getTaylorSwiftQuote(this js.Value, args []js.Value) interface{} {
	// Get the URL as argument
	// args[0] is a js.Value, so we need to get a string out of it
	requestURL := args[0].String()

	// Handler for the Promise
	// We need to return a Promise because HTTP requests are blocking in Go
	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		// Run this code asynchronously
		go func() {
			// Make the HTTP request
			res, err := http.DefaultClient.Get(requestURL)
			if err != nil {
				// Handle errors: reject the Promise if we have an error
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New(err.Error())
				reject.Invoke(errorObject)
				return
			}
			defer res.Body.Close()

			// Read the response body
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				// Handle errors here too
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New(err.Error())
				reject.Invoke(errorObject)
				return
			}

			// "data" is a byte slice, so we need to convert it to a JS Uint8Array object
			arrayConstructor := js.Global().Get("Uint8Array")
			dataJS := arrayConstructor.New(len(data))
			js.CopyBytesToJS(dataJS, data)

			// Create a Response object and pass the data
			responseConstructor := js.Global().Get("Response")
			response := responseConstructor.New(dataJS)

			// Resolve the Promise
			resolve.Invoke(response)
		}()

		// The handler of a Promise doesn't return any value
		return nil
	})

	// Create and return the Promise object
	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

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
	js.Global().Set("getTaylorSwiftQuote", js.FuncOf(getTaylorSwiftQuote))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
