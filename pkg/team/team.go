package team

import (
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/draftstrategy"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/record"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/tradeblock"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/transactioncounter"
)

type Team struct {
	Abbrev                string                                `json:"abbrev"`
	CurrentProjectedRank  int                                   `json:"currentProjectedRank"`
	DivisionId            int                                   `json:"divisionId"`
	DraftDayProjectedRank int                                   `json:"draftDayProjectedRank"`
	DraftStrategy         draftstrategy.DraftStrategy           `json:"draftStrategy"`
	Id                    int                                   `json:"id"`
	IsActive              bool                                  `json:"isActive"`
	Logo                  string                                `json:"logo"`
	LogoType              string                                `json:"logoType"`
	Name                  string                                `json:"name"`
	Owners                []string                              `json:"owners"`
	PlayoffSeed           int                                   `json:"playoffSeed"`
	Points                float32                               `json:"points"`
	PointsAdjusted        float32                               `json:"pointsAdjusted"`
	PointsDelta           float32                               `json:"pointsDelta"`
	PrimaryOwner          string                                `json:"primaryOwner"`
	RankCalculatedFinal   int                                   `json:"rankCalculatedFinal"`
	RankFinal             int                                   `json:"rankFinal"`
	Record                map[string]record.Record              `json:"record"`
	TradeBlock            tradeblock.TradeBlock                 `json:"tradeBlock"`
	TransactionCounter    transactioncounter.TransactionCounter `json:"transactionCounter"`
	ValuesByStat          map[string]float32                    `json:"valuesByStat"`
	WaiverRank            int                                   `json:"waiverRank"`
}
