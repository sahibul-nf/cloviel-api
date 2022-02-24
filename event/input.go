package event

type CompanyInput struct {
	Name             string `json:"name" binding:"required"`
	WebURL           string `json:"web_url"`
	ShortDescription string `json:"short_description"`
}

type SaveCompanyLogoInput struct {
	CompanyID int `form:"company_id" binding:"required"`
}
