package server

import (
	"cleanArchitecture/controller"
	"cleanArchitecture/ent"
	"cleanArchitecture/repository"
	"cleanArchitecture/useCase"
)

type controllers struct {
	user *controller.UserController
}

func newControllers(entClient *ent.Client) *controllers {
	transaction := repository.NewTransactionManager(entClient)

	userRepo := repository.NewUserRepository()
	userInteractor := useCase.NewUserInteractor(userRepo, entClient, transaction)
	userController := controller.NewUserController(userInteractor)

	return &controllers{
		user: userController,
	}
}
