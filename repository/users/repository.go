package users

import (
	"RentSpace/usecases/users"

	"gorm.io/gorm"
)

type MySqlRepository struct {
	conn *gorm.DB
}

func NewUserSQLRepository(gorm *gorm.DB) *MySqlRepository {
	return &(MySqlRepository{
		conn: gorm,
	})
}

func (uuc *MySqlRepository) UpsertUser(user users.Domain) (users.Domain, error) {
	//  ORM Todo
	return users.Domain{}, nil
}

func (uuc *MySqlRepository) InsertDevice(device users.Device, userID int) error {
	return nil
}
