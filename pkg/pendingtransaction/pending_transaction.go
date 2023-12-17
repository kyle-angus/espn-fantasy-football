package pendingtransaction

import "github.com/kyle-angus/espn-fantasy-football-client/pkg/tradeitem"

type PendingTransaction struct {
	BidAmount               int                   `json:"bidAmount"`
	ExecutionType           string                `json:"executionType"`
	Id                      string                `json:"id"`
	IsActiveAsTeamOwner     bool                  `json:"isActiveAsTeamOwner"`
	IsLeagueManager         bool                  `json:"isLeagueManager"`
	IsPending               bool                  `json:"isPending"`
	Items                   []tradeitem.TradeItem `json:"items"`
	MemberId                string                `json:"memberId"`
	ProposedDate            int                   `json:"proposedDate"`
	Rating                  int                   `json:"rating"`
	ScoringPeriodId         int                   `json:"scoringPeriodId"`
	SkipTransactionCounters bool                  `json:"skipTransactionCounters"`
	Status                  string                `json:"status"`
	SuOrder                 int                   `json:"subOrder"`
	TeamActions             map[int]bool          `json:"teamActions"`
	TeamId                  int                   `json:"teamId"`
	Type                    string                `json:"type"`
}
