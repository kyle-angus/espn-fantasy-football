package main

import (
	"log"
	"os"
	"time"

	"github.com/kyle-angus/espn-fantasy-football-client/cmd/espndb/database"
	"github.com/kyle-angus/espn-fantasy-football-client/cmd/espndb/models"
	"github.com/kyle-angus/espn-fantasy-football-client/cmd/espnffd/options"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"
)

type RankProjection struct {
	ID        uint `gorm:"primaryKey"`
	Rank      uint
	Team      uint      `gorm:"foreignKey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt time.Time
}

var client *espnclient.EspnClient

// @todo: i should probably setup a func in the espnClient
// to validate args
func init() {
	opts := options.New(os.Args[1:])

	if opts.EspnS2 == "" {
		log.Fatalln("EspnS2 flag invalid")
	}

	if opts.LeagueId == 0 {
		log.Fatalln("Espn leagueID flag invalid")
	}

	if opts.SwID == "" {
		log.Fatalln("Espn SWID flag invalid")
	}

	client = espnclient.New(opts.LeagueId, opts.EspnS2, opts.SwID)
}

func main() {
	log.Println("ESPN data sync start")
	database.Connect()
	database.Migrate()

	db := database.Instance()

	models.Sync(db, client)
	log.Println("ESPN data sync complete")
}
