package main

import (
	"github.com/danigomez/gobot-exercises/2.SevenSegmentDisplay/display/sevensegment"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot"
)

func main() {

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")

	pins := []*gpio.DirectPinDriver{
		gpio.NewDirectPinDriver(firmataAdaptor, "2"),
		gpio.NewDirectPinDriver(firmataAdaptor, "3"),
		gpio.NewDirectPinDriver(firmataAdaptor, "4"),
		gpio.NewDirectPinDriver(firmataAdaptor, "5"),
		gpio.NewDirectPinDriver(firmataAdaptor, "6"),
		gpio.NewDirectPinDriver(firmataAdaptor, "7"),
		gpio.NewDirectPinDriver(firmataAdaptor, "8"),
		gpio.NewDirectPinDriver(firmataAdaptor, "9"),
	}

	work := func() {
		sevenSegment := sevensegment.NewDisplay(pins)
		sevenSegment.Print([]string{"A", "B", "C", "DP"})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		work)

	robot.Start()
}
