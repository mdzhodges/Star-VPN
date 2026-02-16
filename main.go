package main

import (
	"bufio"
	"fmt"
	"os"
	"star-vpn/commands"
	"strings"
)

func main() {
	//Register Commands
	cmds := commands.RegisterCommands()
	//Init Scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//Grab arguments
		args := strings.Fields(scanner.Text())
		if len(args) == 0 {
			continue
		}
		//Find Command
		cmd, ok := cmds[args[0]]
		if !ok {
			fmt.Printf("Invalid command entry %s\n", args[0])
			continue
		}
		//Execute Command
		cmd(args[1:])

	}
}
