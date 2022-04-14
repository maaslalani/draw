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
