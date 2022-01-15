package main

import gloss "github.com/charmbracelet/lipgloss"

var (
	border       = gloss.RoundedBorder()
	borderWindow = gloss.NewStyle().
			Align(gloss.Center).
			BorderStyle(border).
			BorderForeground(gloss.Color("#5FFFD7"))
	title = gloss.NewStyle().
		Bold(true).
		Underline(true).
		Foreground(gloss.Color("#FFF"))
)
