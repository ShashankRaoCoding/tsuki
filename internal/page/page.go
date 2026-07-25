// Package page defines the Page interface that every page model must implement.
package page

import tea "github.com/charmbracelet/bubbletea"

// Page is the common interface for all page models.
// Implementing this interface is all that is needed to add a new page to the app.
type Page interface {
	Init() tea.Cmd
	// Update handles a message and returns the updated page and an optional command.
	Update(tea.Msg) (Page, tea.Cmd)
	View() string
	SetSize(w, h int)
	// Title returns the human-readable name shown in the status bar.
	Title() string
}
