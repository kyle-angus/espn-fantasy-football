package notificationsettings

type NotificationSettings struct {
	Enabled bool   `json:"enabled"`
	Id      string `json:"id"`
	Type    string `json:"type"`
}
