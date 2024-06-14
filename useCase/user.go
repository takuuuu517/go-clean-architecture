package useCase

import (
	"cleanArchitecture/domain"
	"context"
	"fmt"
)

type IUserRepository interface {
	GetAll(ctx context.Context, dbClient DbClient) ([]*domain.User, error)
	GetById(ctx context.Context, dbClient DbClient, id int) (*domain.User, error)
	Create(ctx context.Context, dbClient DbClient, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, dbClient DbClient, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, dbClient DbClient, id int) error
}

type Input struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserOutputs []*UserOutput

type UserOutput struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserInteractor struct {
	repository  IUserRepository
	transaction ITransaction
	dbClient    DbClient
}

func NewUserInteractor(r IUserRepository, dbClient DbClient, transaction ITransaction) *UserInteractor {
	return &UserInteractor{repository: r, dbClient: dbClient, transaction: transaction}
}

func (i *UserInteractor) HandleGetAll(ctx context.Context) (UserOutputs, error) {
	users, err := i.repository.GetAll(ctx, i.dbClient)
	if err != nil {
		return nil, err
	}

	return domainUsersToUserOutputs(users), nil
}

func (i *UserInteractor) HandleGetById(ctx context.Context, id int) (*UserOutput, error) {
	user, err := i.repository.GetById(ctx, i.dbClient, id)
	if err != nil {
		return nil, err
	}

	return domainUserToUserOutput(user), nil
}

func (i *UserInteractor) HandleCreate(ctx context.Context, input Input) (*UserOutput, error) {
	user := domain.NewUser(nil, input.FirstName, input.LastName, input.Email)

	txErr := i.transaction.WithTx(ctx, func(tx DbClient) error {
		var err error
		user, err = i.repository.Create(ctx, tx, user)
		if err != nil {
			return err
		}

		return thisMightReturnAnError()
	})
	if txErr != nil {
		return nil, txErr
	}
	user, err := i.repository.Create(ctx, i.dbClient, user)
	if err != nil {
		return nil, err
	}

	return domainUserToUserOutput(user), nil
}

func (i *UserInteractor) HandleUpdate(ctx context.Context, id int, input Input) (*UserOutput, error) {
	user, err := i.repository.GetById(ctx, i.dbClient, id)
	if err != nil {
		return nil, err
	}

	user.Update(input.FirstName, input.LastName, input.Email)

	txErr := i.transaction.WithTx(ctx, func(tx DbClient) error {
		user, err = i.repository.Update(ctx, tx, user)
		if err != nil {
			return err
		}

		return thisMightReturnAnError()
	})
	if txErr != nil {
		return nil, txErr
	}

	return domainUserToUserOutput(user), nil
}

func (i *UserInteractor) HandleDelete(ctx context.Context, id int) error {
	return i.repository.Delete(ctx, i.dbClient, id)
}

func domainUserToUserOutput(user *domain.User) *UserOutput {
	return &UserOutput{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
func domainUsersToUserOutputs(users []*domain.User) UserOutputs {
	var outputs UserOutputs
	for _, user := range users {
		outputs = append(outputs, domainUserToUserOutput(user))
	}

	return outputs
}

func thisMightReturnAnError() error {
	//return nil
	return fmt.Errorf("errorrrrr")
}
