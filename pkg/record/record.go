package record

type StreakType string

var (
	Losses StreakType = "LOSS"
	Win    StreakType = "WIN"
)

type Record struct {
	GamesBack     float32    `json:"gamesBack"`
	Losses        uint       `json:"losses"`
	Percentage    float32    `json:"percentage"`
	PointsAgainst float32    `json:"pointsAgainst"`
	PointsFor     float32    `json:"pointsFor"`
	StreakLength  uint       `json:"streakLength"`
	StreakType    StreakType `json:"streakType"`
	Ties          uint       `json:"ties"`
	Wins          uint       `json:"wins"`
}
