package ai_client

type GetRateByTwoRatingsResponse struct {
	Stations []string `json:"stations"`
	Period   []string `json:"dates"`
}

type GetBrigV1Request struct {
	Text string `json:"text"`
}

type GetBrigV1Response struct {
	Stations []string `json:"stations"`
	Period   []string `json:"dates"`
}

type BrigV1Response struct {
	Stations []string `json:"stations"`
	Period   []string `json:"dates"`
}

type GetVendorAIEURequest struct {
	Text string `json:"text"`
}

type GetVendorAIEUResponse struct {
	Stations []string `json:"stations"`
	Period   []string `json:"dates"`
}

type GetVendorAIRURequest struct {
	Text string `json:"text"`
}

type GetVendorAIRUResponse struct {
	Stations []string `json:"stations"`
	Period   []string `json:"dates"`
}
