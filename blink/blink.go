package main

import (
	"machine"
	"time"
)

var UNIT = 150
var DOT = 1 * UNIT
var DASH = 3 * UNIT
var INTRA_CHAR = 1 * UNIT
var INTER_CHAR = 3 * UNIT
var WORD = 7 * UNIT
var BREAK = 3 * WORD

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	helloWorld(led)
}

func helloWorld(led machine.Pin) {
	for {
		dot(led)
		dot(led)
		dot(led)
		dot(led)
		time.Sleep(time.Millisecond * time.Duration(INTER_CHAR))
		dot(led)
		time.Sleep(time.Millisecond * time.Duration(INTER_CHAR))
		dot(led)
		dash(led)
		dot(led)
		dot(led)
		time.Sleep(time.Millisecond * time.Duration(INTER_CHAR))
		dot(led)
		dash(led)
		dot(led)
		dot(led)
		time.Sleep(time.Millisecond * time.Duration(INTER_CHAR))
		dash(led)
		dash(led)
		dash(led)

		time.Sleep(time.Millisecond * time.Duration(WORD))

		dot(led)
		dash(led)
		dash(led)
		time.Sleep(time.Millisecond * time.Duration(INTER_CHAR))
		dash(led)
		dash(led)
		dash(led)
		time.Sleep(time.Millisecond * time.Duration(INTER_CHAR))
		dot(led)
		dash(led)
		dot(led)
		time.Sleep(time.Millisecond * time.Duration(INTER_CHAR))
		dot(led)
		dash(led)
		dot(led)
		dot(led)
		time.Sleep(time.Millisecond * time.Duration(INTER_CHAR))
		dash(led)
		dot(led)
		dot(led)

		time.Sleep(time.Millisecond * time.Duration(BREAK))
	}
}

func dot(led machine.Pin) {
	led.High()
	time.Sleep(time.Millisecond * time.Duration(DOT))
	led.Low()
	time.Sleep(time.Millisecond * time.Duration(INTRA_CHAR))
}

func dash(led machine.Pin) {
	led.High()
	time.Sleep(time.Millisecond * time.Duration(DASH))
	led.Low()
	time.Sleep(time.Millisecond * time.Duration(INTRA_CHAR))
}
