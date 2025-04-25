package config

type TimerSettings struct {
	WorkMinutes             int
	ShortBreakMinutes       int
	LongBreakMinutes        int
	PomodorosUntilLongBreak int
	SoundFilePath           string
}

func NewSettings(workMinutes int, shortBreakMinutes int, longBreakMinutes int, pomodorosUntilLongBreak int) TimerSettings {
	return TimerSettings{
		WorkMinutes:             workMinutes,
		ShortBreakMinutes:       shortBreakMinutes,
		LongBreakMinutes:        longBreakMinutes,
		PomodorosUntilLongBreak: pomodorosUntilLongBreak,
		SoundFilePath:           "assets/notification.mp3",
	}
}

func DefaultSettings() TimerSettings {
	return TimerSettings{
		WorkMinutes:             25,
		ShortBreakMinutes:       5,
		LongBreakMinutes:        15,
		PomodorosUntilLongBreak: 4,
		SoundFilePath:           "assets/notification.mp3",
	}
}
