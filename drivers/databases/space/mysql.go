package space

import (
	"RentSpace/businesses/spaces"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type mysqlSpaceRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) spaces.Repository {
	return &mysqlSpaceRepository{
		Conn: conn,
	}
}

func (sr *mysqlSpaceRepository) GetAllSpace(ctx context.Context) ([]spaces.Domain, error) {
	rec := []Space{}

	resp := sr.Conn.Find(&rec)

	if resp.Error != nil {
		return []spaces.Domain{}, resp.Error
	}

	spaceDomains := []spaces.Domain{}
	for _, value := range rec {
		spaceDomains = append(spaceDomains, value.toDomain())
	}

	fmt.Println(spaceDomains)

	return spaceDomains, nil

}

func (sr *mysqlSpaceRepository) AddSpace(ctx context.Context, space *spaces.Domain) error {
	rec := fromDomain(*space)

	result := sr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}