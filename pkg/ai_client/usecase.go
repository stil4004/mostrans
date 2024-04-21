package ai_client

import (
	"context"
)

type UseCase interface {
	GetBrigV1(ctx context.Context, promt string) (GetBrigV1Response, error)
	GetVendorAIEU(ctx context.Context, promt string) (GetVendorAIEUResponse, error)
	GetVendorAIRU(ctx context.Context, promt string) (GetVendorAIRUResponse, error)
}
