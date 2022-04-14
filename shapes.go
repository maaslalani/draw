package main

import "github.com/charmbracelet/lipgloss"

type Point struct {
	x, y int
}

var BorderVertical string = "│"
var BorderHorizontal string = "─"
var BorderTopLeft string = "┌"
var BorderTopRight string = "┐"
var BorderBottomLeft string = "└"
var BorderBottomRight string = "┘"

func (m *model) DrawShape(p1, p2 Point) {
	if p1.x == p2.x || p1.y == p2.y {
		m.DrawLine(p1, p2)
	} else {
		m.DrawBox(p1, p2)
	}
}

func (m *model) DrawLine(p1, p2 Point) {
	minx, miny := min(p1.x, p2.x), min(p1.y, p2.y)
	maxx, maxy := max(p1.x, p2.x), max(p1.y, p2.y)

	style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.color))

	var char string
	if minx == maxx {
		char = BorderVertical
	} else if miny == maxy {
		char = BorderHorizontal
	}

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			m.canvas[y][x] = style.Render(char)
		}
	}
}

func (m *model) DrawBox(p1, p2 Point) {
	minx, miny := min(p1.x, p2.x), min(p1.y, p2.y)
	maxx, maxy := max(p1.x, p2.x), max(p1.y, p2.y)

	style := lipgloss.NewStyle().Foreground(lipgloss.Color(m.color))

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			switch {
			case x == minx && y == miny:
				m.canvas[y][x] = style.Render(BorderTopLeft)
			case x == maxx && y == miny:
				m.canvas[y][x] = style.Render(BorderTopRight)
			case x == minx && y == maxy:
				m.canvas[y][x] = style.Render(BorderBottomLeft)
			case x == maxx && y == maxy:
				m.canvas[y][x] = style.Render(BorderBottomRight)
			case x == minx || x == maxx:
				m.canvas[y][x] = style.Render(BorderVertical)
			case y == miny || y == maxy:
				m.canvas[y][x] = style.Render(BorderHorizontal)
			}
		}
	}
}
