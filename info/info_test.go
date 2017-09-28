package info

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	info, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.NotEmpty(t, info)
}
