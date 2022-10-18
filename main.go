package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const DEFAULT_DRAW_FILE = "/tmp/draw.txt"
const DRAW_FILE_KEY = "DRAW_FILE"

func main() {
	var helpFlag = flag.Bool("help", false, "help")
	flag.Parse()

	if *helpFlag {
		fmt.Println(help())
		os.Exit(0)
	}

	m := &model{}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())

	_, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// Save the user's beautiful drawing, in the event they want to
	// appreciate it later which they can do with `cat`.
	//
	// cat $DRAW_FILE || cat /tmp/draw.txt
	//
	file, ok := os.LookupEnv(DRAW_FILE_KEY)
	if !ok {
		file = DEFAULT_DRAW_FILE
	}
	// If a file is given as an argument, use that instead.
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	os.WriteFile(file, []byte(m.View()), 0644)
}
