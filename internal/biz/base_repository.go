package biz

import "golang.org/x/net/context"

type Key interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~string
}

type Offset interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~string
}

// BaseRepository is the generic interface that all other repositories for concrete types
// would implement defining the CRUDL operations for the type.
type BaseRepository[T interface{}, K Key, O Offset] interface {
	// This is sample function only
	Create(ctx context.Context, e *T) error

	// TODO : Uncomment this because you will have to implement these in your service
	//FindById(ctx context.Context, id K) (*T, error)
	//Update(ctx context.Context, e *T) error
	//Delete(ctx context.Context, id K) error
	//List(ctx context.Context, offset O, limit int) ([]*T, error)
}
