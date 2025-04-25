package discord

import (
	"time"

	"github.com/hugolgst/rich-go/client"
)

func Login() {
	client.Login("1365081816997953616")
}

func UpdatePresence(state string, startTime time.Time, endTime time.Time) {
	client.SetActivity(client.Activity{
		State:   state,
		Details: "Modo Pomodoro",
		Timestamps: &client.Timestamps{
			Start: &startTime,
			End:   &endTime,
		},
	})
}
