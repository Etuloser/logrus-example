package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	logger := New()
	logger.Info("success")
	assert.NotNil(t, logger)
}
