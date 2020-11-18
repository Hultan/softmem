package player

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os"
	"time"
)

type SoundPlayer struct {
	bufferCorrect   *beep.Buffer
	bufferIncorrect *beep.Buffer
}

func NewSoundPlayer() (*SoundPlayer, error) {
	p := new(SoundPlayer)

	buffer, format ,err := p.loadSound("assets/sounds/correct.wav")
	if err != nil {
		return nil, err
	}
	p.bufferCorrect = buffer

	buffer, format ,err = p.loadSound("assets/sounds/incorrect.wav")
	if err != nil {
		return nil, err
	}
	p.bufferIncorrect = buffer

	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	return p, nil
}

func (p *SoundPlayer) loadSound(path string) (*beep.Buffer, *beep.Format, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		return nil, nil, err
	}
	defer streamer.Close()

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)

	return buffer, &format, nil
}

func (p *SoundPlayer) Initialize() error {
	return nil
}

func (p *SoundPlayer) Close() {
	p.bufferCorrect = nil
	p.bufferIncorrect = nil
}

func (p *SoundPlayer) PlayCorrect() error {

	shot := p.bufferCorrect.Streamer(0, p.bufferCorrect.Len())
	speaker.Play(shot)

	return nil
}

func (p *SoundPlayer) PlayIncorrect() error {

	shot := p.bufferIncorrect.Streamer(0, p.bufferIncorrect.Len())
	speaker.Play(shot)

	return nil
}
