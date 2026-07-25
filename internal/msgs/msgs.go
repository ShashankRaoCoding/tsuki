// Package msgs defines shared message types used for inter-page communication.
package msgs

// PageID identifies a page in the application.
type PageID int

const (
	Home     PageID = iota
	Notes
	Settings
)

// NavigateMsg is sent by a page when it wants to navigate to another page.
type NavigateMsg struct {
	To PageID
}
