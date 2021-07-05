package ui

import (
	"io/fs"

	"github.com/knipferrc/fm/internal/config"
	"github.com/knipferrc/fm/internal/dirtree"
	"github.com/knipferrc/fm/internal/pane"
	"github.com/knipferrc/fm/internal/statusbar"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Main App model
type model struct {
	primaryPane          pane.Model
	secondaryPane        pane.Model
	textInput            textinput.Model
	loader               spinner.Model
	dirTree              dirtree.Model
	statusBar            statusbar.Model
	previousKey          tea.KeyMsg
	itemToMove           fs.FileInfo
	activeMarkdownSource string
	previousDirectory    string
	initialMoveDirectory string
	showCommandBar       bool
	inMoveMode           bool
	ready                bool
}

func NewModel() model {
	cfg := config.GetConfig()

	// Create a new textinput
	input := textinput.NewModel()
	input.Prompt = "❯ "
	input.CharLimit = 250

	// Create a new spinner with some styling based on the config
	l := spinner.NewModel()
	l.Spinner = spinner.Dot
	l.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(cfg.Colors.Spinner))

	// Create a new dirtree
	dirTree := dirtree.NewModel(
		cfg.Settings.ShowIcons,
		cfg.Colors.DirTree.SelectedItem,
		cfg.Colors.DirTree.UnselectedItem,
	)

	// Initialize the primary pane as active and pass in some config values
	primaryPane := pane.NewModel(
		true,
		cfg.Settings.RoundedPanes,
		true,
		cfg.Colors.Pane.ActiveBorderColor,
		cfg.Colors.Pane.InactiveBorderColor,
	)

	// Initialize the secondary pane as inactive and pass in some config values
	secondaryPane := pane.NewModel(
		false,
		cfg.Settings.RoundedPanes,
		true,
		cfg.Colors.Pane.ActiveBorderColor,
		cfg.Colors.Pane.InactiveBorderColor,
	)

	// Initialize a status bar passing in config values
	statusBar := statusbar.NewModel(
		statusbar.Color{
			Background: cfg.Colors.StatusBar.SelectedFile.Background,
			Foreground: cfg.Colors.StatusBar.SelectedFile.Foreground,
		},
		statusbar.Color{
			Background: cfg.Colors.StatusBar.Bar.Background,
			Foreground: cfg.Colors.StatusBar.Bar.Foreground,
		},
		statusbar.Color{
			Background: cfg.Colors.StatusBar.TotalFiles.Background,
			Foreground: cfg.Colors.StatusBar.TotalFiles.Foreground,
		},
		statusbar.Color{
			Background: cfg.Colors.StatusBar.Logo.Background,
			Foreground: cfg.Colors.StatusBar.Logo.Foreground,
		},
	)

	return model{
		primaryPane:          primaryPane,
		secondaryPane:        secondaryPane,
		textInput:            input,
		loader:               l,
		dirTree:              dirTree,
		statusBar:            statusBar,
		previousKey:          tea.KeyMsg{},
		itemToMove:           nil,
		activeMarkdownSource: "",
		previousDirectory:    "",
		initialMoveDirectory: "",
		showCommandBar:       false,
		inMoveMode:           false,
		ready:                false,
	}
}