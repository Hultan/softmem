package mainWindow

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	player "github.com/hultan/softmemo/sounds"
	gtkhelper "github.com/hultan/softteam/gtk"
	"os"
	"strconv"
	"strings"
)

type MainWindow struct {
	window *gtk.ApplicationWindow
	picker *NumberPicker
	player *player.Player
	label  *gtk.Label
	hint   *gtk.Label
}

func NewMainWindow() *MainWindow {
	mainWindow := new(MainWindow)
	return mainWindow
}

func (m *MainWindow) OpenMainWindow(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new gtk helper
	helper := gtkhelper.GtkHelperNewFromFile("resources/main.glade")
	// Get the main window from the glade file
	mainWindow, err := helper.GetApplicationWindow("main_window")
	errorCheck(err)

	m.window = mainWindow

	// Set up main window
	mainWindow.SetApplication(app)
	mainWindow.SetTitle(applicationTitle)
	mainWindow.SetDefaultSize(800, 600)

	// Hook up the destroy event
	mainWindow.Connect("destroy", func() {
		m.CloseMainWindow()
	})

	label, err := helper.GetLabel("memory_label")
	answer, err := helper.GetLabel("answer")
	entry, err := helper.GetEntry("memory_entry")
	image, err := helper.GetImage("image")
	hint, err := helper.GetLabel("hint_label")
	m.label = label
	m.hint = hint

	errorCheck(err)

	picker := NewNumberPicker()
	picker.Initialize()
	m.picker = picker

	player := player.NewPlayer()
	player.Initialize()
	m.player = player

	m.getNextNumber()

	entry.Connect("activate", func() {
		var result = ""

		text, err := entry.GetText()
		errorCheck(err)
		if strings.ToLower(m.picker.current.Word) == strings.ToLower(text) {
			result = fmt.Sprintf("CORRECT : %s = %s", m.picker.current.Number, m.picker.current.Word)
			m.picker.current.Correct += 1
			player.PlayCorrect()
		} else {
			result = fmt.Sprintf("WRONG : %s = %s", m.picker.current.Number, m.picker.current.Word)
			m.picker.current.Correct -= 1
			player.PlayIncorrect()
		}
		m.picker.current.HasChanged = true
		answer.SetText(result)
		entry.SetText("")

		image.SetFromFile(m.GetImagePath())

		m.getNextNumber()
	})

	entry.Connect("key-press-event", func(entry *gtk.Entry, event *gdk.Event) {
		keyEvent := gdk.EventKeyNewFromEvent(event)

		if keyEvent.KeyVal() == keyValF1 {
			hint.SetText(m.GetHint(m.picker.current.Number))
		} else if keyEvent.KeyVal() == keyValF2 {
			entry.Activate()
		}
	})
	entry.GrabFocus()

	// Show the main window
	mainWindow.ShowAll()
}

func (m *MainWindow) CloseMainWindow() {
	m.picker.UpdateStatistics()
	m.player.Close()
}

func (m *MainWindow) getNextNumber() {
	next, err := m.picker.GetNextNumber()
	if err != nil {
		panic(err)
	}
	m.label.SetText(fmt.Sprintf("%s", next))
	m.hint.SetText("")
}

func (m *MainWindow) GetHint(number string) string {
	var result string
	num := strings.Trim(number, " ")

	for i := 0; i < len(num); i++ {
		if result != "" {
			result += " : "
		}
		num, _ := strconv.Atoi(num[i : i+1])
		result += m.GetMnemonicsForNumber(num)
	}

	return "(" + result + ")"
}

func (m *MainWindow) GetMnemonicsForNumber(number int) string {
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

func (m *MainWindow) GetImagePath() string {
	length := len(strings.Trim(m.picker.current.Number, " "))
	if length == 1 {
		return m.GetSingleImagePath()
	} else {
		return m.GetDoubleImagePath()
	}
}

func (m *MainWindow) GetSingleImagePath() string {
	return fmt.Sprintf("resources/images/single/%s.png", strings.Trim(m.picker.current.Number, " "))
}

func (m *MainWindow) GetDoubleImagePath() string {
	return fmt.Sprintf("resources/images/double/%s.png", strings.Trim(m.picker.current.Number, " "))
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
