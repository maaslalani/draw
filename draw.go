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
	// anchor is used to anchor the anchor when a user selects a position
	// to keep in mind. I.e. when a user is drawing a box
	anchor     struct{ x, y int }
	textAnchor struct{ x, y, ix, iy int }
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
		m.backupCanvas = make([][]string, msg.Height)
		for i := range m.canvas {
			m.canvas[i] = make([]string, msg.Width)
			m.backupCanvas[i] = make([]string, msg.Width)
		}
		m.backup()
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseRelease:
			// If the mouse is released at the same position as the anchor,
			// we want to allow the user to insert text at the anchor.
			if msg.X == m.anchor.x && msg.Y == m.anchor.y {
				m.textAnchorSet(msg.X, msg.Y)
			}
			m.anchorReset()
			return m, nil
		case tea.MouseLeft:
			// When the user clicks on the mouse, we want to write the
			// character to the current position of the mouse in the grid, so
			// that we can draw it later.
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.color))
			m.canvas[msg.Y][msg.X] = style.Render(m.character)
		case tea.MouseRight:
			if m.anchorIsSet() {
				// Anchor was already set, now draw the box based on the
				// anchor and current position of the mouse.
				m.restore()
				m.DrawShape(Point{m.anchor.x, m.anchor.y}, Point{msg.X, msg.Y})
			} else {
				// anchor was not set, so set it to the current position
				// of the mouse and create a backup of the canvas.
				m.backup()
				m.anchorSet(msg.X, msg.Y)
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.textAnchorReset()
			m.anchorReset()
		case "ctrl+z":
			m.restore()
		case "backspace":
			if m.textAnchorIsSet() {
				m.textAnchorLeft()
				m.canvas[m.textAnchor.y][m.textAnchor.x] = " "
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// Use the number keys to select the color.
			// The colors are: red, green, yellow, blue, magenta, cyan, white.
			// which correspond to the colors in the terminal / lipgloss colors.
			m.color = msg.String()
		case "enter":
			if m.textAnchorIsSet() {
				m.textAnchor.x = m.textAnchor.ix
				m.textAnchorDown()
			}
		default:
			var character = msg.String()

			if strings.HasPrefix(character, "alt+") {
				character = character[4:]
			}

			if m.textAnchorIsSet() {
				// If the textAnchor is set, we want to allow the user to
				// insert text at the anchor.
				style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.color))
				m.canvas[m.textAnchor.y][m.textAnchor.x] = style.Render(character)
				m.textAnchorRight()
			}
			// Otherwise, we will want to change the character that is being
			// used to the character that the user just typed.
			m.character = character
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
