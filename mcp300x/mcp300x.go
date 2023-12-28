//go:build linux

// Package mcp300x implements a sensor model supporting mcp300x adc sensor.
package mcp300x

import (
	"context"

	"go.viam.com/rdk/components/board/genericlinux/buses"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

var Model = resource.NewModel("hazalmestci", "sensor", "mcp3001-2-go")

// Registering the component model on init is how we make sure the new model is picked up and usable.
func init() {
	resource.RegisterComponent(
		sensor.API,
		Model,
		resource.Registration[sensor.Sensor, *Mcp300xConfig]{Constructor: newSensor})
}

func newSensor(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
	logger logging.Logger,
) (sensor.Sensor, error) {
	newConfig, err := resource.NativeConfig[*Mcp300xConfig](conf)
	if err != nil {
		return nil, err
	}
	mcp := mcp300x{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
		// Attributes necessary for this sensor component config
		pins:       newConfig.Pins,
		bus:        buses.NewSpiBus(newConfig.SpiBus),
		chipSelect: newConfig.ChipSelect,
	}

	return &mcp, nil
}

// mcp300x is a sensor device that returns values read by the channels
type mcp300x struct {
	resource.Named
	resource.AlwaysRebuild
	resource.TriviallyCloseable
	logger logging.Logger
	// Maps the sensor names which are strings to the channel pins the sensor is connected to, which are ints
	pins map[string]int
	bus  buses.SPI
	// Most of the times 0 or 1
	chipSelect string
}

// Readings return results of reading the ADC
func (s *mcp300x) Readings(ctx context.Context, _ map[string]interface{}) (map[string]interface{}, error) {
	handle, err := s.bus.OpenHandle()
	if err != nil {
		return nil, err
	}
	defer handle.Close()

	results := map[string]interface{}{}
	for name, pin := range s.pins {
		s.logger.Debugw("reading the next pin", "name", name, "pin", pin)
		var tx [2]byte
		// We need a 1 as a start bit, and before that we can have many zero's as we want
		// The next bit says whether to read single-ended mode, so we set it to 1
		// Next bit says which channel to read from
		// Next bit we set to 1 to read the most significant bit first in the response
		// The bit after that is a null bit, so you start reading data
		// Then the ten bits of data at the end should all be zeros
		if pin == 0 {
			tx[0] = byte(0b1101000)
		} else {
			// They are the same length and only differ in channel bit
			tx[0] = byte(0b1111000)
		}
		tx[1] = 0

		rx, err := handle.Xfer(ctx, 1000000, s.chipSelect, 0, tx[:])
		if err != nil {
			return nil, err
		}

		value := 0x03FF & ((int(rx[0]) << 8) | int(rx[1]))
		results[name] = value
	}

	return results, nil
}
