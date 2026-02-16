package commands

import (
	"fmt"
	"os"
)

type Commands func(args []string)

func RegisterCommands() map[string]Commands {
	return map[string]Commands{
		"exit": func(args []string) { os.Exit(1) },
		"version": func(args []string) {
			fmt.Println("star-vpn v0.1.0")
		},
	}
}
