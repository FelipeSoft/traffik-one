package port

import "context"

type BackendRepository interface {
	Save(ctx context.Context) error
}