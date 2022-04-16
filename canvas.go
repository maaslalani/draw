package main

import "strings"

// backup is a helper function that copies the current canvas to the backup
// canvas.
func (m *model) backup() {
	for i := range m.canvas {
		for j := range m.canvas[i] {
			m.backupCanvas[i][j] = m.canvas[i][j]
		}
	}
}

// restore is a helper function that copies the backup canvas to the current
// canvas.
func (m *model) restore() {
	for i := range m.canvas {
		for j := range m.canvas[i] {
			m.canvas[i][j] = m.backupCanvas[i][j]
		}
	}
}

// load is a helper function that loads contents into the canvas.
func (m *model) load(contents string) {
	lines := strings.Split(contents, "\n")
	for i, line := range lines {
		if i >= len(m.canvas) {
			break
		}
		for j, c := range line {
			if j >= len(m.canvas[i]) {
				break
			}
			m.canvas[i][j] = string(c)
		}
	}
}
