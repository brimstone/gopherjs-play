// +build !generate

//go:generate go run assets.go generate.go assets_generate.go client_generate.go

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/shurcooL/httpgzip"
	"golang.org/x/net/websocket"
)

func readWebsocket(ws *websocket.Conn, outChan chan string, errChan chan error) {
	var message string
	var err error
	for {
		// read in a message
		err = websocket.Message.Receive(ws, &message)
		if err != nil {
			errChan <- err
			return
		}
		// send it out of our channel
		outChan <- message
	}
}
func websocketHandler(ws *websocket.Conn) {
	// when we exit the function, close the socket
	defer ws.Close()
	// start our reader in the background
	clientmsg := make(chan string)
	clienterr := make(chan error)
	//websocket.Message.Send(ws, "Howdy ho!")
	go readWebsocket(ws, clientmsg, clienterr)
	for {
		select {
		case data := <-clientmsg:
			log.Println("Received:", data)
		case e := <-clienterr:
			if e.Error() != "EOF" {
				log.Println("Error from websocket: ", e.Error())
			}
			return

		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(websocketHandler))
	http.Handle("/assets/", httpgzip.FileServer(assets, httpgzip.FileServerOptions{ServeError: httpgzip.Detailed}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		f, _ := assets.Open("/assets/index.html")
		io.Copy(w, f)

	})

	fmt.Println("ready")
	http.ListenAndServe(":8080", nil)
}
