package main

import (
	"fmt"
	"os"

	"github.com/Pauloo27/gotroller/cli/polybar"
)

type Mode func()

var modes = map[string]Mode{
	"polybar-dmenu": polybar.WithDmenu,
	"polybar-gui":   polybar.WithGUI,
}

func listModes() {
	fmt.Print("Valid modes: ")

	var prefix string
	for modeName := range modes {
		fmt.Printf("%s%s", prefix, modeName)
		if prefix == "" {
			prefix = ", "
		}
	}

	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing mode.")
		listModes()
		os.Exit(-1)
	}

	mode, ok := modes[os.Args[1]]
	if !ok {
		fmt.Println("Invalid mode.")
		listModes()
		os.Exit(-1)
	}

	mode()
}
