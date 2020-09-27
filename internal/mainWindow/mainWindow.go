package mainWindow

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	soundPlayer "github.com/hultan/softmem/internal/soundPlayer"
	gtkHelper "github.com/hultan/softteam-tools/pkg/gtk-helper"
	"github.com/hultan/softteam/messagebox"
	"os"
	"strconv"
	"strings"
)

type MainWindow struct {
	window      *gtk.ApplicationWindow
	helper      *gtkHelper.GtkHelper
	soundPlayer *soundPlayer.SoundPlayer
	picker      *NumberPicker
	answer      *gtk.Label
	image       *gtk.Image
	entry       *gtk.Entry
	label       *gtk.Label
	hint        *gtk.Label
}

func NewMainWindow() *MainWindow {
	mainWindow := new(MainWindow)
	return mainWindow
}

func (m *MainWindow) OnStartup(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new gtk helper
	m.helper = gtkHelper.GtkHelperNewFromFile("main.glade")
}

func (m *MainWindow) OnActivate(app *gtk.Application) {
	// Get the main window from the glade file
	mainWindow, err := m.helper.GetApplicationWindow("main_window")
	errorCheck(err)
	m.window = mainWindow

	// Set up main window
	mainWindow.SetApplication(app)
	mainWindow.SetTitle(applicationTitle)
	mainWindow.SetDefaultSize(800, 600)

	// Hook up the destroy event
	m.window.Connect("destroy", m.closeMainWindow)

	m.getGTKObjects()

	picker := NewNumberPicker()
	m.picker = picker

	player, err := soundPlayer.NewSoundPlayer()
	errorCheck(err)
	m.soundPlayer = player

	m.getNextNumber()

	m.entry.Connect("activate", m.onEntryActivate)
	m.entry.Connect("key-press-event", m.onKeyPressed)
	m.entry.GrabFocus()

	// Show the main window
	mainWindow.ShowAll()
}

func (m *MainWindow) OnShutdown(app *gtk.Application) {
	m.window = nil
	m.helper = nil
	m.picker = nil
	m.answer = nil
	m.entry = nil
	m.soundPlayer = nil
	m.label = nil
	m.hint = nil
	m.image = nil
}

func (m *MainWindow) closeMainWindow() {

	buttons := []messagebox.Button{{"Absolutely!", gtk.RESPONSE_OK}, {"Hell no!", gtk.RESPONSE_CANCEL}}
	box2 := messagebox.NewMessageBoxWithButtons("Update statistics?", "Do you want to update statistics?", m.window, buttons)

	if box2.Warning() == gtk.RESPONSE_OK {
		m.picker.UpdateStatistics()
	}
	//box := messagebox.NewMessageBox("Update statistics?", "Do you want to update statistics?", m.window)
	//if box.QuestionYesNo() {
	//	m.picker.UpdateStatistics()
	//}

	m.soundPlayer.Close()
}

func (m *MainWindow) onEntryActivate() {
	var result = ""

	text, err := m.entry.GetText()
	errorCheck(err)
	if strings.ToLower(m.picker.current.Word) == strings.ToLower(text) {
		result = fmt.Sprintf("CORRECT : %s = %s", m.picker.current.Number, m.picker.current.Word)
		m.picker.current.Correct += 1
		m.soundPlayer.PlayCorrect()
	} else {
		result = fmt.Sprintf("WRONG : %s = %s", m.picker.current.Number, m.picker.current.Word)
		m.picker.current.Correct -= 1
		m.soundPlayer.PlayIncorrect()
	}
	m.picker.current.HasChanged = true
	m.answer.SetText(result)
	m.entry.SetText("")

	m.image.SetFromFile(m.getImagePath())

	m.getNextNumber()
}

func (m *MainWindow) onKeyPressed(entry *gtk.Entry, event *gdk.Event) {
	keyEvent := gdk.EventKeyNewFromEvent(event)

	if keyEvent.KeyVal() == keyValF1 {
		m.hint.SetText(m.getHint(m.picker.current.Number))
	} else if keyEvent.KeyVal() == keyValF2 {
		entry.Activate()
	}
}

func (m *MainWindow) getGTKObjects() {
	label, err := m.helper.GetLabel("memory_label")
	answer, err := m.helper.GetLabel("answer")
	entry, err := m.helper.GetEntry("memory_entry")
	image, err := m.helper.GetImage("image")
	hint, err := m.helper.GetLabel("hint_label")
	m.answer = answer
	m.image = image
	m.entry = entry
	m.label = label
	m.hint = hint

	errorCheck(err)
}

func (m *MainWindow) getNextNumber() {
	next, err := m.picker.GetNextNumber()
	if err != nil {
		panic(err)
	}
	m.label.SetText(fmt.Sprintf("%s", next))
	m.hint.SetText("")
}

func (m *MainWindow) getHint(number string) string {
	var result string
	num := strings.Trim(number, " ")

	for i := 0; i < len(num); i++ {
		if result != "" {
			result += " : "
		}
		num, _ := strconv.Atoi(num[i : i+1])
		result += m.getMnemonicsForNumber(num)
	}

	return "(" + result + ")"
}

func (m *MainWindow) getMnemonicsForNumber(number int) string {
	switch number {
	case 0:
		return "S"
	case 1:
		return "T"
	case 2:
		return "N"
	case 3:
		return "M"
	case 4:
		return "R"
	case 5:
		return "L"
	case 6:
		return "Sh"
	case 7:
		return "K"
	case 8:
		return "F"
	case 9:
		return "P"
	default:
		return ""
	}
}

func (m *MainWindow) getImagePath() string {
	length := len(strings.Trim(m.picker.current.Number, " "))
	if length == 1 {
		return m.getSingleImagePath()
	} else {
		return m.getDoubleImagePath()
	}
}

func (m *MainWindow) getSingleImagePath() string {
	return fmt.Sprintf("assets/images/single/%s.png", strings.Trim(m.picker.current.Number, " "))
}

func (m *MainWindow) getDoubleImagePath() string {
	return fmt.Sprintf("assets/images/double/%s.png", strings.Trim(m.picker.current.Number, " "))
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
