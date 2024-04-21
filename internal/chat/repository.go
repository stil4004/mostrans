package chat

import "context"

type Repository interface {
	GetInfoFromBatch(ctx context.Context, req GetInfoFromBatchRequest) (GetInfoFromBatchResponse, error)
	GetOneStation(ctx context.Context, req GetOneStationRequest) (GetOneStationResponse, error)
}
