package acquisitionsettings

type AcquisitionSettings struct {
	AcquisitionBudget             int      `json:"acquisitionBudget"`
	AcquisitionLimit              int      `json:"acquisitionLimit"`
	AcquisitionType               string   `json:"acquisitionType"` // @todo: appears to be an enum
	IsUsingAcquisitionBudget      bool     `json:"isUsingAcquisitionBudget"`
	MatchupAcquisitionLimit       int      `json:"matchupAcquisitionLimit"`
	MatchupLimitPerScordingPeriod bool     `json:"matchupLimitPerScoringPeriod"`
	MinimumBid                    int      `json:"minimumBid"`
	WaiverHours                   int      `json:"waiverHours"`
	WaiverOrderReset              bool     `json:"waiverOrderReset"`
	WaiverProcessDays             []string `json:"waiverProcessDays"` // @todo: day name strings, mon - sunday.
	WaiverProcessHour             int      `json:"waiverProcessHour"`
}
