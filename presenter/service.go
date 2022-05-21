package presenter

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

type Service interface {
	CreateNewPresenter(input PresenterInput) (Presenter, error)
}

func (s service) CreateNewPresenter(input PresenterInput) (Presenter, error) {

	presenter := Presenter{}
	presenter.Name = input.Name
	presenter.ShortDescription = input.ShortDescription
	presenter.EventID = input.EventID

	newPresenter, err := s.repository.CreatePresenter(presenter)
	if err != nil {
		return newPresenter, err
	}

	return newPresenter, nil
}
