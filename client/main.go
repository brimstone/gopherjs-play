package main

import (
	"log"

	"github.com/gopherjs/jquery"
)

const (
	// INPUT holds our user's name
	INPUT = "input#name"
	// OUTPUT holds greeting
	OUTPUT = "span#output"
)

var jQuery = jquery.NewJQuery

func main() {
	log.Println("Hello world")
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

}
