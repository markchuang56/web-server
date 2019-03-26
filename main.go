package main

import (
	"fmt"
	"os"

	//"../../pkg/cmd/server"
	"go-heroku/fitbit-demo/web-server/pkg/cmd/server"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
