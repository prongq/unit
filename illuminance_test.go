package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIlluminance(t *testing.T) {

	// SI
	assert.Equal(t, 1e0, (1 * Lux).Lux())
}
