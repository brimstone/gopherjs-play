// +build generate

package main

import (
	"fmt"
	"log"

	"github.com/shurcooL/vfsgen"
)

func assetsMain() {

	fmt.Println("Slurping up assets")
	err := vfsgen.Generate(assets, vfsgen.Options{
		BuildTags: "!dev",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
