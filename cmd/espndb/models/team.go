package models

import "github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"

type Team struct {
	ID           uint `gorm:"primaryKey"`
	Abbreviation string
	EspnID       uint `gorm:"index"`
	Name         string
	Owner        string `gorm:"index"`
	Logo         string
}

func fetchTeams(client *espnclient.EspnClient) {

}
