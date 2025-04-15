package chui

import "github.com/charmbracelet/lipgloss"

// ========= Banner =========
// func cText(color, msg string) string {
// 	return color + msg + "{{ .AnsiColor.Default }}"
// }

// ========= LipGloss =========

var Header = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#EEEEEE")).
	Bold(true)

var EvenRow = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#777777"))

var OddRow = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#AAAAAA"))

var Step = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#EEEEEE")).
	Bold(true)

var SubStep = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#228822")).
	Bold(true)

var Success = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#228822"))

var Info = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#EEEEEE"))

var Warning = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF8800")).
	Bold(true)

var Error = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#EE0000")).
	Bold(true)
