package cmd

import (
	"fmt"
	"os"

	"github.com/joaovictorsoars/pomogoro/cmd/app/start"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(start.StartCmd)
}

var rootCmd = &cobra.Command{
	Use:   "pomogoro",
	Short: "CLI para gerenciamento de tempo com a tática Pomodoro",
	Long:  "Ferramenta de linha de comando para inicializar um timer utilizando a tática Pomodoro.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}
