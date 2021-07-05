package ui

import (
	"log"
	"os"

	"github.com/knipferrc/fm/internal/config"
	"github.com/knipferrc/fm/internal/constants"
	"github.com/knipferrc/fm/internal/utils"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cfg := config.GetConfig()
	startDir := ""

	if len(os.Args) > 1 {
		startDir = os.Args[1]
	}

	// Get the initial directory listing to be displayed
	if _, err := os.Stat(startDir); err == nil {
		cmds = append(cmds, m.updateDirectoryListing(startDir))
	} else if cfg.Settings.StartDir == constants.HomeDirectory {
		homeDir, err := utils.GetHomeDirectory()
		if err != nil {
			log.Fatal(err)
		}

		cmds = append(cmds, m.updateDirectoryListing(homeDir))
	} else if _, err := os.Stat(cfg.Settings.StartDir); err == nil {
		cmds = append(cmds, m.updateDirectoryListing(cfg.Settings.StartDir))
	} else {
		cmds = append(cmds, m.updateDirectoryListing(cfg.Settings.StartDir))
	}

	cmds = append(cmds, textinput.Blink)
	cmds = append(cmds, spinner.Tick)

	return tea.Batch(cmds...)
}