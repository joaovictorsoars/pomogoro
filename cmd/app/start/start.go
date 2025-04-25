package start

import (
	"context"
	"fmt"
	"os"

	"github.com/joaovictorsoars/pomogoro/config"
	"github.com/joaovictorsoars/pomogoro/internal/discord"
	"github.com/joaovictorsoars/pomogoro/internal/session"
	"github.com/joaovictorsoars/pomogoro/pkg/utils"
	"github.com/spf13/cobra"
)

func execute() {
	fmt.Println("🍅 Pomodoro Timer iniciado!")
	fmt.Println("Pressione Ctrl+C para sair")
	fmt.Println("--------------------------------")

	discord.Login()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	utils.HandleTerminationSignals(cancel)

	var settings config.TimerSettings

	if !isDefault {
		workMinutes := utils.InputInteger("Digite os minutos de trabalho: ")
		shortBreakMinutes := utils.InputInteger("Digite os minutos da pausa curta: ")
		longBreakMinutes := utils.InputInteger("Digite os minutos da pausa longa: ")
		pomodorosUntilLongBreak := utils.InputInteger("Quantos pomodoros até a pausa longa? ")

		settings = config.NewSettings(workMinutes, shortBreakMinutes, longBreakMinutes, pomodorosUntilLongBreak)
	} else {
		settings = config.DefaultSettings()
	}

	sessionManager := session.NewManager(ctx, settings)
	sessionManager.Start()

	<-ctx.Done()
	fmt.Println("\nFinalizando aplicação...")
	os.Exit(0)
}

var isDefault bool

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicializa o timer Pomodoro",
	Long:  `Inicializa o timer Pomodoro`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func init() {
	StartCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "Use default value for pomodoro config")
}
