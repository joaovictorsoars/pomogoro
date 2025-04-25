package sound

import (
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

type SoundPlayer struct {
	soundFilePath string
}

func NewSoundPlayer(soundFilePath string) *SoundPlayer {
	return &SoundPlayer{
		soundFilePath: soundFilePath,
	}
}

func Play(player *SoundPlayer) {
	f, err := os.Open(player.soundFilePath)

	if err != nil {
		return
	}

	streamer, format, err := mp3.Decode(f)

	if err != nil {
		return
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
