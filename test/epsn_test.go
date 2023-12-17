package test

import (
	"testing"

	"github.com/kyle-angus/espn-fantasy-football-client/pkg/config"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"
	"github.com/stretchr/testify/assert"
)

func TestNew_ReturnsNewInstanceWithProvidedConfig(t *testing.T) {
	cfg := &config.EspnConfig{}
	assert.NotNil(t, cfg)

	espn := espnclient.New(cfg)
	assert.Nil(t, espn)

	cfg.Season = 2023
	espn = espnclient.New(cfg)
	assert.Nil(t, espn)

	cfg.LeagueId = 12345678
	espn = espnclient.New(cfg)
	assert.Nil(t, espn)

	cfg.EspnS2 = "test_espns2"
	espn = espnclient.New(cfg)
	assert.Nil(t, espn)

	cfg.SWID = "test_swid"
	espn = espnclient.New(cfg)
	assert.NotNil(t, espn)
}
