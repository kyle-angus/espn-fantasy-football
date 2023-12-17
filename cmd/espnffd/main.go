package main

import (
	"log"
	"os"

	"github.com/kyle-angus/espn-fantasy-football-client/cmd/espnffd/options"
	"github.com/kyle-angus/espn-fantasy-football-client/cmd/espnffd/routes"
)

func main() {

	opts := options.New(os.Args[1:])

	if opts.EspnS2 == "" {
		log.Fatalln("EspnS2 flag invalid")
	}

	if opts.LeagueId == "" {
		log.Fatalln("Espn leagueID flag invalid")
	}

	if opts.SwID == "" {
		log.Fatalln("Espn SWID flag invalid")
	}

	log.Println("Starting ESPN Fantasy Football Daemon...")

	r := routes.Setup()
	r.Run()

	log.Println("ESPN Fantasy Football Daemon exiting...")

}
