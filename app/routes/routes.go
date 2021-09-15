package routes

import (
	"RentSpace/controllers/customers"
	"RentSpace/controllers/owners"
	"RentSpace/controllers/spaces"
	"RentSpace/controllers/transactions"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	CustomerController    customers.CustomerController
	OwnerController       owners.OwnerController
	SpaceController       spaces.SpaceController
	TransactionController transactions.TransactionController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	customers := e.Group("customers")
	customers.GET("/login", cl.CustomerController.LoginCustomer)
	customers.POST("/register", cl.CustomerController.AddCustomer)
	customers.GET("/list", cl.CustomerController.GetAllCustomer)
	customers.GET("/token", cl.CustomerController.CreateToken)
	customers.GET("/email", cl.CustomerController.GetByEmail)

	owners := e.Group("owners")
	owners.POST("/register", cl.OwnerController.AddOwner)
	owners.GET("/list", cl.OwnerController.GetAllOwner)

	spaces := e.Group("spaces")
	spaces.POST("/register", cl.SpaceController.AddSpace)
	spaces.GET("/list", cl.SpaceController.GetAllSpace)

	transactions := e.Group("transaction")
	transactions.POST("/checkoutspace", cl.TransactionController.AddTransaction)
}
