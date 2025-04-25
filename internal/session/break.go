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

type BreakSession struct {
	ctx         context.Context
	settings    config.TimerSettings
	soundPlayer *sound.SoundPlayer
	isLongBreak bool
}

func NewBreakSession(
	ctx context.Context,
	settings config.TimerSettings,
	soundPlayer *sound.SoundPlayer,
	isLongBreak bool,
) *BreakSession {
	return &BreakSession{
		ctx:         ctx,
		settings:    settings,
		soundPlayer: soundPlayer,
		isLongBreak: isLongBreak,
	}
}

func (b *BreakSession) Execute() bool {
	var breakTime int
	var breakMsg string

	if b.isLongBreak {
		breakTime = b.settings.LongBreakMinutes
		breakMsg = "🌴 Iniciando pausa longa"
	} else {
		breakTime = b.settings.ShortBreakMinutes
		breakMsg = "☕ Iniciando pausa curta"
	}

	select {
	case <-b.ctx.Done():
		return false
	default:
		fmt.Printf("%s de %d minutos\n", breakMsg, breakTime)
	}

	breakTimer := timer.NewTimer(breakTime*60, "Pausa")
	endTime := time.Duration(breakTime) * time.Minute

	discord.UpdatePresence(fmt.Sprintf("💤 Descansando por %d", breakTime), time.Now(), time.Now().Add(endTime))
	if !breakTimer.Run(b.ctx, timer.DisplayTimeRemaining) {
		return false
	}

	select {
	case <-b.ctx.Done():
		return false
	default:
		sound.Play(b.soundPlayer)
		fmt.Println("\n⏰ Pausa finalizada! Pronto para voltar ao trabalho?")
		fmt.Println("--------------------------------")
		return true
	}
}
