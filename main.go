package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const DEFAULT_DRAW_FILE = "/tmp/draw.txt"
const DRAW_FILE_KEY = "DRAW_FILE"

func main() {
	m := &model{}

	p := tea.NewProgram(m)

	p.EnableMouseAllMotion()
	p.EnterAltScreen()

	if err := p.Start(); err != nil {
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
	os.WriteFile(file, []byte(m.View()), 0644)
}
