package main

func help() string {
	return `Draw in your terminal.

Usage: draw [file]

Key bindings:
  ctrl+c:       Quit
  ctrl+z:       Undo last stroke
  any key:      Draw with that character
  1-9:          Change brush color, ANSI colors

Mouse:
  Left click:   Draw
  Right click:  Set text anchor
  Right drag:   Draw box

Environment variables:
  DRAW_FILE:    The file to save the drawing to.
                Default: /tmp/draw.txt

Arguments:
  [file]:       The file to save the drawing to.
                Overrides $DRAW_FILE.
`
}
