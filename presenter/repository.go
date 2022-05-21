package presenter

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreatePresenter(presenter Presenter) (Presenter, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) CreatePresenter(presenter Presenter) (Presenter, error) {

	err := r.db.Create(&presenter).Error
	if err != nil {
		return presenter, err
	}

	return presenter, nil
}
