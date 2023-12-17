package scoringitem

type ScoringItem struct {
	IsReverseItem   bool               `json:"isReverseItem"`
	LeagueRanking   float32            `json:"leagueRanking"`
	LeagueTotal     float32            `json:"leagueTotal"`
	Points          float32            `json:"points"`
	PointsOverrides map[string]float32 `json:"pointsOverrides"`
	StatId          int                `json:"statId"`
}
