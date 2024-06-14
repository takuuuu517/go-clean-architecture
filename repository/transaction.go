package repository

import (
	"cleanArchitecture/ent"
	"cleanArchitecture/useCase"
	"context"
	"fmt"
)

type Transaction struct {
	entClient *ent.Client
}

func NewTransactionManager(entClient *ent.Client) *Transaction {
	return &Transaction{entClient: entClient}
}

func (t *Transaction) WithTx(ctx context.Context, fn func(tx useCase.DbClient) error) error {
	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
