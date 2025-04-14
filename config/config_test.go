package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	conf, err := New()
	assert.NoError(t, err)
	assert.Equal(t, 8181, conf.PORT)
}
