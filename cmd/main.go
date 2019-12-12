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
	// the log.Println goes to console log
	log.Println("Module loaded")
	fmt.Printf("and can write using %s\n", "fmt")
}

func main() {
	// register/export functions to JavaScript Global context
	js.Global().Set("goFunction", js.FuncOf(fromJsToGo))

	buttonClicked := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		output.Set("innerHTML", "from bound event")
		return nil
			//cb.Release() // release the function if the button will not be clicked again
		return nil
	})
	js.Global().Get("document").Call("querySelector", "#button-2").Call("addEventListener", "click", buttonClicked)

	// Call a function in the main that will interact with the DOM
	fromGoToDOM()

	// As this is an all, it run and end. To keep it alive (don't exit)
	// we create a channel and wait foerver
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
