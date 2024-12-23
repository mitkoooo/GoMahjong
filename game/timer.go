package game

type Timer struct {
	elapsedTime int
}

func NewTimer() *Timer {
	newTimer := &Timer{
		elapsedTime: 0,
	}

	return newTimer
}

func Tick(t *Timer) {
	t.elapsedTime++
}