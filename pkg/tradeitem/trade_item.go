package tradeitem

type TradeItem struct {
	FromLineupSlotId    int    `json:"fromLineupSlotId"`
	FromTeamId          int    `json:"fromTeamId"`
	IsKeeper            bool   `json:"isKeeper"`
	OverallPickupNumber int    `json:"overallPickupNumber"`
	PlayerId            int    `json:"playerId"`
	ToLineupSlotId      int    `json:"toLineupSlotId"`
	ToTeamId            int    `json:"toTeamId"`
	Type                string `json:"type"`
}
