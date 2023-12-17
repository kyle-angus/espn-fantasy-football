package tradesettings

import "time"

type TradeSettings struct {
	AllowOutOfUniverse bool      `json:"allowOutOfUniverse"`
	DeadlineDate       time.Time `json:"deadlineDate"`
	Max                int       `json:"max"`
	RevisionHours      int       `json:"revisionHours"`
	VetoVotesRequired  int       `json:"vetoVotesRequired"`
}
