package main

import (
	torrent "github.com/anacrolix/torrent"
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
	home "github.com/mitchellh/go-homedir"
)

type torrentDownloadStarted struct{}

func downloadTorrent(t *torrent.Torrent) tea.Cmd {
	return func() tea.Msg {
		<-t.GotInfo()
		t.DownloadAll()
		return torrentDownloadStarted{}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width - gloss.Width(border.Right+border.Left)
		m.height = msg.Height - gloss.Width(border.Bottom+border.Top)
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			m.client.Close()
			return m, tea.Quit
		} else if m.addPrompt.enabled {
			return addPromptKeyPress(m, msg)
		} else {
			return defaultKeyPress(m, msg)
		}
	}
	return m, nil
}

func addPromptKeyPress(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEsc:
		m.addPrompt = initialAddPrompt()
	case tea.KeyEnter:
		var t *torrent.Torrent
		if m.addPrompt.magnet {
			t, _ = m.client.AddMagnet(m.addPrompt.input.Value())
		} else {
			path, _ := home.Expand(m.addPrompt.input.Value())
			t, _ = m.client.AddTorrentFromFile(path)
		}
		m.addPrompt = initialAddPrompt()
		return m, downloadTorrent(t)
	default:
		var cmd tea.Cmd
		m.addPrompt.input, cmd = m.addPrompt.input.Update(msg)
		return m, cmd
	}
	return m, nil
}

func defaultKeyPress(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		m.client.Close()
		return m, tea.Quit
	case "m":
		m.addPrompt.input.Focus()
		m.addPrompt.enabled = true
	case "t":
		m.addPrompt.magnet = false
		m.addPrompt.input.Focus()
		m.addPrompt.enabled = true
	}
	return m, nil
}
