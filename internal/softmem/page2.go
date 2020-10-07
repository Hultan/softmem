package softmem

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"math/rand"
	"strconv"
)

type page2 struct {
	MainForm      *MainForm
	image         [5]*gtk.Image
	numbers       [5]int
	entry         *gtk.Entry
	answer        *gtk.Label
	correctAnswer string
	hint          *gtk.Label
}

func NewPage2(m *MainForm) *page2 {
	page2 := new(page2)
	page2.MainForm = m
	return page2
}

func (p *page2) init() {
	p.getGTKObjects()

	_, err := p.entry.Connect("activate", p.onEntryActivate)
	errorCheck(err)

	_, err = p.entry.Connect("key-press-event", p.onKeyPressed)
	errorCheck(err)

	p.loadImages()

	p.correctAnswer = p.getCorrectAnswer()
}

func (p *page2) getGTKObjects() {
	entry, err := p.MainForm.helper.GetEntry("page2_entry_answer")
	errorCheck(err)
	answer, err := p.MainForm.helper.GetLabel("page2_label_answer")
	errorCheck(err)
	hint, err := p.MainForm.helper.GetLabel("page2_label_hint")
	errorCheck(err)

	p.entry = entry
	p.answer = answer
	p.hint = hint

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("page2_image%v", i)
		image, err := p.MainForm.helper.GetImage(name)
		errorCheck(err)
		p.image[i] = image
	}
}

func (p *page2) loadImages() {
	var imagePath = ""
	var number = 0
	for i := 0; i < 5; i++ {
		number = rand.Intn(110)
		p.numbers[i] = number
		if number < 10 {
			// Get single digit path
			imagePath = getImagePath(number)
		} else {
			// Get double digits path
			number = number - 10
			p.numbers[i] = number
			numberAsString := strconv.Itoa(number)
			if len(numberAsString) == 1 {
				numberAsString = "0" + numberAsString
			}
			imagePath = getImagePathByString(numberAsString)
		}
		p.image[i].SetFromFile(imagePath)
	}
	p.hint.SetText("")
}

func (p *page2) onKeyPressed(entry *gtk.Entry, event *gdk.Event) {
	keyEvent := gdk.EventKeyNewFromEvent(event)

	if keyEvent.KeyVal() == keyValF1 {
		p.hint.SetText(getLongHint())
	}
}

func (p *page2) onEntryActivate(entry *gtk.Entry) {
	answer, err := entry.GetText()
	if err != nil {
		// What to do here?
	}
	if answer == p.correctAnswer {
		p.answer.SetText("Correct!")
	} else {
		p.answer.SetText(fmt.Sprintf("WRONG! Your answer was %v, the correct answer was %v", answer, p.correctAnswer))
	}

	p.entry.SetText("")
	p.loadImages()
}

func (p *page2) getCorrectAnswer() string {
	var result = ""

	for i := 0; i < 5; i++ {
		if len(result) > 0 {
			result += " "
		}
		result += strconv.Itoa(p.numbers[i])
	}

	return result
}
