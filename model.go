package main

import (
	"os"
	"time"

	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	width, height         int
	client                *torrent.Client
	torrentMeta           map[metainfo.Hash]torrentMetadata
	selected              metainfo.Hash
	help                  help.Model
	err                   error
	addPrompt             modelAddPrompt
	viewingTorrentDetails bool
}

type modelAddPrompt struct {
	enabled bool
	dir     bool
	torrent textinput.Model
	saveDir textinput.Model
}

type torrentMetadata struct {
	added   time.Time
	created time.Time
	comment string
	program string
}

func initialAddPrompt() modelAddPrompt {
	input := textinput.New()
	input.Width = 32

	s := modelAddPrompt{
		enabled: false,
		dir:     false,
		torrent: input,
		saveDir: input,
	}
	return s
}

func initialModel() (model, error) {
	config := torrent.NewDefaultClientConfig()
	config.Logger = log.Discard
	config.Seed = true

	metadataDirectory := os.TempDir()
	if metadataStorage, err := storage.NewDefaultPieceCompletionForDir(metadataDirectory); err != nil {
		config.DefaultStorage = storage.NewMMap("")
	} else {
		config.DefaultStorage = storage.NewMMapWithCompletion("", metadataStorage)
	}

	if client, err := torrent.NewClient(config); err != nil {
		return model{}, err
	} else {
		m := model{
			client:      client,
			torrentMeta: make(map[metainfo.Hash]torrentMetadata),
			help:        help.New(),
			addPrompt:   initialAddPrompt(),
		}
		return m, nil
	}
}
