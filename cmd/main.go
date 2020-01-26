package main

import (
	"fmt"

	"github.com/darkraiden/odysseus/internal/logs"
	"github.com/darkraiden/odysseus/internal/template"
	"github.com/darkraiden/odysseus/internal/whatsmyip"
)

func main() {
	// wire our template and logger together
	logger := logs.NewLogger()

	l := template.New(logger, "Info")

	l.Log("Welcome to Odysseus")

	ip, err := whatsmyip.GetLocalIp()
	if err != nil {
		panic(err)
	}

	l.Log(fmt.Sprintf("Your local IP Address is: %s", *ip))
}
