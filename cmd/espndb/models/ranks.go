package models

import (
	"time"

	"github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"
)

type Rank struct {
	ID        uint `gorm:"primaryKey"`
	Team      uint `gorm:"foreignKey"`
	Rank      int
	Season    uint
	Week      int
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}

func fetchRanks(client *espnclient.EspnClient) {}
