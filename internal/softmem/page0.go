package softmem

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"strings"
)

type page0 struct {
	MainForm      *MainForm
	answer        *gtk.Label
	image         *gtk.Image
	entry         *gtk.Entry
	label         *gtk.Label
	hint          *gtk.Label
	currentNumber string
	currentWord   string
}

func NewPage0(m *MainForm) *page0 {
	page0 := new(page0)
	page0.MainForm = m
	return page0
}

func (p *page0) init() {
	p.getGTKObjects()

	p.getNextNumber()

	_, err := p.entry.Connect("activate", p.onEntryActivate)
	errorCheck(err)
	_, err = p.entry.Connect("key-press-event", p.onKeyPressed)
	errorCheck(err)
	p.entry.GrabFocus()
}

func (p *page0) getGTKObjects() {
	label, err := p.MainForm.helper.GetLabel("page0_label_memory")
	errorCheck(err)
	answer, err := p.MainForm.helper.GetLabel("page0_label_answer")
	errorCheck(err)
	entry, err := p.MainForm.helper.GetEntry("page0_entry_answer")
	errorCheck(err)
	image, err := p.MainForm.helper.GetImage("page0_image")
	errorCheck(err)
	hint, err := p.MainForm.helper.GetLabel("page0_label_hint")
	errorCheck(err)

	p.answer = answer
	p.image = image
	p.entry = entry
	p.label = label
	p.hint = hint
}

func (p *page0) getNextNumber() {
	next, err := p.MainForm.picker.GetNextNumber()
	if err != nil {
		panic(err)
	}
	p.label.SetText(fmt.Sprintf("%s", next))
	p.hint.SetText("")
}

func (p *page0) onEntryActivate() {
	var result = ""

	text, err := p.entry.GetText()
	errorCheck(err)
	if strings.ToLower(p.currentWord) == strings.ToLower(text) {
		result = fmt.Sprintf("CORRECT : %s = %s", p.currentNumber, p.currentWord)

		// TODO : Update statistics, Correct += 1

		// Ignore sound errors
		_ = p.MainForm.soundPlayer.PlayCorrect()
	} else {
		result = fmt.Sprintf("WRONG : %s = %s", p.currentNumber, p.currentWord)

		// TODO : Update statistics, Correct -= 1

		// Ignore sound errors
		_ = p.MainForm.soundPlayer.PlayIncorrect()
	}
	// TODO : Set has changed
	//p.MainForm.picker.current.HasChanged = true

	p.answer.SetText(result)
	p.entry.SetText("")

	p.image.SetFromFile(getImagePathByString(p.currentNumber))

	p.getNextNumber()
}

func (p *page0) onKeyPressed(entry *gtk.Entry, event *gdk.Event) {
	keyEvent := gdk.EventKeyNewFromEvent(event)

	if keyEvent.KeyVal() == keyValF1 {
		p.hint.SetText(getHint(p.currentNumber))
	} else if keyEvent.KeyVal() == keyValF2 {
		entry.Activate()
	}
}

func (p *page0) onShutDown() {
	p.answer = nil
	p.entry = nil
	p.label = nil
	p.hint = nil
	p.image = nil
}

func (p *page0) onCloseMainWindow() {
	p.MainForm.picker.UpdateStatistics()
}
