package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func HandleTerminationSignals(cancelFunc func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\n🛑 Pomodoro Timer finalizado pelo usuário!")
		cancelFunc()
	}()
}
