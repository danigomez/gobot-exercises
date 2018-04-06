package main

import (
	"github.com/danigomez/gobot-exercises/2.SevenSegmentDisplay/display/sevensegment"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot"
)

func main() {

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	connection := gobot.Connection(firmataAdaptor)
	pins := []*gpio.DirectPinDriver{
		gpio.NewDirectPinDriver(connection, "2"),
		gpio.NewDirectPinDriver(connection, "3"),
		gpio.NewDirectPinDriver(connection, "4"),
		gpio.NewDirectPinDriver(connection, "5"),
		gpio.NewDirectPinDriver(connection, "6"),
		gpio.NewDirectPinDriver(connection, "7"),
		gpio.NewDirectPinDriver(connection, "8"),
		gpio.NewDirectPinDriver(connection, "9"),
	}

	sevenSegment := sevensegment.NewDisplay(pins)

	sevenSegment.Print([]string{"A", "B"})
}
