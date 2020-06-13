package player

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os"
	"time"
)

type Player struct {
	bufferCorrect   *beep.Buffer
	bufferIncorrect *beep.Buffer
}

func NewPlayer() *Player {
	return new(Player)
}

func (p *Player) Initialize() error {
	f, err := os.Open("resources/sounds/correct.wav")
	if err != nil {
		return err
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		return err
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	p.bufferCorrect = buffer
	streamer.Close()

	f, err = os.Open("resources/sounds/incorrect.wav")
	if err != nil {
		return err
	}

	streamer, format, err = wav.Decode(f)
	if err != nil {
		return err
	}
	buffer = beep.NewBuffer(format)
	buffer.Append(streamer)
	p.bufferIncorrect = buffer
	streamer.Close()

	return nil
}

func (p *Player) Close() {
	p.bufferCorrect = nil
	p.bufferIncorrect = nil
}

func (p *Player) PlayCorrect() error {

	shot := p.bufferCorrect.Streamer(0, p.bufferCorrect.Len())
	speaker.Play(shot)

	return nil
}

func (p *Player) PlayIncorrect() error {

	shot := p.bufferIncorrect.Streamer(0, p.bufferIncorrect.Len())
	speaker.Play(shot)

	return nil
}
