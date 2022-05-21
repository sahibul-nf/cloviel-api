package presenter

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindPresenterByID(ID int) (Presenter, error)
	CreatePresenter(presenter Presenter) (Presenter, error)
	UpdatePresenter(presenter Presenter) (Presenter, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) FindPresenterByID(ID int) (Presenter, error) {

	var presenter Presenter
	
	err := r.db.Where("id = ?", ID).Find(&presenter).Error
	if err != nil {
		return presenter, err
	}

	return presenter, nil
}

func (r repository) CreatePresenter(presenter Presenter) (Presenter, error) {

	err := r.db.Create(&presenter).Error
	if err != nil {
		return presenter, err
	}

	return presenter, nil
}

func (r repository) UpdatePresenter(presenter Presenter) (Presenter, error) {
	
	err := r.db.Save(&presenter).Error
	if err != nil {
		return presenter, err
	}

	return presenter, nil
}
