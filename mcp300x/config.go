package mcp300x

import (
	"errors"
)

type Mcp300xConfig struct {
	Pins       map[string]int `json:"pins"`
	SpiBus     string         `json:"spi_bus"`
	ChipSelect string         `json:"chip_select"`
}

func (config *Mcp300xConfig) Validate(path string) ([]string, error) {
	if config.SpiBus == "" {
		return nil, errors.New("you need to specify the SPI bus")
	}
	if config.ChipSelect == "" {
		return nil, errors.New("you need to specify the chip select")
	}
	return []string{}, nil
}
