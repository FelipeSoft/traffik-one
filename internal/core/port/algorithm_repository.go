package port

import "context"

type AlgorithmsRepository interface {
	Get(ctx context.Context) (string, error)
	Set(ctx context.Context, algorithm string) error
}
