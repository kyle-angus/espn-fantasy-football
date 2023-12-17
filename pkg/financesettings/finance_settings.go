package financesettings

type FinanceSettings struct {
	EntryFee           float32 `json:"entryFee"`
	MiscFee            float32 `json:"miscFee"`
	PerLoss            float32 `json:"perLoss"`
	PerTrade           float32 `json:"perTrade"`
	PlayerAcquisition  float32 `json:"playerAcquisition"`
	PlayerDrop         float32 `json:"playerDrop"`
	PlayerMoveToActive float32 `json:"playerMoveToActive"`
	PlayerMoveToIR     float32 `json:"playerMoveToIR"`
}
