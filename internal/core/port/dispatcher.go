package port

import "context"

type Dispatcher interface {
	Start(ctx context.Context)
	Dispatch(args any)
}
