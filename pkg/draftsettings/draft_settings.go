package draftsettings

type DraftSettings struct {
	AuctionBudget     uint   `json:"auctionBudget"`
	AvailableDate     uint   `json:"availableDate"`
	Date              uint   `json:"date"`
	IsTradingEnabled  bool   `json:"isTradingEnabled"`
	KeeperCount       uint   `json:"keeperCount"`
	KeeperCountFuture uint   `json:"keeperCountFuture"`
	KeeperOrderType   string `json:"keeperOrderType"` // @todo: appears to be enum
	LeagueSubType     string `json:"leagueSubType"`   // @todo: appears to be enum
	OrderType         string `json:"orderType"`       // @todo: appears to be enum
	PickOrder         []int  `json:"pickOrder"`
	TimePerSelection  int    `json:"timePerSelection"`
	Type              string `json:"type"` // @todo: appears to be enum
}
