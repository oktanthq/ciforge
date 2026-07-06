package ci

import "context"

type Handler func(ctx context.Context, request any) (any, error)
type Middleware func(Handler) Handler
