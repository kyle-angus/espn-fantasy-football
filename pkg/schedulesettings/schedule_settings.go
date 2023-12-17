package schedulesettings

import "github.com/kyle-angus/espn-fantasy-football-client/pkg/division"

type ScheduleSettings struct {
	Divisions                  []division.Division `json:"divisions"`
	MatchupPeriodCount         int                 `json:"matchupPeriodCount"`
	MatchupPeriodLength        int                 `json:"matchupPeriodLength"`
	MatchupPeriods             map[string][]int    `json:"matchupPeriods"`
	PeriodTypeId               int                 `json:"periodTypeId"`
	PlayoffMatchupPeriodLength int                 `json:"playoffMatchupPeriodLength"`
	PlayoffSeedingRule         string              `json:"playoffSeedingRule"` // @todo: appears to be enum
	PlayoffSeedingRuleBy       int                 `json:"playoffSeedingRuleBy"`
	PlayoffTeamCount           int                 `json:"playoffTeamCount"`
}
