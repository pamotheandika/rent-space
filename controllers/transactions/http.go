package transactions

import (
	"RentSpace/businesses/transactions"
	"RentSpace/controllers"
	"RentSpace/controllers/transactions/request"
	"RentSpace/controllers/transactions/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUsecase transactions.Usecase
}

func NewTransactionController(tc transactions.Usecase) *TransactionController {
	return &TransactionController{
		transactionUsecase: tc,
	}
}

func (ctrl *TransactionController) GetAllTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.transactionUsecase.GetAllTransaction(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	respController := []response.Transaction{}
	for _, value := range resp {
		respController = append(respController, response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, respController)
}

func (ctrl *TransactionController) AddTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Transaction{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	err := ctrl.transactionUsecase.AddTransaction(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, "Transaction Success")
}
