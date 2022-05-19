package event

import "gorm.io/gorm"

type Repository interface {
	CreateEvent(event Event) (Event, error)
	UpdateEvent(event Event) (Event, error)
	DeleteEvent(ID int) (bool, error)
	FindEventByID(ID int) (Event, error)
	FindAllEvent() ([]Event, error)
	FindAllEventByUserID(userID int) ([]Event, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
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

func (r *repository) DeleteEvent(ID int) (bool, error) {

	var event Event

	err := r.db.Delete(&event, ID).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) FindEventByID(ID int) (Event, error) {

	var event Event

	err := r.db.Preload("User").Preload("Company").Where("id = ?", ID).Find(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) FindAllEvent() ([]Event, error) {

	var events []Event

	err := r.db.Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repository) FindAllEventByUserID(userID int) ([]Event, error) {

	var events []Event

	err := r.db.Where("user_id = ?", userID).Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}
