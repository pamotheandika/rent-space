package spaces

import (
	"RentSpace/businesses/spaces"
	controllers "RentSpace/controllers"
	"RentSpace/controllers/spaces/request"
	"RentSpace/controllers/spaces/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SpaceController struct {
	spaceUsacase spaces.Usecase
}

func NewSpaceController(sc spaces.Usecase) *SpaceController {
	return &SpaceController{
		spaceUsacase: sc,
	}
}

func (ctrl *SpaceController) GetAllSpace(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.spaceUsacase.GetAllSpace(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	respController := []response.Space{}
	for _, value := range resp {
		respController = append(respController, response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, respController)
}

func (ctrl *SpaceController) AddSpace(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Space{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.spaceUsacase.AddSpace(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Success Registered")
}
