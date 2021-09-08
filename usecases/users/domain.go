package users

import "time"

type Domain struct {
	ID        int
	Name      string
	Password  string
	Device    []Device
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Device struct {
	ID        int
	Name      string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Role struct {
	ID   int
	Name string
}

type Usecase interface {
	// Create user Domain
	CreateUser(user Domain) (int, error)
	// Add Device when user register the device
	AddDevice(device Device, userID int) error
}

type Repository interface {
	UpsertUser(user Domain) (Domain, error)
	InsertDevice(device Device, userID int) error
}
