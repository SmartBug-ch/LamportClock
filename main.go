package main

// create a struct for the clock

type clock struct {
	tick int64
}

func NewClock(tick int64) *clock {
	return &clock{
		tick: tick,
	}
}

func GetNewTick(clock clock) int64 {
	clock.tick += 1
	return clock.tick
}
func CompTicks(tickOne int64, tickTwo int64) int64 {
	if tickOne > tickTwo {
		return tickOne
	} else {
		return tickTwo
	}
}
