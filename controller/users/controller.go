package users

import (
	"RentSpace/usecases/users"

	"github.com/labstack/echo"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(us users.Usecase) *UserController {
	return &UserController{
		userUsecase: us,
	}
}

func (us *UserController) UpsertDevice(c echo.Context)  error {
	// send to bussiness
	ctx := c.Request().Context()

	req := Request{}
	if err := c.Bind(&req); err != nil {
		
	}
}