package users

import (
	"RentSpace/usecases/users"
	"time"
)

type Record struct {
	ID        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomain(record users.Domain) Record {
	return Record{
		ID:       record.ID,
		Name:     record.Name,
		Password: record.Password,
	}
}
