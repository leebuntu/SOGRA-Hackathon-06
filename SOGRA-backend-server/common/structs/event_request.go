package common

type SearchEventRequest struct {
	UnformattedArea string `json:"area"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
}
