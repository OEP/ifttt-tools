package common

import (
	"os"
)

type Config interface {
	GetIFTTTKey() string
}

type config struct {
	iftttKey string
}

func (c *config) GetIFTTTKey() string { return c.iftttKey }

func NewDefaultConfig() Config {
	return &config{
		iftttKey: os.Getenv("IFTTT_KEY"),
	}
}
