package drivers

import (
	customerDomain "RentSpace/businesses/customers"
	customerDB "RentSpace/drivers/databases/customer"

	ownerDomain "RentSpace/businesses/owners"
	ownerDB "RentSpace/drivers/databases/owner"

	spaceDomain "RentSpace/businesses/spaces"
	spaceDB "RentSpace/drivers/databases/space"

	transactionDomain "RentSpace/businesses/transactions"
	transactionDB "RentSpace/drivers/databases/transaction"

	paymentDomain "RentSpace/businesses/payments"
	paymentDB "RentSpace/drivers/databases/payment"

	districtDomain "RentSpace/businesses/districts"
	districtDB "RentSpace/drivers/databases/district"

	"gorm.io/gorm"
)

func NewCustomerRepository(conn *gorm.DB) customerDomain.Repository {
	return customerDB.NewMySQLRepository(conn)
}

func NewOwnerRepository(conn *gorm.DB) ownerDomain.Repository {
	return ownerDB.NewMySQLRepository(conn)
}

func NewSpaceRepository(conn *gorm.DB) spaceDomain.Repository {
	return spaceDB.NewMySQLRepository(conn)
}

func NewTransactionRepository(conn *gorm.DB) transactionDomain.Repository {
	return transactionDB.NewMySQLRepository(conn)
}

func NewPaymentRepository(conn *gorm.DB) paymentDomain.Repository {
	return paymentDB.NewMySQLRepository(conn)
}

func NewDistrictRepository(conn *gorm.DB) districtDomain.Repository {
	return districtDB.NewMySQLRepository(conn)
}
