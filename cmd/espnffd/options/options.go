package options

import "flag"

type Options struct {
	LeagueId string
	EspnS2   string
	SwID     string
}

func New(args []string) *Options {
	opts := Options{}

	flag.StringVar(&opts.LeagueId, "league", "", "ESPN Fantasy Football League ID")
	flag.StringVar(&opts.EspnS2, "espns2", "", "ESPN Fantasy Football EpsnS2 Token")
	flag.StringVar(&opts.SwID, "swid", "", "ESPN Fantasy Football SW ID Token")

	flag.CommandLine.Parse(args)

	return &opts
}
