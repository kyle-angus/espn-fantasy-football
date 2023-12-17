package test

import (
	"testing"

	"github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"
	"github.com/stretchr/testify/assert"
)

func TestNew_ReturnsNewInstanceWithProvidedConfig(t *testing.T) {
	espn := espnclient.New(0, "", "")
	assert.Nil(t, espn)

	espn = espnclient.New(12345678, "", "")
	assert.Nil(t, espn)

	espn = espnclient.New(12345678, "test_espns2", "")
	assert.Nil(t, espn)

	espn = espnclient.New(12345678, "test_espns2", "test_swid")
	assert.NotNil(t, espn)
}
