package timer

import "fmt"

func DisplayTimeRemaining(mins, secs int, timerType string) {
	fmt.Printf("\r%s: %02d:%02d restantes", timerType, mins, secs)
}
