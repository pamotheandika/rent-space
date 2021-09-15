package owners

import (
	"RentSpace/businesses/owners"
	controllers "RentSpace/controllers"
	"RentSpace/controllers/owners/request"
	"RentSpace/controllers/owners/response"
	"context"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type OwnerController struct {
	ownerUscase owners.Usecase
}

func NewOwnerController(cc owners.Usecase) *OwnerController {
	return &OwnerController{
		ownerUscase: cc,
	}
}

func (ctrl *OwnerController) GetAllOwner(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.ownerUscase.GetAllOwner(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Owner{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (ctrl *OwnerController) GetByEmail(c echo.Context, email string) error {
	owner, err := ctrl.ownerUscase.GetByEmail(context.Background(), email)

	if err == nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, owner)
}

func (ctrl *OwnerController) AddOwner(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Owners{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.ownerUscase.AddOwner(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Success Registered")
}
