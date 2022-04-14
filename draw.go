package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	// canvas contains the cells of the terminal and stores the strings +
	// styles of the characters written to the grid.
	canvas [][]string
	// backupCanvas is a backup of the canvas that we can use to undo the
	// last action.
	backupCanvas [][]string
	// character is the current character selected to be written to the grid.
	character string
	// color is the current color selected that will be used to write the
	// `character` to the grid.
	color string
	// cursor is used to anchor the cursor when a user selects a position
	// to keep in mind. I.e. when a user is drawing a box
	cursor struct{ x, y int }
}

// Init initializes the model with the initial state.
func (m *model) Init() tea.Cmd {
	m.character = "*"
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// On window resize, we want to create a replica grid
		// (called the canvas) for the user to draw on.
		//
		// This will be a 2D slice of strings. We use strings and not runes so
		// that we can store the style of the character drawn as well so that
		// each cell can be a different style / color.
		m.canvas = make([][]string, msg.Height)
		for i := range m.canvas {
			m.canvas[i] = make([]string, msg.Width)
		}
		m.backupCanvas = make([][]string, msg.Height)
		for i := range m.backupCanvas {
			m.backupCanvas[i] = make([]string, msg.Width)
		}
		m.backup()
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseRelease:
			m.cursorReset()
			return m, nil
		case tea.MouseLeft:
			// When the user clicks on the mouse, we want to write the
			// character to the current position of the mouse in the grid, so
			// that we can draw it later.
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.color))
			m.canvas[msg.Y][msg.X] = style.Render(m.character)
		case tea.MouseRight:
			if m.cursorIsSet() {
				// Cursor was already set, now draw the box based on the
				// cursor and current position of the mouse.
				m.restore()
				m.DrawShape(Point{m.cursor.x, m.cursor.y}, Point{msg.X, msg.Y})
			} else {
				// Cursor was not set, so set it to the current position
				// of the mouse and create a backup of the canvas.
				m.backup()
				m.cursor.x = msg.X
				m.cursor.y = msg.Y
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.cursorReset()
			m.restore()
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// Use the number keys to select the color.
			// The colors are: red, green, yellow, blue, magenta, cyan, white.
			// which correspond to the colors in the terminal / lipgloss colors.
			m.color = msg.String()
		default:
			// Otherwise, we will want to change the character that is being
			// used to the character that the user just typed.
			m.character = msg.String()
		}
	}
	return m, nil
}

// View renders the grid to the screen. Since we write the character & style to
// the grid, we can simply print out the grid.
func (m model) View() string {
	var s strings.Builder
	for _, row := range m.canvas {
		for _, cell := range row {
			if cell == "" {
				s.WriteString(" ")
			} else {
				s.WriteString(cell)
			}
		}
		s.WriteString("\n")
	}
	return strings.TrimSuffix(s.String(), "\n")
}

func (m *model) cursorIsSet() bool {
	return m.cursor.x != 0 && m.cursor.y != 0
}

func (m *model) cursorReset() {
	m.cursor.x = 0
	m.cursor.y = 0
}

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
