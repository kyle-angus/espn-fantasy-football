package transactioncounter

type TransactionCounter struct {
	AcquisitionBudgetSpent   int            `json:"acquisitionBudgetSpent"`
	Acquisitions             int            `json:"acquisitions"`
	Drops                    int            `json:"drops"`
	MatchupAcquisitionTotals map[string]int `json:"matchupAcquisitionTotals"`
	Misc                     int            `json:"misc"`
	MoveToActive             int            `json:"moveToActive"`
	MoveToIR                 int            `json:"moveToIR"`
	Paid                     float32        `json:"paid"`
	TeamCharges              float32        `json:"teamCharges"`
	Trades                   int            `json:"trades"`
}
