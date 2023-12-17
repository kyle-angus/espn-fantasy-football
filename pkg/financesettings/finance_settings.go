package financesettings

type FinanceSettings struct {
	EntryFee           int `json:"entryFee"`
	MiscFee            int `json:"miscFee"`
	PerLoss            int `json:"perLoss"`
	PerTrade           int `json:"perTrade"`
	PlayerAcquisition  int `json:"playerAcquisition"`
	PlayerDrop         int `json:"playerDrop"`
	PlayerMoveToActive int `json:"playerMoveToActive"`
	PlayerMoveToIR     int `json:"playerMoveToIR"`
}
