package main

import (
	"os"
	"path/filepath"

	gbuild "github.com/gopherjs/gopherjs/build"
)

func gopherjsMain() {

	options := &gbuild.Options{
		CreateMapFile: true,
		Verbose:       true,
		Minify:        true,
	}
	s := gbuild.NewSession(options)
	//var pkgObj string
	currentDirectory, err := os.Getwd()
	sourceDirectory := filepath.Join(currentDirectory, "client")
	assetsDirectory := filepath.Join(currentDirectory, "assets")

	err = s.BuildDir(sourceDirectory,
		sourceDirectory,
		filepath.Join(assetsDirectory, "client.js"))
	if err != nil {
		panic(err)
	}
}
