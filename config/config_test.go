package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	path := "../config.yaml"
	config, err := Parse(path)
	assert.NoError(t, err, "There should not be any error while parsing the correct config")
	assert.NotNil(t, config, "Config should not be nil")
}
