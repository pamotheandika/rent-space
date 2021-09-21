package routes

import (
	_middlewareauth "RentSpace/app/middleware"
	"RentSpace/controllers/cities"
	"RentSpace/controllers/customers"
	"RentSpace/controllers/owners"
	"RentSpace/controllers/payments"
	"RentSpace/controllers/spaces"
	"RentSpace/controllers/transactions"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware         middleware.JWTConfig
	CustomerController    customers.CustomerController
	OwnerController       owners.OwnerController
	SpaceController       spaces.SpaceController
	TransactionController transactions.TransactionController
	PaymentController     payments.PaymentController
	CitiesController      cities.CitiesController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	customers := e.Group("customers")
	customers.POST("/register", cl.CustomerController.AddCustomer)
	customers.GET("/list", cl.CustomerController.GetAllCustomer)
	customers.GET("/login", cl.CustomerController.LoginCustomer)
	customers.GET("/email", cl.CustomerController.GetByEmail)

	owners := e.Group("owners")
	owners.POST("/register", cl.OwnerController.AddOwner)
	owners.GET("/list", cl.OwnerController.GetAllOwner)
	owners.GET("/getbycity", cl.OwnerController.GetOwnerByCity)
	owners.GET("/login", cl.OwnerController.LoginOwner)

	spaces := e.Group("spaces", middleware.JWTWithConfig(cl.JWTMiddleware))
	spaces.POST("/register", cl.SpaceController.AddSpace, _middlewareauth.RoleValidation("owner"))
	spaces.GET("/list", cl.SpaceController.GetAllSpace)

	transactions := e.Group("transaction", middleware.JWTWithConfig(cl.JWTMiddleware))
	transactions.POST("/checkoutspace", cl.TransactionController.AddTransaction)
	transactions.GET("/listransactions", cl.TransactionController.GetAllTransaction)

	payments := e.Group("payment")
	payments.POST("/addpayment", cl.PaymentController.AddPayment)

	cities := e.Group("cities")
	cities.GET("/list", cl.CitiesController.GetCities)

}
