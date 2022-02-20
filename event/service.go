package event

type Service interface {
	CreateNewCompany(input CompanyInput) (Company, error)
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

	newCompany, err := s.repository.CreateNewCompany(company)
	if err != nil {
		return newCompany, err
	}

	return newCompany, err
}
