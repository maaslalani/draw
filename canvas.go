package main

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
