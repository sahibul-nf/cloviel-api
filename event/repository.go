package event

import "gorm.io/gorm"

type Repository interface {
	CreateNewCompany(company Company) (Company, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateNewCompany(company Company) (Company, error) {

	err := r.db.Create(&company).Error
	if err != nil {
		return company, err
	}

	return company, err
}
