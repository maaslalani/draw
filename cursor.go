package main

// cursorIsSet returns whether the cursor is set.
func (m *model) cursorIsSet() bool {
	return m.cursor.x != 0 && m.cursor.y != 0
}

// cursorSet sets the cursor to the given position.
func (m *model) cursorSet(x, y int) {
	m.cursor.x = x
	m.cursor.y = y
}

// cursorReset resets the cursor to the origin.
func (m *model) cursorReset() {
	m.cursor.x = 0
	m.cursor.y = 0
}
