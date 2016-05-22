package main

import (
	"log"

	"github.com/gopherjs/websocket"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

const (
	// INPUT holds our user's name
	INPUT = "input#name"
	// OUTPUT holds greeting
	OUTPUT = "span#greeting"
)

var jQuery = jquery.NewJQuery

func websocketAddress() string {
	location := js.Global.Get("window").Get("location")
	protocol := location.Get("protocol").String()
	host := location.Get("host").String()

	if protocol == "http:" {
		protocol = "ws:"
	} else {
		protocol = "wss:"
	}

	return protocol + "//" + host + "/ws"
}

func websocketSetup() error {
	c, err := websocket.Dial(websocketAddress())
	if err != nil {
		return err
	}
	log.Println("Websocket open")

	jQuery(INPUT).On(jquery.KEYUP, func(e jquery.Event) {
		name := jQuery(e.Target).Val()
		name = jquery.Trim(name)
		log.Println("Writing to the server")
		_, err = c.Write([]byte("User entered:" + name))
	})

	return nil
}

func main() {

	log.Println("Hello world")

	go websocketSetup()

	jQuery(INPUT).On(jquery.KEYUP, func(e jquery.Event) {

		name := jQuery(e.Target).Val()
		name = jquery.Trim(name)

		//show welcome message:
		if len(name) > 0 {
			jQuery(OUTPUT).SetText("Welcome to GopherJS, " + name + "!")
		} else {
			jQuery(OUTPUT).Empty()
		}
	})

	/*
	   c, err := websocket.Dial("ws://localhost/socket") // Blocks until connection is established
	   if err != nil { panic(err) }

	   buf := make([]byte, 1024)
	   n, err = c.Read(buf) // Blocks until a WebSocket frame is received
	   if err != nil { panic(err) }
	   doSomethingWithData(buf[:n])

	   _, err = c.Write([]byte("Hello!"))
	   if err != nil { panic(err) }

	   err = c.Close()
	   if err != nil { panic(err) }
	*/
}
