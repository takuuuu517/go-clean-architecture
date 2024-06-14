package repository

import (
	"cleanArchitecture/domain"
	"cleanArchitecture/ent"
	"context"
)

type UserRepository struct {
	entClient *ent.Client
}

func NewUserRepository(entClient *ent.Client) *UserRepository {
	return &UserRepository{entClient: entClient}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	entUsers, err := r.entClient.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return entUsersToDomainUsers(entUsers), nil
}

func (r *UserRepository) GetById(ctx context.Context, id int) (*domain.User, error) {
	entUser, err := r.entClient.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return entUserToDomainUser(entUser), nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	entUser, err := r.entClient.User.Create().
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entUserToDomainUser(entUser), nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	entUser, err := r.entClient.User.UpdateOneID(user.ID).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entUserToDomainUser(entUser), nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	return r.entClient.User.DeleteOneID(id).Exec(ctx)
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
