package member

import "github.com/kyle-angus/espn-fantasy-football-client/pkg/notificationsettings"

type Member struct {
	DisplayName          string                                      `json:"displayName"`
	FirstName            string                                      `json:"firstName"`
	Id                   string                                      `json:"id"`
	IsLeagueCreator      bool                                        `json:"isLeagueCreator"`
	IsLeagueManager      bool                                        `json:"isLeagueManager"`
	LastName             string                                      `json:"lastName"`
	NotificationSettings []notificationsettings.NotificationSettings `json:"notificationSettings"`
}
