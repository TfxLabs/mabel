package main

import (
	"os"
	"strings"
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
	clientConfig          *torrent.ClientConfig
	torrentMeta           map[metainfo.Hash]torrentMetadata
	selected              metainfo.Hash
	help                  help.Model
	err                   error
	addPrompt             modelAddPrompt
	viewingTorrentDetails bool
	portStartupFailure    portStartupFailure
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

type portStartupFailure struct {
	enabled bool
	port    textinput.Model
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

func initialPortStartupFailure() portStartupFailure {
	input := textinput.New()
	input.Width = 32

	return portStartupFailure{port: input}
}

func genMabelConfig() *torrent.ClientConfig {
	config := torrent.NewDefaultClientConfig()
	config.Logger = log.Discard
	config.Seed = true

	metadataDirectory := os.TempDir()
	if metadataStorage, err := storage.NewDefaultPieceCompletionForDir(metadataDirectory); err != nil {
		config.DefaultStorage = storage.NewMMap("")
	} else {
		config.DefaultStorage = storage.NewMMapWithCompletion("", metadataStorage)
	}

	return config
}

func initialModel() (model, error) {
	config := genMabelConfig()
	client, err := torrent.NewClient(config)

	m := model{
		client:             client,
		clientConfig:       config,
		torrentMeta:        make(map[metainfo.Hash]torrentMetadata),
		help:               help.New(),
		addPrompt:          initialAddPrompt(),
		portStartupFailure: initialPortStartupFailure(),
	}

	if err != nil {
		msg := err.Error()
		switch {
		case strings.HasPrefix(msg, "subsequent listen"), strings.HasPrefix(msg, "first listen"):
			m.portStartupFailure.enabled = true
		default:
			return model{}, err
		}
	}
	return m, nil
}
