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
	images        [5]*gtk.Image
	lastImages    [5]*gtk.Image
	numbers       [5]string
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

	_= p.entry.Connect("activate", p.onEntryActivate)
	_ = p.entry.Connect("key-press-event", p.onKeyPressed)

	p.loadImages()
}

func (p *page2) getGTKObjects() {
	entry := p.MainForm.builder.getObject("page2_entry_answer").(*gtk.Entry)
	answer := p.MainForm.builder.getObject("page2_label_answer").(*gtk.Label)
	hint := p.MainForm.builder.getObject("page2_label_hint").(*gtk.Label)

	p.entry = entry
	p.answer = answer
	p.hint = hint

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("page2_image%v", i)
		image := p.MainForm.builder.getObject(name).(*gtk.Image)
		p.images[i] = image

		name = fmt.Sprintf("page2_last_image%v", i)
		image = p.MainForm.builder.getObject(name).(*gtk.Image)
		p.lastImages[i] = image
	}
}

func (p *page2) loadImages() {
	p.moveImagesToLastImages()
	var imagePath = ""
	var number = 0
	for i := 0; i < 5; i++ {
		number = rand.Intn(110)
		if number < 10 {
			// Get single digit path
			p.numbers[i] = strconv.Itoa(number)
			imagePath = getImagePathByString(p.numbers[i])
		} else {
			// Get double digits path
			numberAsString := strconv.Itoa(number - 10)
			if len(numberAsString) == 1 {
				numberAsString = "0" + numberAsString
			}
			p.numbers[i] = numberAsString
			imagePath = getImagePathByString(numberAsString)
		}
		p.images[i].SetFromFile(imagePath)

	}
	p.hint.SetText("")
	p.correctAnswer = p.getCorrectAnswer()
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
		p.answer.SetText(fmt.Sprintf("CORRECT! The following number, %v, corresponds to the following images:", answer))
	} else {
		p.answer.SetText(fmt.Sprintf("WRONG! Your answer was %v, the correct answer was %v:", answer, p.correctAnswer))
	}

	//answers := strings.Split(answer, " ")
	//for _, value := range answers {
	//	// TODO : Update statistics, correct +=1 and -=1
	//}

	p.entry.SetText("")
	p.loadImages()
}

func (p *page2) getCorrectAnswer() string {
	var result = ""

	for i := 0; i < 5; i++ {
		if len(result) > 0 {
			result += " "
		}
		result += p.numbers[i]
	}

	return result
}

func (p *page2) moveImagesToLastImages() {
	for i := 0; i < 5; i++ {
		p.lastImages[i].SetFromFile(getImagePathByString(p.numbers[i]))
	}
}
