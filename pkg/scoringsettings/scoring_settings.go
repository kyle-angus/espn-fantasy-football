package scoringsettings

import "github.com/kyle-angus/espn-fantasy-football-client/pkg/scoringitem"

type ScoringSettings struct {
	AllowOutOfPositionScoring bool                      `json:"allowOUtOfPositionScoring"`
	HomeTeamBonus             int                       `json:"homeTeamBonus"`
	MatchupTieRule            string                    `json:"matchupTieRule"` // @todo: appears to be enum
	MatchupTieRuleBy          int                       `json:"matchupTieRuleBy"`
	PlayerRankType            string                    `json:"playerRankType"` // @todo: appears to be enum
	PlayoffHomeTeamBonus      int                       `json:"playoffHomeTeamBonus"`
	PlayoffMatchupTieRule     string                    `json:"playoffMatchupTieRule"` // @todo: appears to be enum
	PlayoffMatchupTieRuleBy   int                       `json:"playoffMatchupTieRuleBy"`
	SccoringItems             []scoringitem.ScoringItem `json:"scoringItems"`
	ScoringType               string                    `json:"scoringType"` // @todo: appears to be enum
}
