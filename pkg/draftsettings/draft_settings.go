package draftsettings

import "time"

type DraftSettings struct {
	AuctionBudget     int       `json:"auctionBudget"`
	AvailableDate     time.Time `json:"availableDate"`
	Date              time.Time `json:"date"`
	IsTradingEnabled  bool      `json:"isTradingEnabled"`
	KeeperCount       bool      `json:"keeperCount"`
	KeeperCountFuture int       `json:"keeperCountFuture"`
	KeeperOrderType   string    `json:"keeperOrderType"` // @todo: appears to be enum
	LeagueSubType     string    `json:"leagueSubType"`   // @todo: appears to be enum
	OrderType         string    `json:"orderType"`       // @todo: appears to be enum
	PickOrder         []int     `json:"pickOrder"`
	TimePerSelection  int       `json:"timePerSelection"`
	Type              string    `json:"type"` // @todo: appears to be enum
}
