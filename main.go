package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, WebAssembly!")

	js.Global().Set("formatJSON", js.FuncOf(jsFormatJSON))

	<-c
}

func jsAlert() {
	jsWindow := js.Global().Get("window")
	jsWindow.Call("alert", "Something went wrong!")
}

func jsFormatJSON(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		jsAlert()
		return "Invalid number of arguments passed"
	}

	input := args[0].String()

	output, err := formatJSON(input)
	if err != nil {
		jsAlert()
		return err.Error()
	}

	return output
}

func formatJSON(input string) (string, error) {
	var raw interface{}

	err := json.Unmarshal([]byte(input), &raw)
	if err != nil {
		return "", err
	}

	fmt.Printf("raw = %+v\n", raw)

	indent := "  "
	output, err := json.MarshalIndent(raw, "", indent)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
