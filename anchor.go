package main

// anchorIsSet returns whether the anchor is set.
func (m *model) anchorIsSet() bool {
	return m.anchor.x != 0 && m.anchor.y != 0
}

// anchorSet sets the anchor to the given position.
func (m *model) anchorSet(x, y int) {
	m.anchor.x = x
	m.anchor.y = y
}

// anchorReset resets the anchor to the origin.
func (m *model) anchorReset() {
	m.anchor.x = 0
	m.anchor.y = 0
}

// textAnchorIsSet returns whether the text anchor is set.
func (m *model) textAnchorIsSet() bool {
	return m.textAnchor.x != 0 && m.textAnchor.y != 0
}

// textAnchorReset resets the text anchor to the origin.
func (m *model) textAnchorReset() {
	m.textAnchor.ix = 0
	m.textAnchor.iy = 0
	m.textAnchor.x = 0
	m.textAnchor.y = 0
}

// textAnchorSet sets the text anchor to the given position.
func (m *model) textAnchorSet(x, y int) {
	m.textAnchor.ix = x
	m.textAnchor.iy = y
	m.textAnchor.x = x
	m.textAnchor.y = y
}

// textAnchorRight moves the text anchor right by one spot.
func (m *model) textAnchorRight() {
	if m.textAnchor.x >= len(m.canvas[len(m.canvas)-1])-1 {
		return
	}
	m.textAnchor.x++
}

// textAnchorUp moves the text anchor up by one spot.
func (m *model) textAnchorUp() {
	if m.textAnchor.y <= 0 {
		return
	}
	m.textAnchor.y--
}

// textAnchorDown moves the text anchor down by one spot.
func (m *model) textAnchorDown() {
	if m.textAnchor.y >= len(m.canvas)-1 {
		return
	}
	m.textAnchor.y++
}

// textAnchorLeft moves the text anchor left by one spot.
func (m *model) textAnchorLeft() {
	if m.textAnchor.x <= 0 {
		return
	}
	m.textAnchor.x--
}
