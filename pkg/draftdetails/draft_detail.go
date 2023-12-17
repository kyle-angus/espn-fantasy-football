package draftdetails

type DraftDetail struct {
	CompleteDate int64 `json:"completeDate"`
	Drafted      bool  `json:"drafted"`
	InProgress   bool  `json:"inProgress"`
}
