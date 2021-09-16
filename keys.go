package main

import "github.com/charmbracelet/bubbles/key"

type homeKeyMap struct {
	quit       key.Binding
	help       key.Binding
	addTorrent key.Binding
}

func (k homeKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.help, k.quit}
}

func (k homeKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.addTorrent},
		{k.help, k.quit},
	}
}

var homeKeys = homeKeyMap{
	quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	addTorrent: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add torrent"),
	),
}

type addPromptKeyMap struct {
	quit    key.Binding
	back    key.Binding
	forward key.Binding
}

func (k addPromptKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.forward, k.back, k.quit}
}

func (k addPromptKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.back, k.forward},
		{k.quit},
	}
}

var addPromptKeys = addPromptKeyMap{
	quit: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("⎋", "home"),
	),
	back: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("⇧↹", "previous"),
	),
	forward: key.NewBinding(
		key.WithKeys("↵"),
		key.WithHelp("↵", "next"),
	),
}
