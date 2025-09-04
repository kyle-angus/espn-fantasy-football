package models

import (
	"log"

	"github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"
	"gorm.io/gorm"
)

func Sync(db *gorm.DB, client *espnclient.EspnClient) {

	var season uint = 2025
	currentWeek := calculateCurrentWeek()

	teams, err := client.GetTeams(season)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, t := range teams {
		team := Team{
			// @todo: Setup team model and create/update it in database
		}

		// @todo: use saved team id here
		rank := Rank{
			Rank:   t.CurrentProjectedRank,
			Season: season,
			Team:   uint(team.ID),
			Week:   currentWeek,
		}
		if err := db.Create(&rank); err != nil {
			log.Printf("Failed to insert rank for team %d: %v", team.ID, err)
		}
	}
}

func calculateCurrentWeek() int {
	// @todo: finish this lol
	return 0
}
