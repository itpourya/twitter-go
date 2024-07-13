package main

import (
	"fmt"
	"twitter-go-api/internal/server"
)

func main() {

	srv := server.NewServer()

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start srv: %s", err))
	}
}
