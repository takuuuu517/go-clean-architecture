package repository

import (
	"cleanArchitecture/domain"
	"cleanArchitecture/ent"
	"cleanArchitecture/useCase"
	"context"
	"errors"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAll(ctx context.Context, dbClient useCase.DbClient) ([]*domain.User, error) {
	entClient, err := getEntClient(dbClient)
	if err != nil {
		return nil, err
	}

	entUsers, err := entClient.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return entUsersToDomainUsers(entUsers), nil
}

func (r *UserRepository) GetById(ctx context.Context, dbClient useCase.DbClient, id int) (*domain.User, error) {
	entClient, err := getEntClient(dbClient)
	if err != nil {
		return nil, err
	}

	entUser, err := entClient.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return entUserToDomainUser(entUser), nil
}

func (r *UserRepository) Create(ctx context.Context, dbClient useCase.DbClient, user *domain.User) (*domain.User, error) {
	entClient, err := getEntClient(dbClient)
	if err != nil {
		return nil, err
	}

	entUser, err := entClient.User.Create().
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entUserToDomainUser(entUser), nil
}

func (r *UserRepository) Update(ctx context.Context, dbClient useCase.DbClient, user *domain.User) (*domain.User, error) {
	entClient, err := getEntClient(dbClient)
	if err != nil {
		return nil, err
	}

	entUser, err := entClient.User.UpdateOneID(user.ID).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entUserToDomainUser(entUser), nil
}

func (r *UserRepository) Delete(ctx context.Context, dbClient useCase.DbClient, id int) error {
	entClient, err := getEntClient(dbClient)
	if err != nil {
		return err
	}

	return entClient.User.DeleteOneID(id).Exec(ctx)
}

func entUserToDomainUser(entUer *ent.User) *domain.User {
	return &domain.User{
		ID:        entUer.ID,
		FirstName: entUer.FirstName,
		LastName:  entUer.LastName,
		Email:     entUer.Email,
		CreatedAt: entUer.CreatedAt,
		UpdatedAt: entUer.UpdatedAt,
	}
}

func entUsersToDomainUsers(entUsers []*ent.User) []*domain.User {
	var domainUsers []*domain.User
	for _, entUser := range entUsers {
		domainUsers = append(domainUsers, entUserToDomainUser(entUser))
	}
	return domainUsers
}

func getEntClient(dbClient useCase.DbClient) (*ent.Client, error) {
	entClient, ok := dbClient.(*ent.Client)
	if !ok {
		tx, ok := dbClient.(*ent.Tx)
		if ok {
			entClient = tx.Client()
		} else {
			return nil, errors.New("invalid dbClient type")
		}
	}
	return entClient, nil
}
