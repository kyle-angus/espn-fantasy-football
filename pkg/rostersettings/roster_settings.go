package rostersettings

type RosterSettings struct {
	IsBenchUnlimited       bool           `json:"isBenchUnlimited"`
	IsUsingUndroppableList bool           `json:"isUsingUndroppableList"`
	LineupLocktimeType     string         `json:"lineupLocktimeType"` // @todo: appears to be enum
	LineupSlotCounts       map[string]int `json:"lineupSlotCounts"`
	LineupSlotStatLimits   map[string]int `json:"lineupSlotStatLimits"` // @todo: may be incorrect type
	MoveLimit              int            `json:"moveLimit"`
	PositionLimits         map[string]int `json:"positionLimits"`
	RosterLocktimeType     string         `json:"rosterLocktimeType"` // @todo: appears to be enum
	UniverseIds            []int          `json:"universeIds"`
}
