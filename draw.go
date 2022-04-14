package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	canvas    [][]string
	character string
	color     string
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.canvas = make([][]string, msg.Height-1)
		for i := range m.canvas {
			m.canvas[i] = make([]string, msg.Width)
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseLeft:
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.color))
			m.canvas[msg.Y][msg.X] = style.Render(m.character)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.color = msg.String()
		default:
			m.character = msg.String()
		}
	}
	return m, nil
}

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
	return s.String()
}
