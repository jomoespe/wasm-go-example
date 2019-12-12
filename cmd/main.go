// +build js,wasm

package main

import (
	"fmt"
	"log"
	"syscall/js"
	"time"
)

const (
	wait = 3 * time.Second
)

var (
	output = js.Global().Get("document").Call("querySelector", "#output")
)

func init() {
	// the log.Println and the fmt.Print goes to console log
	log.Println("Module loaded")
	fmt.Printf("and can write using %s\n", "fmt")
}

func main() {
	// register/export functions to JavaScript Global context
	js.Global().Set("goFunction", js.FuncOf(fromJsToGo))

	// create an element (an H2) and append to docuemnt body
	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerHTML", "this element have been created from WASM")
	js.Global().Get("document").Get("body").Call("appendChild", h2)

	buttonClicked := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		output.Set("innerHTML", "from event bound in WASM")
		//buttonClicked.Release() // release the function if the button will not be clicked again
		return nil
	})
	js.Global().Get("document").Call("querySelector", "#button-2").Call("addEventListener", "click", buttonClicked)

	// Call a function in the main that will interact with the DOM
	fromGoToDOM()

	// To keep it alive (don't exit) we create a channel and wait foerver
	c := make(chan bool)
	<-c
}

func fromGoToDOM() {
	output.Set("innerHTML", fmt.Sprintf("Wait %s seconds....", wait))

	// Let's create a gorutine that will sleep 3 seconds
	c := make(chan string)
	go func() {
		time.Sleep(wait)
		c <- "from Gorutine!"
	}()
	s := <-c
	
	output.Set("innerHTML", s)
}

func fromJsToGo(this js.Value, args []js.Value) interface{} {
	message := args[0].String()
	output.Set("innerHTML", message)
	return nil
}
