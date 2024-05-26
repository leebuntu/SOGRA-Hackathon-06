package common_struct

type SearchEventRequest struct {
	UnformattedArea string `json:"area"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
}

type CommonEventListResponse struct {
	EventList []int `json:"event_list"`
}

type GetEventOutlineResponse struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	ThumbnailURL string `json:"thumbnail"`
}

type GetEventFullResponse struct {
	Title   string `json:"title"`
	Context string `json:"context"`

	PhotoURLs []string `json:"photos"`
}
