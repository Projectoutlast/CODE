package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	dev := NewLogger("dev")
	assert.NotNil(t, dev)

	test := NewLogger("test")
	assert.NotNil(t, test)

	prod := NewLogger("prod")
	assert.NotNil(t, prod)

	unknown := NewLogger("unknown")
	assert.NotNil(t, unknown)
}
