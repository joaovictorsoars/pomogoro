package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Exibe a versão do Pomogoro",
	Long:  `Exibe a versão atual do Pomogoro instalada no sistema`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pomogoro v0.0.1")
	},
}
