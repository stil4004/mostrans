package chat

import "context"

type UseCase interface {
	Cache()
	ProcessMessage(ctx context.Context, req ProcessMessageRequest) (ProcessMessageResponse, error)
}
