package request

import "RentSpace/usecases/users"

type Users struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (us *Users) GotoDomain() users.Domain {
	return users.Domain{
		Name:     us.Name,
		Password: us.Password,
	}
}
