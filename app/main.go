package main

import (
	_driverFactory "RentSpace/drivers"

	_customerUsecase "RentSpace/businesses/customers"
	_customerController "RentSpace/controllers/customers"
	_customerRepo "RentSpace/drivers/databases/customer"

	_ownerUsecase "RentSpace/businesses/owners"
	_ownerController "RentSpace/controllers/owners"
	_ownerRepo "RentSpace/drivers/databases/owner"

	_spaceUsecase "RentSpace/businesses/spaces"
	_spaceController "RentSpace/controllers/spaces"
	_spaceRepo "RentSpace/drivers/databases/space"

	_transactionUsecase "RentSpace/businesses/transactions"
	_transactionController "RentSpace/controllers/transactions"
	_transactionRepo "RentSpace/drivers/databases/transaction"

	_routes "RentSpace/app/routes"

	_dbDriver "RentSpace/drivers/mysql"

	_middleware "RentSpace/app/middleware"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service run on Debug Mode")
	}
}

func DBConnection(db *gorm.DB) {
	db.AutoMigrate(
		&_customerRepo.Customer{},
		&_ownerRepo.Owner{},
		&_spaceRepo.Space{},
		&_transactionRepo.Transaction{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitialDB()
	DBConnection(db)

	configJWT := _middleware.ConfigJwt{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	customerRepo := _driverFactory.NewCustomerRepository(db)
	customerUsecase := _customerUsecase.NewCustomerUsace(timeoutContext, customerRepo, &configJWT)
	customerCtrl := _customerController.NewCustomerController(customerUsecase)

	ownerRepo := _driverFactory.NewOwnerRepository(db)
	ownerUsecase := _ownerUsecase.NewOwnerUsace(timeoutContext, ownerRepo, &configJWT)
	ownerCtrl := _ownerController.NewOwnerController(ownerUsecase)

	spaceRepo := _driverFactory.NewSpaceRepository(db)
	spaceUsecase := _spaceUsecase.NewSpaceUsace(timeoutContext, spaceRepo)
	spaceCtrl := _spaceController.NewSpaceController(spaceUsecase)

	transactionRepo := _driverFactory.NewTransactionRepository(db)
	transactionUsecase := _transactionUsecase.NewTransactionUsecase(timeoutContext, transactionRepo, spaceUsecase)
	transactionCtrl := _transactionController.NewTransactionController(transactionUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:         configJWT.Init(),
		CustomerController:    *customerCtrl,
		OwnerController:       *ownerCtrl,
		SpaceController:       *spaceCtrl,
		TransactionController: *transactionCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
