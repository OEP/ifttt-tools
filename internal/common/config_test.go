package common_test

import (
	"os"
	"testing"

	"github.com/OEP/ifttt-tools/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultConfig(t *testing.T) {
	assert := assert.New(t)
	oldValue := os.Getenv("IFTTT_KEY")

	os.Setenv("IFTTT_KEY", "key")
	cfg := common.NewDefaultConfig()
	assert.Equal(cfg.GetIFTTTKey(), "key")

	os.Setenv("IFTTT_KEY", oldValue)
}
