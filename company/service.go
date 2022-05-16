package company

type Service interface {
	CreateNewCompany(input CompanyInput) (Company, error)
	SaveCompanyLogo(ID int, fileLocation string) (Company, error)
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
