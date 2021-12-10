package ui

type Table interface {
	// SetHeaders defines table headers
	SetHeaders(...string)

	// Append adds rows to the table
	Append(...interface{})

	// Render does the actual render of the table
	Render()
}
