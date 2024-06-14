package useCase

import "context"

type ITransaction interface {
	WithTx(ctx context.Context, fn func(tx DbClient) error) error
}

type DbClient interface {
}
