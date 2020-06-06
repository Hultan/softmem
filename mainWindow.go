package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	gtkhelper "github.com/hultan/softteam/gtk"
	"math/rand"
	"os"
	"strings"
	"time"
)

type MainWindow struct {
	window   *gtk.ApplicationWindow
	database *database
	PEG      map[string]string
	currentNumber string
	currentWord string
}

func NewMainWindow() *MainWindow {
	mainWindow := new(MainWindow)
	mainWindow.PEG = make(map[string]string)
	return mainWindow
}

func (m *MainWindow) OpenMainWindow(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new gtk helper
	helper := gtkhelper.GtkHelperNewFromFile("main.glade")
	// Get the main window from the glade file
	mainWindow, err := helper.GetApplicationWindow("main_window")
	errorCheck(err)

	m.window = mainWindow

	// Set up main window
	mainWindow.SetApplication(app)
	mainWindow.SetTitle("SoftTeam Memo v. 1.0")
	mainWindow.SetDefaultSize(800, 600)

	// Hook up the destroy event
	mainWindow.Connect("destroy", func() {
		m.CloseMainWindow()
	})

	label, err := helper.GetLabel("memory_label")
	entry, err := helper.GetEntry("memory_entry")
	errorCheck(err)

	// Load PEGs
	m.database = new(database)
	numbers, err := m.database.GetAllNumbers()
	errorCheck(err)
	for _, item := range numbers {
		m.PEG[item.Number] = item.Word
	}

	// Set a seed for RNG
	rand.Seed(time.Now().UTC().UnixNano())
	next := m.GetNextNumber()
	label.SetText(fmt.Sprintf("%s", next))

	entry.Connect("activate", func() {
		var result = ""

		text, err := entry.GetText()
		errorCheck(err)
		if text != "" {
			if strings.ToLower(m.currentWord) == strings.ToLower(text) {
				result = " (correct)"
			} else {
				result = fmt.Sprintf(" (WRONG : %s = %s)", m.currentNumber, m.currentWord)
			}
			entry.SetText("")
		}
		next := m.GetNextNumber()
		label.SetText(fmt.Sprintf("%s %s", next, result))
	})

	// Show the main window
	mainWindow.ShowAll()
}

func (m *MainWindow) CloseMainWindow() {
}

func (m *MainWindow) GetNextNumber() string {
	// get next number
	length := len(m.PEG)
	next := rand.Intn(length)

	var index = 0
	for key, value := range m.PEG {
		if index == next {
			m.currentNumber = key
			m.currentWord = value
			return key
		}
		index ++
	}
	return ""
}
