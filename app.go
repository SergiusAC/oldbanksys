package main

import (
	"fmt"
	"oldbanksys/commands"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		commands.ShowHelp()
		return
	}

	err := commands.Dispatch(os.Args[1], os.Args[2:])

	if err != nil {
		fmt.Println(err)
	}
}
