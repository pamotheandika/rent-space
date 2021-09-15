package request

import "RentSpace/businesses/transactions"

type Transaction struct {
	InvoiceNumber   string `json:"invoicenumber"`
	TransactionDate string `json:"transactiondate"`
	IDSpace         int    `json:"idspace"`
	IDOwner         int    `json:"idowner"`
	IDCustomer      int    `json:"idcustomer"`
	RentTotalMonth  int    `json:"rentotalmonth"`
	Cost            int    `json:"cost"`
	TotalCost       int    `json:"totalcost"`
}

func (req *Transaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		InvoiceNumber:   req.InvoiceNumber,
		TransactionDate: req.TransactionDate,
		IDSpace:         req.IDSpace,
		IDOwner:         req.IDOwner,
		IDCustomer:      req.IDCustomer,
		RentTotalMonth:  req.RentTotalMonth,
		Cost:            req.Cost,
		TotalCost:       req.TotalCost,
	}
}
