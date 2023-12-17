package response

import (
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/draftdetails"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/member"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/pendingtransaction"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/schedule"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/settings"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/status"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/team"
)

type EspnResponse struct {
	Id                  int                                     `json:"id"`
	GameId              int                                     `json:"gameId"`
	DraftDetail         draftdetails.DraftDetail                `json:"draftDetail"`
	Members             []member.Member                         `json:"members"`
	PendingTransactions []pendingtransaction.PendingTransaction `json:"pendingTransactions"`
	Schedule            []schedule.Schedule                     `json:"schedule"`
	ScoringPeriodId     int                                     `json:"scoringPeriodId"`
	SeasonId            int                                     `json:"seasonId"`
	SegmentId           int                                     `json:"segmentId"`
	Settings            settings.Settings                       `json:"settings"`
	Status              status.Status                           `json:"status"`
	Teams               []team.Team                             `json:"teams"`
}
