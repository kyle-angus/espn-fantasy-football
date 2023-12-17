package options

import "flag"

type Options struct {
	LeagueId uint
	EspnS2   string
	SwID     string
}

func New(args []string) *Options {
	opts := Options{}

	flag.UintVar(&opts.LeagueId, "league", 0, "ESPN Fantasy Football League ID")
	flag.StringVar(&opts.EspnS2, "espns2", "", "ESPN Fantasy Football EpsnS2 Token")
	flag.StringVar(&opts.SwID, "swid", "", "ESPN Fantasy Football SW ID Token")

	flag.CommandLine.Parse(args)

	return &opts
}
