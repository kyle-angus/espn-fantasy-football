package tradesettings

type TradeSettings struct {
	AllowOutOfUniverse bool `json:"allowOutOfUniverse"`
	DeadlineDate       uint `json:"deadlineDate"`
	Max                int  `json:"max"`
	RevisionHours      int  `json:"revisionHours"`
	VetoVotesRequired  int  `json:"vetoVotesRequired"`
}
