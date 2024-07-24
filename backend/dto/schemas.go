package dto

type RegisterCompany struct {
	CompanyName  string `json:"company_name"`
	CompanyEmail string `json:"company_email"`
	AdminEmail   string `json:"email"`
	AdminName    string `json:"admin_name"`
	Password     string `json:"password"`
}
