package chat

type ChatRequest struct {
	AIType      int    `json:"ai_type"`
	MessageText string `json:"text"`
}

type ProcessMessageRequest struct {
	AIType      int    `json:"ai_type"`
	MessageText string `json:"text"`
}

type ProcessMessageResponse struct {
	ResponseMessage string `json:"message"`
}

type GetInfoFromBatchRequest struct {
	Periods  []string `json:"periods"`
	Stations []string `json:"stations"`
}

type GetInfoFromBatchResponse struct {
	Periods    []string `json:"periods"`
	Stations   []string `json:"stations"`
	PeopleFlow int      `json:"passenger_count"`
}

type GetOneStationRequest struct {
	Station string `db:"station_name"`
	Date    string `db:"date"`
}

type GetOneStationResponse struct {
}
