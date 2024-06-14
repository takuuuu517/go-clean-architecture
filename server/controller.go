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
	userRepo := repository.NewUserRepository(entClient)
	userInteractor := useCase.NewUserInteractor(userRepo)
	userController := controller.NewUserController(userInteractor)

	return &controllers{
		user: userController,
	}
}
