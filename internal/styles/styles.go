// Package styles defines shared lipgloss colours and styles for Tsuki.
package styles

import "github.com/charmbracelet/lipgloss"

// Palette — moon-inspired colours.
var (
	ColorPrimary   = lipgloss.Color("#C9B8FF") // lavender
	ColorSecondary = lipgloss.Color("#89DCEB") // sky blue
	ColorAccent    = lipgloss.Color("#F5C2E7") // pink
	ColorMuted     = lipgloss.Color("#585B70") // dim grey
	ColorText      = lipgloss.Color("#CDD6F4") // off-white
	ColorError     = lipgloss.Color("#F38BA8") // red
	ColorSuccess   = lipgloss.Color("#A6E3A1") // green
)

// Text styles.
var (
	Title = lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true)

	Subtitle = lipgloss.NewStyle().
		Foreground(ColorSecondary)

	Selected = lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true)

	Normal = lipgloss.NewStyle().
		Foreground(ColorText)

	Muted = lipgloss.NewStyle().
		Foreground(ColorMuted)

	ErrorStyle = lipgloss.NewStyle().
		Foreground(ColorError)

	SuccessStyle = lipgloss.NewStyle().
		Foreground(ColorSuccess)

	Help = lipgloss.NewStyle().
		Foreground(ColorMuted)

	Divider = lipgloss.NewStyle().
		Foreground(ColorMuted)
)

// MoonBanner is the decorative header shown on the home page.
const MoonBanner = `
   .  ·  ✦   ·   ✧  ·  ✦  ·   .
    🌙   T  S  U  K  I   🌙
   ·  ✧   ·  .   ✦  ·   ✧  .  ·`
