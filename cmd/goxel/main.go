package main

import (
	"github.com/chengxuncc/goxel"
	"io"
	"log"
)

func main() {
	log.SetOutput(io.Discard)

	goxel.NewGoXel().Run()
}
