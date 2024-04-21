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
	Periods  []string `json:"periods" db:"date"`
	Stations []string `json:"stations" db:"station"`
}

type GetInfoFromBatchResponse struct {
	PeopleFlow int      `json:"passenger_count" db:"passenger_count"`
	Periods    []string `json:"periods" db:"-"`
	Stations   string   `json:"stations" db:"-"`
}

type GetOneStationRequest struct {
	Station string `db:"station_name"`
	Date    string `db:"date"`
}

type GetOneStationResponse struct {
	Station string `json:"station"`
	Date    string `json:"date"`
	Flow    int    `json:"passenger_count" db:"ps_count"`
}
