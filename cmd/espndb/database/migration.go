package database

import "github.com/kyle-angus/espn-fantasy-football-client/cmd/espndb/models"

func Migrate() {
	DB.AutoMigrate(&models.Team{})
	DB.AutoMigrate(&models.Rank{})
}
