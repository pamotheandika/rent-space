package response

import "RentSpace/usecases/users"

type Response struct {
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Password string       `json:"password"`
	Role     RoleResponse `json:"role"`
}

type RoleResponse struct {
	Name string `json:"name"`
}

func fromDomain(us *users.Domain) Response {
	return Response{
		Name:     us.Name,
		Password: us.Password,
		Role: RoleResponse{
			Name: us.Role.Name,
		},
	}
}
