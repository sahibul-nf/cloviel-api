package event

type CompanyFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	WebURL           string `json:"web_url"`
	LogoURL          string `json:"logo_url"`
	ShortDescription string `json:"short_description"`
}

func FormatCompany(company Company) CompanyFormatter {

	formatter := CompanyFormatter{}
	formatter.ID = company.ID
	formatter.Name = company.Name
	formatter.WebURL = company.WebURL
	formatter.LogoURL = company.LogoURL
	formatter.ShortDescription = company.ShortDescription

	return formatter
}
