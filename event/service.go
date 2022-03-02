package event

type Service interface {
	CreateNewCompany(input CompanyInput) (Company, error)
	SaveCompanyLogo(ID int, fileLocation string) (Company, error)
	CreateNewEvent(input EventInput) (Event, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) CreateNewCompany(input CompanyInput) (Company, error) {

	company := Company{}
	company.Name = input.Name
	company.WebURL = input.WebURL
	company.ShortDescription = input.ShortDescription

	newCompany, err := s.repository.CreateCompany(company)
	if err != nil {
		return newCompany, err
	}

	return newCompany, err
}

func (s *service) SaveCompanyLogo(ID int, fileLocation string) (Company, error) {

	company, err := s.repository.FindCompanyByID(ID)
	if err != nil {
		return company, err
	}

	company.LogoURL = fileLocation
	updateCompany, err := s.repository.UpdateCompany(company)
	if err != nil {
		return company, err
	}

	return updateCompany, nil
}

func (s *service) CreateNewEvent(input EventInput) (Event, error) {

	event := Event{}
	event.Title = input.Title
	event.Description = input.Description
	event.Perks = input.Perks
	event.StartDate = input.StartDate
	event.ClosingRegistration = input.ClosingRegistration
	event.LimitedTo = input.LimitedTo
	event.CompanyID = input.CompanyID
	event.UserID = input.User.ID
	event.Status = "on going"

	// TODO: fill CategoryID variabel future

	newEvent, err := s.repository.CreateEvent(event)
	if err != nil {
		return newEvent, err
	}

	return newEvent, nil
}
