package customers

import (
	"RentSpace/businesses/customers"
	controllers "RentSpace/controllers"
	"RentSpace/controllers/customers/request"
	"RentSpace/controllers/customers/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type CustomerController struct {
	customerUscase customers.Usecase
}

func NewCustomerController(cc customers.Usecase) *CustomerController {
	return &CustomerController{
		customerUscase: cc,
	}
}

func (ctrl *CustomerController) CreateToken(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.QueryParam("email")
	password := c.QueryParam("password")

	token, err := ctrl.customerUscase.CreateToken(ctx, email, password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	return controllers.NewSuccessResponse(c, response)
}

func (ctrl *CustomerController) GetByEmail(c echo.Context) error {
	ctx := c.Request().Context()
	cst := new(request.Customers)

	if err := c.Bind(cst); err != nil {
		return err
	}

	customers, _ := ctrl.customerUscase.GetByEmail(ctx, cst.Email)

	return controllers.NewSuccessResponse(c, customers)
}

func (ctrl *CustomerController) LoginCustomer(c echo.Context) error {
	ctx := c.Request().Context()
	cst := new(request.Customers)

	if err := c.Bind(cst); err != nil {
		return err
	}

	customers, _ := ctrl.customerUscase.LoginCustomer(ctx, cst.Email, cst.Password)

	return controllers.NewSuccessResponse(c, customers)
}

func (ctrl *CustomerController) GetAllCustomer(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.customerUscase.GetAllCustomer(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Customer{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseController)
}

func (ctrl *CustomerController) AddCustomer(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Customers{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.customerUscase.AddCustomer(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Success Registered")
}
