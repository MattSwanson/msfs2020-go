package main

import (
	"fmt"
	"time"

	"github.com/lian/msfs2020-go/simconnect"
)

type Events struct {
	ToggleNavLights simconnect.DWORD
	ToggleAPMaster  simconnect.DWORD
}

func main() {
	s, err := simconnect.New("Transmit Event")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Flight Simulator!")

	events := &Events{
		ToggleNavLights: s.GetEventID(),
		ToggleAPMaster:  s.GetEventID(),
	}

	err = s.MapClientEventToSimEvent(events.ToggleNavLights, "TOGGLE_NAV_LIGHTS")
	if err != nil {
		panic(err)
	}

	err = s.MapClientEventToSimEvent(events.ToggleAPMaster, "AP_MASTER")
	if err != nil {
		panic(err)
	}

	for {

		err = s.TransmitClientID(events.ToggleNavLights, 0)
		if err != nil {
			panic(err)
		}

		err = s.TransmitClientID(events.ToggleAPMaster, 0)
		if err != nil {
			panic(err)
		}

		time.Sleep(2000 * time.Millisecond)

	}

	fmt.Println("close")

	if err = s.Close(); err != nil {
		panic(err)
	}
}
