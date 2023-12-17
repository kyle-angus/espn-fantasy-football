package settings

import (
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/acquisitionsettings"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/draftsettings"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/financesettings"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/rostersettings"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/schedulesettings"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/scoringsettings"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/tradesettings"
)

type Settings struct {
	AcquisitionSettings acquisitionsettings.AcquisitionSettings `json:"acquisitionSettings"`
	DraftSettings       draftsettings.DraftSettings             `json:"draftSettings"`
	FinanceSettings     financesettings.FinanceSettings         `json:"financeSettings"`
	IsCustomizable      bool                                    `json:"isCustomizable"`
	IsPublic            bool                                    `json:"isPublic"`
	Name                string                                  `json:"name"`
	RestrictionType     string                                  `json:"restrictionType"` // @todo: appears to be enum
	RosterSettings      rostersettings.RosterSettings           `json:"rosterSettings"`
	ScheduleSettings    schedulesettings.ScheduleSettings       `json:"scheduleSettings"`
	ScoringSettings     scoringsettings.ScoringSettings         `json:"scoringSettings"`
	Size                int                                     `json:"size"`
	TradeSettings       tradesettings.TradeSettings             `json:"tradeSettings"`
}
