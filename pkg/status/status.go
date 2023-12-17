package status

import (
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/creationinfo"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/lastupdateinfo"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/waiverprocessstatus"
)

type Status struct {
	ActivatedDate            uint                                    `json:"activatedDate"`
	CreatedAsLeagueType      int                                     `json:"createdAsLeagueType"`
	CreationInfo             creationinfo.CreationInfo               `json:"creationInfo"`
	CurrentLeagueType        int                                     `json:"currentLeagueType"`
	CurrentMatchupPeriod     int                                     `json:"currentMatchupPeriod"`
	FinalScoringPeriod       int                                     `json:"finalScoringPeriod"`
	FirstScoringPeriod       int                                     `json:"firstScoringPeriod"`
	IsActive                 bool                                    `json:"isActive"`
	IsExpired                bool                                    `json:"isExpired"`
	IsFull                   bool                                    `json:"isFull"`
	IsPlayoffMatchupEdited   bool                                    `json:"isPlayoffMatchupEdited"`
	IsToBeDeleted            bool                                    `json:"isToBeDeleted"`
	IsViewable               bool                                    `json:"isViewable"`
	IsWaiverOrderEdited      bool                                    `json:"isWaiverOrderEdited"`
	LastUpdateInfo           lastupdateinfo.LastUpdateInfo           `json:"lastUpdateInfo"`
	LatestScoringPeriod      int                                     `json:"latestScoringPeriod"`
	PreviousSeasons          []int                                   `json:"previousSeasons"`
	StandingsUpdateDate      uint                                    `json:"standingsUpdateDate"`
	TeamsJoined              int                                     `json:"teamsJoined"`
	TransactionScoringPeriod int                                     `json:"transactionScoringPeriod"`
	WaiverLastExecutionDate  uint                                    `json:"waiverLastExecutionDate"`
	WaiverProcessStatus      waiverprocessstatus.WaiverProcessStatus `json:"waiverProcessStatus"`
}
