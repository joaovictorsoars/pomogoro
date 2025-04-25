package session

import (
	"context"
	"fmt"
	"time"

	"github.com/joaovictorsoars/pomogoro/config"
	"github.com/joaovictorsoars/pomogoro/internal/discord"
	"github.com/joaovictorsoars/pomogoro/internal/sound"
	"github.com/joaovictorsoars/pomogoro/internal/timer"
)

type WorkSession struct {
	ctx           context.Context
	settings      config.TimerSettings
	soundPlayer   *sound.SoundPlayer
	pomodoroCount int
}

func NewWorkSession(
	ctx context.Context,
	settings config.TimerSettings,
	soundPlayer *sound.SoundPlayer,
	pomodoroCount int,
) *WorkSession {
	return &WorkSession{
		ctx:           ctx,
		settings:      settings,
		soundPlayer:   soundPlayer,
		pomodoroCount: pomodoroCount,
	}
}

func (w *WorkSession) Execute() bool {
	fmt.Printf("\n🍅 Iniciando Pomodoro #%d - %d minutos de foco!\n",
		w.pomodoroCount, w.settings.WorkMinutes)

	workTime := w.settings.WorkMinutes * 60
	workTimer := timer.NewTimer(workTime, "Trabalho")

	endTime := time.Duration(workTime)

	discord.UpdatePresence(fmt.Sprintf("💼 Foco total por %dm", w.settings.WorkMinutes), time.Now(), time.Now().Add(endTime))
	if !workTimer.Run(w.ctx, timer.DisplayTimeRemaining) {
		return false
	}

	select {
	case <-w.ctx.Done():
		return false
	default:
		sound.Play(w.soundPlayer)
		fmt.Println("\n✅ Pomodoro concluído! Hora de uma pausa.")
		return true
	}
}
