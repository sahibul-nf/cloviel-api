package event

import "gorm.io/gorm"

type Repository interface {
	CreateCompany(company Company) (Company, error)
	UpdateCompany(company Company) (Company, error)
	FindCompanyByID(ID int) (Company, error)
	CreateEvent(event Event) (Event, error)
	UpdateEvent(event Event) (Event, error)
	FindEventByID(ID int) (Event, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCompany(company Company) (Company, error) {

	err := r.db.Create(&company).Error
	if err != nil {
		return company, err
	}

	return company, err
}

func (r *repository) UpdateCompany(company Company) (Company, error) {

	err := r.db.Save(&company).Error
	if err != nil {
		return company, err
	}

	return company, nil
}

func (r *repository) FindCompanyByID(ID int) (Company, error) {
	var company Company

	err := r.db.Where("id = ?", ID).Find(&company).Error
	if err != nil {
		return company, err
	}

	return company, nil
}

func (r *repository) CreateEvent(event Event) (Event, error) {
	
	err := r.db.Create(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) UpdateEvent(event Event) (Event, error) {
	
	err := r.db.Save(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) FindEventByID(ID int) (Event, error) {
	
	var event Event

	err := r.db.Where("id = ?", ID).Find(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}