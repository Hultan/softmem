package softmem

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	//soundPlayer "github.com/hultan/softmem/internal/soundPlayer"
	"os"
)

type MainForm struct {
	window      *gtk.ApplicationWindow
	builder     *SoftBuilder
	//soundPlayer *soundPlayer.SoundPlayer
	picker      *NumberPicker

	page0 *page0
	//page1		*page1
	page2 *page2
}

func NewMainWindow() *MainForm {
	mainWindow := new(MainForm)
	return mainWindow
}

func (m *MainForm) OnStartup() {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new gtk helper
	m.builder = newSoftBuilder("main.glade")
}

func (m *MainForm) OnActivate(app *gtk.Application) {
	// Get the main window from the glade file
	m.window = m.builder.getObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.window.SetApplication(app)
	m.window.SetTitle(fmt.Sprintf("%s - %s", applicationName, applicationVersion))
	m.window.SetDefaultSize(800, 600)

	// Hook up the destroy event
	_ = m.window.Connect("destroy", m.onCloseMainWindow)

	// Initialize number picker
	picker := NewNumberPicker()
	m.picker = picker

	// Initialize page 0
	m.page0 = NewPage0(m)
	m.page0.init()

	// Initialize page 2
	m.page2 = NewPage2(m)
	m.page2.init()

	// Initialize sound player
	//player, err := soundPlayer.NewSoundPlayer()
	//errorCheck(err)
	//m.soundPlayer = player

	// Show the main window
	m.window.ShowAll()
}

func (m *MainForm) OnShutdown() {
	m.window = nil
	m.builder = nil
	m.picker = nil
	//m.soundPlayer = nil

	m.page0.onShutDown()
}

func (m *MainForm) onCloseMainWindow() {

	buttons := []Button{{"Absolutely!", gtk.RESPONSE_OK}, {"Hell no!", gtk.RESPONSE_CANCEL}}
	box2 := NewMessageBoxWithButtons("Update statistics?", "Do you want to update statistics?", m.window, buttons)

	if box2.Warning() == gtk.RESPONSE_OK {
		m.page0.onCloseMainWindow()
	}
	//box := messagebox.NewMessageBox("Update statistics?", "Do you want to update statistics?", m.window)
	//if box.QuestionYesNo() {
	//	m.picker.UpdateStatistics()
	//}

	//m.soundPlayer.Close()
}
