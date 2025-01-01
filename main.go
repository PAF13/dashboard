package main

import (
	"log"

	"github.com/PAF13/dashboard/internal/httpserver"
	"github.com/PAF13/dashboard/internal/tui"
)

func main() {
	log.Println("starting")
	logBuf := &httpserver.LogBuffer{}
	log.SetOutput(logBuf) // Redirect log output to the buffer.
	go httpserver.Server(logBuf)

	if true {
		tui.TempFunc()
	}
}
