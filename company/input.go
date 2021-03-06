package company

type CompanyInput struct {
	Name             string `json:"name" binding:"required"`
	WebURL           string `json:"web_url"`
	ShortDescription string `json:"short_description"`
}

type CompanyLogoInput struct {
	CompanyID int `form:"company_id" binding:"required"`
}
