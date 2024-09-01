package dto

type RegisterCompany struct {
	CompanyName  string `json:"company_name"`
	CompanyEmail string `json:"company_email"`
	AdminEmail   string `json:"admin_email"`
	AdminName    string `json:"admin_name"`
	Password     string `json:"admin_password,-"`
}

type UserDetails struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password,-"`
}

type AppRegistrationDetails struct {
	Name        string   `json:"name"`
	OwnerEmails []string `json:"owner_emails"`
}
