package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	antim "github.com/pavelpatrin/go-ansi-to-image"
)

const DEFAULT_DRAW_NAME = "/tmp/draw"
const DEFAULT_DRAW_EXT = ".txt"
const DEFAULT_DRAW_FILE = DEFAULT_DRAW_NAME + DEFAULT_DRAW_EXT
const DRAW_FILE_KEY = "DRAW_FILE"
const DRAW_NAME_KEY = "DRAW_NAME"
const DRAW_EXT_KEY = "DRAW_EXT"

func main() {
	var helpFlag = flag.Bool("help", false, "help")
	flag.Parse()

	os.Args = append(os.Args, "") // so that the next line does never panics
	file := filePath(os.Args[1])

	if *helpFlag {
		fmt.Fprintln(os.Stderr, help())
		os.Exit(0)
	}

	m := &model{}

	p := tea.NewProgram(m)

	p.EnableMouseAllMotion()
	p.EnterAltScreen()

	must(p.Start())

	// Save the user's beautiful drawing, in the event they want to
	// appreciate it later which they can do with `cat`.
	//
	// cat $DRAW_FILE || cat "$DRAW_NAME$DRAW_EXT" || cat /tmp/draw.txt

	switch ext := strings.ToLower(filepath.Ext(file)); ext {
	default:
		fmt.Fprintf(os.Stderr, "this application is not equipped to produce %q files. Rendering an eponymous text file instead\n", ext)
		fallthrough

	case ".txt":
		must(os.WriteFile(file, []byte(m.View()), 0644))

	case ".png":
		cfg := antim.DefaultConfig
		cfg.PageRows = len(m.canvas)    // windowSize.Height
		cfg.PageCols = len(m.canvas[0]) // windowSize.Width

		c, err := antim.NewConverter(cfg)
		must(err)

		must(c.Parse(m.View()))

		data, err := c.ToPNG()
		must(err)

		must(os.WriteFile(file, data, 0644))
	}
}

// arg if the arg is non-empty
// name + ext if both keys are set
// file if key is set
// default file otherwise
func filePath(file string) string {
	if file != "" {
		return file
	}

	name, ok := os.LookupEnv(DRAW_NAME_KEY)
	if ok {
		ext, ok := os.LookupEnv(DRAW_EXT_KEY)
		if ok {
			return name + ext
		}
	}

	file, ok = os.LookupEnv(DRAW_FILE_KEY)
	if ok {
		return file
	}

	return DEFAULT_DRAW_FILE
}

// exit on error
func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
