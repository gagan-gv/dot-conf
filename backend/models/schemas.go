package models

import (
	"github.com/lib/pq"
)

type Company struct {
	ID           int64  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	AdminId      string `json:"admin_id" gorm:"unique"`
	RegisteredOn string `json:"registered_on"`
	ModifiedOn   string `json:"modified_on,omitempty"`
	DocumentPath string `json:"document_path"`
}

type User struct {
	ID           string     `json:"id" gorm:"primaryKey"`
	Email        string     `json:"email" gorm:"unique"`
	Name         string     `json:"name"`
	Password     string     `json:"password,-"`
	Role         Role       `json:"role"`
	CompanyID    int64      `json:"company_id"`
	Status       UserStatus `json:"status"`
	RegisteredOn string     `json:"registered_on"`
}

type App struct {
	ID         string         `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	Owners     pq.StringArray `json:"owners" gorm:"type:text[]"`
	CompanyID  int64          `json:"company_id"`
	AppKey     string         `json:"app_key" gorm:"unique"`
	CreatedBy  string         `json:"created_by"`
	CreatedOn  string         `json:"created_on"`
	ModifiedBy string         `json:"modified_by,omitempty"`
	ModifiedOn string         `json:"modified_on,omitempty"`
}

type Config struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        Type   `json:"type"`
	Value       any    `json:"value" gorm:"type:text"`
	ServiceID   string `json:"service_id"`
	CreatedBy   string `json:"created_by"`
	CreatedOn   string `json:"created_on"`
	ModifiedBy  string `json:"modified_by,omitempty"`
	ModifiedOn  string `json:"modified_on,omitempty"`
}
