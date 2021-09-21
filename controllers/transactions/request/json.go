package request

import "RentSpace/businesses/transactions"

type Transaction struct {
	InvoiceNumber   string `json:"invoicenumber"`
	TransactionDate string `json:"transactiondate"`
	SpaceID         int    `json:"idspace"`
	OwnerID         int    `json:"idowner"`
	CustomerID      int    `json:"idcustomer"`
	RentTotalMonth  int    `json:"rentotalmonth"`
	Cost            int    `json:"cost"`
	TotalCost       int    `json:"totalcost"`
}

func (req *Transaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		InvoiceNumber:   req.InvoiceNumber,
		TransactionDate: req.TransactionDate,
		SpaceID:         req.SpaceID,
		OwnerID:         req.OwnerID,
		CustomerID:      req.CustomerID,
		RentTotalMonth:  req.RentTotalMonth,
		Cost:            req.Cost,
		TotalCost:       req.TotalCost,
	}
}
