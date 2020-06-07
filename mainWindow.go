package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	gtkhelper "github.com/hultan/softteam/gtk"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type MainWindow struct {
	window        *gtk.ApplicationWindow
	database      *database
	PEG           map[string]string
	currentNumber string
	currentWord   string
	isSingle      bool
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
	//mainWindow.SetDefaultSize(800, 600)

	// Hook up the destroy event
	mainWindow.Connect("destroy", func() {
		m.CloseMainWindow()
	})

	label, err := helper.GetLabel("memory_label")
	answer, err := helper.GetLabel("answer")
	entry, err := helper.GetEntry("memory_entry")
	image, err := helper.GetImage("image")
	upper, err := helper.GetSpinButton("upper")
	lower, err := helper.GetSpinButton("lower")
	hint,err := helper.GetLabel("hint_label")
	includeSingle, err := helper.GetCheckButton("include_single_digit_checkbox")

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
	upperLimit := upper.GetValueAsInt()
	lowerLimit := lower.GetValueAsInt()
	singleLimit := includeSingle.GetActive()
	next := m.GetNextNumber(lowerLimit, upperLimit, singleLimit)
	label.SetText(fmt.Sprintf("%s", next))
	hint.SetText("")

	entry.Connect("activate", func() {
		var result = ""

		text, err := entry.GetText()
		errorCheck(err)
		if strings.ToLower(m.currentWord) == strings.ToLower(text) {
			result = fmt.Sprintf("CORRECT : %s = %s", m.currentNumber, m.currentWord)
		} else {
			result = fmt.Sprintf("WRONG : %s = %s", m.currentNumber, m.currentWord)
		}
		answer.SetText(result)
		entry.SetText("")

		if m.isSingle {
			image.SetFromFile(fmt.Sprintf("images/single/%s.png", strings.Trim(m.currentNumber, " ")))
		} else {
			image.SetFromFile(fmt.Sprintf("images/double/%s.png", strings.Trim(m.currentNumber, " ")))
		}
		upperLimit := upper.GetValueAsInt()
		lowerLimit := lower.GetValueAsInt()
		singleLimit := includeSingle.GetActive()
		next := m.GetNextNumber(lowerLimit, upperLimit, singleLimit)
		label.SetText(fmt.Sprintf("%s", next))
		hint.SetText("")
	})

	entry.Connect("key-press-event", func(entry *gtk.Entry, event *gdk.Event) {
		keyEvent := gdk.EventKeyNewFromEvent(event)

		if keyEvent.KeyVal() == 65470 {
			hint.SetText(m.GetMnemonics(m.currentNumber))
		}
	})
	entry.GrabFocus()

	// Show the main window
	mainWindow.ShowAll()
}

func (m *MainWindow) CloseMainWindow() {
}

func (m *MainWindow) GetNextNumber(lowerLimit, upperLimit int, singleLimit bool) string {
	// get next number
	for true {
		length := len(m.PEG)
		next := rand.Intn(length)

		var index = 0
		for key, value := range m.PEG {
			if index == next {
				if singleLimit {
					keyNum, _ := strconv.Atoi(key)
					if keyNum >= lowerLimit && keyNum <= upperLimit {
						m.currentNumber = key
						m.currentWord = value
						m.isSingle = true
						return key
					}
				} else {
					singleDigit := strings.Trim(key, " ")
					if len(singleDigit) == 2 {
						keyNum, _ := strconv.Atoi(key)
						if keyNum >= lowerLimit && keyNum <= upperLimit {
							m.currentNumber = key
							m.currentWord = value
							m.isSingle = false
							return key
						}
					}
				}
			}
			index++
		}
	}
	return ""
}

func (m *MainWindow) GetMnemonics(number string) string {
	var result string

	for i:=0;i<len(number);i++ {
		if result !="" {
			result +=  " : "
		}
		num, _:=strconv.Atoi(number[i:i+1])
		result += m.GetMnemonicsForNumber(num)
	}

	return "(" + result + ")"
}

func (m *MainWindow) GetMnemonicsForNumber(number int) string {
	switch number {
	case 0: return "S,Z"
	case 1: return "T,D"
	case 2: return "N"
	case 3: return "M"
	case 4: return "R"
	case 5: return "L"
	case 6: return "J,Sh,Ch"
	case 7: return "G,K"
	case 8: return "F,V"
	case 9: return "P,B"
	default: return ""
	}
}