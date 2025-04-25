package session

import (
	"context"

	"github.com/joaovictorsoars/pomogoro/config"
	"github.com/joaovictorsoars/pomogoro/internal/sound"
)

type Manager struct {
	ctx           context.Context
	settings      config.TimerSettings
	soundPlayer   *sound.SoundPlayer
	pomodoroCount int
}

func NewManager(ctx context.Context, settings config.TimerSettings) *Manager {

	return &Manager{
		ctx:           ctx,
		settings:      settings,
		soundPlayer:   sound.NewSoundPlayer(settings.SoundFilePath),
		pomodoroCount: 0,
	}
}

func (m *Manager) Start() {
	go m.runPomodoroLoop()
}

func (m *Manager) runPomodoroLoop() {
	for {
		select {
		case <-m.ctx.Done():
			return
		default:
			m.pomodoroCount++

			workSession := NewWorkSession(m.ctx, m.settings, m.soundPlayer, m.pomodoroCount)
			if !workSession.Execute() {
				return
			}

			select {
			case <-m.ctx.Done():
				return
			default:
				isLongBreak := (m.pomodoroCount%m.settings.PomodorosUntilLongBreak == 0)
				breakSession := NewBreakSession(m.ctx, m.settings, m.soundPlayer, isLongBreak)
				if !breakSession.Execute() {
					return
				}
			}
		}
	}
}
