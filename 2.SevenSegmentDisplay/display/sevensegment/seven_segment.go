package sevensegment

import (
	"gobot.io/x/gobot/drivers/gpio"
	"log"
	"github.com/pkg/errors"
)

var segments = []string{"A", "B", "C", "D", "E", "F", "G", "DP"}

type SevenSegment struct {
	Pins []*gpio.DirectPinDriver
}

func NewDisplay(pins []*gpio.DirectPinDriver) SevenSegment {
	if len(pins) != len(segments) {
		log.Fatalf("You only entered %v pins, the required segments are %s", len(pins), segments)
	}
	display := SevenSegment{pins}
	display.clear()
	return display
}

func (display SevenSegment) Print(segments []string)  {
	for _, segment := range segments {
		selectedPin, err := display.getPinForSegment(segment)

		if err != nil {
			log.Printf("There was an error processing segment %s", segment)
			log.Print(err.Error())
			continue
		}

		selectedPin.Off()
	}
}

func (display SevenSegment) getPinForSegment(segment string) (pin *gpio.DirectPinDriver, err error) {
	index := -1

	for idx, seg := range segments {
		if seg == segment {
			index = idx
			break
		}
	}
	if index != -1 {
		pin = display.Pins[index]
	} else {
		err = errors.Errorf("The segment %s doesn't have a linked pin.", segment)
	}

	return pin, err

}

func (display SevenSegment) clear()  {
	for _, segment := range segments {
		selectedPin, err := display.getPinForSegment(segment)

		if err != nil {
			log.Printf("There was an error processing segment %s", segment)
			log.Print(err.Error())
			continue
		}

		selectedPin.On()
	}
}
