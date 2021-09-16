package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrNewsIDResource = errors.New("(NewsID) not found or empty")

	ErrNewsTitleResource = errors.New("(NewsTitle) not found or empty")

	ErrCategoryNotFound = errors.New("category not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrEmailCustomerResource = errors.New("(Email) not found or empty")

	ErrEmailPasswordNotFound = errors.New("(Email) or (Password) cannot be empty")
)
