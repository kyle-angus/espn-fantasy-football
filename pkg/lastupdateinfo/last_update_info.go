package lastupdateinfo

// @todo: matches creation info. combine both into a struct and reuse it
type LastUpdateInfo struct {
	ClientAddress string `json:"clientAddress"`
	Platform      string `json:"platform"`
	Source        string `json:"source"`
}
