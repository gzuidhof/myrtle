package main

import (
	"log"
	"net"
	"os"
	"strings"

	"github.com/gzuidhof/myrtle/example/server"
)

func main() {
	address := os.Getenv("MYRTLE_SERVER_ADDR")
	srv, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("myrtle example server listening on http://localhost%s", localPort(address))
	log.Fatal(srv.ListenAndServe(address))
}

func localPort(address string) string {
	trimmed := strings.TrimSpace(address)
	if trimmed == "" {
		return ":8380"
	}

	if strings.HasPrefix(trimmed, ":") {
		return trimmed
	}

	_, port, err := net.SplitHostPort(trimmed)
	if err == nil {
		return ":" + port
	}

	return ":8380"
}
