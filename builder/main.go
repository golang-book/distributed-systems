package main

import (
	"log"

	"github.com/gopherjs/gopherjs/build"
)

func main() {
	log.SetFlags(0)

	session := build.NewSession(&build.Options{
		GOPATH: "/Users/caleb/src/github.com/golang-book/distributed-systems/example",
	})
	archive, err := session.BuildImportPath("runner")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(archive)

	// err = http.ListenAndServe("127.0.0.1:5000", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
