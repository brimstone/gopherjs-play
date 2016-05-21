// +build !generate

//go:generate go run assets.go generate.go assets_generate.go client_generate.go

package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/shurcooL/go/gzip_file_server"
)

func main() {
	http.Handle("/assets/", gzip_file_server.New(assets))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		f, _ := assets.Open("/assets/index.html")
		io.Copy(w, f)

	})

	fmt.Println("ready")
	http.ListenAndServe(":8080", nil)
}
