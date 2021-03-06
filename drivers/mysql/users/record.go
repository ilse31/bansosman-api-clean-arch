package users

import (
	"bansosman/bussiness/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	NIK       int    `gorm:"unique"`
	Name      string `gorm:"unique"`
	Email     string
	Password  string
	FotoRumah string
	Gaji      int
	Alamat    string
}

func ToDomain(rec Users) users.Domain {
	return users.Domain{
		ID:        int(rec.ID),
		NIK:       rec.NIK,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Gaji:      rec.Gaji,
		Alamat:    rec.Alamat,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		NIK:       rec.NIK,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Gaji:      rec.Gaji,
		Alamat:    rec.Alamat,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		NIK:      domain.NIK,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Gaji:     domain.Gaji,
		Alamat:   domain.Alamat,
	}
}
func FromDomainUpdate(domain users.Domain) Users {
	return Users{
		ID:       uint(domain.ID),
		NIK:      domain.NIK,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Gaji:     domain.Gaji,
		Alamat:   domain.Alamat,
	}
}

func ToDomainArray(user []Users) []users.Domain {
	var response []users.Domain

	for _, val := range user {
		response = append(response, ToDomain(val))
	}
	return response
}
