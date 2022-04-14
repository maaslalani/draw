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
	// character is the current character selected to be written to the grid.
	character string
	// color is the current color selected that will be used to write the
	// `character` to the grid.
	color string
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
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseLeft:
			// When the user clicks on the mouse, we want to write the
			// character to the current position of the mouse in the grid, so
			// that we can draw it later.
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.color))
			m.canvas[msg.Y][msg.X] = style.Render(m.character)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
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
