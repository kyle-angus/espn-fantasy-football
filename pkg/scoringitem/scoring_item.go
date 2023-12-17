package scoringitem

type ScoringItem struct {
	IsReverseItem   bool           `json:"isReverseItem"`
	LeagueRanking   int            `json:"leagueRanking"`
	LeagueTotal     int            `json:"leagueTotal"`
	Points          int            `json:"points"`
	PointsOverrides map[string]int `json:"pointsOverrides"`
	StatId          int            `json:"statId"`
}
