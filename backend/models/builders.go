package models

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type CompanyBuilder struct {
	Name         string
	Email        string
	AdminId      string
	RegisteredOn string
	ModifiedOn   string
	DocumentPath string
}

func NewCompanyBuilder() *CompanyBuilder {
	year, month, date := time.Now().Date()
	return &CompanyBuilder{
		Name:         "",
		Email:        "",
		AdminId:      "",
		RegisteredOn: fmt.Sprintf("%02d-%02d-%d", date, month, year),
		ModifiedOn:   "",
		DocumentPath: "",
	}
}

func (builder *CompanyBuilder) SetName(name string) *CompanyBuilder {
	builder.Name = name
	return builder
}

func (builder *CompanyBuilder) SetEmail(email string) *CompanyBuilder {
	builder.Email = email
	return builder
}

func (builder *CompanyBuilder) SetAdminId(adminId string) *CompanyBuilder {
	builder.AdminId = adminId
	return builder
}

func (builder *CompanyBuilder) SetModifiedOn(modifiedOn string) *CompanyBuilder {
	builder.ModifiedOn = modifiedOn
	return builder
}

func (builder *CompanyBuilder) SetDocumentPath(documentPath string) *CompanyBuilder {
	builder.DocumentPath = documentPath
	return builder
}

func (builder *CompanyBuilder) Build() *Company {
	return &Company{
		Name:         builder.Name,
		Email:        builder.Email,
		AdminId:      builder.AdminId,
		RegisteredOn: builder.RegisteredOn,
		ModifiedOn:   builder.ModifiedOn,
		DocumentPath: builder.DocumentPath,
	}
}

type UserBuilder struct {
	ID           string
	Email        string
	Name         string
	Password     string
	Role         Role
	CompanyID    int64
	Status       UserStatus
	RegisteredOn string
}

func NewUserBuilder() *UserBuilder {
	year, month, date := time.Now().Date()
	return &UserBuilder{
		ID:           uuid.New().String(),
		Email:        "",
		Name:         "",
		Password:     "",
		Role:         USER,
		CompanyID:    0,
		Status:       ACTIVE,
		RegisteredOn: fmt.Sprintf("%02d-%02d-%d", date, month, year),
	}
}

func (builder *UserBuilder) SetEmail(email string) *UserBuilder {
	builder.Email = email
	return builder
}

func (builder *UserBuilder) SetName(name string) *UserBuilder {
	builder.Name = name
	return builder
}

func (builder *UserBuilder) SetPassword(password string) *UserBuilder {
	builder.Password = password
	return builder
}

func (builder *UserBuilder) SetRole(role Role) *UserBuilder {
	builder.Role = role
	return builder
}

func (builder *UserBuilder) SetCompanyId(companyId int64) *UserBuilder {
	builder.CompanyID = companyId
	return builder
}

func (builder *UserBuilder) SetStatus(status UserStatus) *UserBuilder {
	builder.Status = status
	return builder
}

func (builder *UserBuilder) Build() *User {
	return &User{
		ID:           builder.ID,
		Email:        builder.Email,
		Name:         builder.Name,
		Password:     builder.Password,
		Role:         builder.Role,
		CompanyID:    builder.CompanyID,
		Status:       builder.Status,
		RegisteredOn: builder.RegisteredOn,
	}
}

type ServiceBuilder struct {
	ID         string
	Name       string
	Owners     []string
	CompanyID  int64
	ServiceKey string
	CreatedBy  string
	CreatedOn  string
	ModifiedBy string
	ModifiedOn string
}

func NewServiceBuilder() *ServiceBuilder {
	serviceKey, _ := generateServiceKey()
	year, month, date := time.Now().Date()

	return &ServiceBuilder{
		ID:         uuid.New().String(),
		Name:       "",
		Owners:     []string{},
		CompanyID:  0,
		ServiceKey: serviceKey,
		CreatedBy:  "",
		CreatedOn:  fmt.Sprintf("%02d-%02d-%d", date, month, year),
		ModifiedBy: "",
		ModifiedOn: "",
	}
}

func (builder *ServiceBuilder) SetName(name string) *ServiceBuilder {
	builder.Name = name
	return builder
}

func (builder *ServiceBuilder) SetCompanyId(companyId int64) *ServiceBuilder {
	builder.CompanyID = companyId
	return builder
}

func (builder *ServiceBuilder) SetCreatedBy(createdBy string) *ServiceBuilder {
	builder.CreatedBy = createdBy
	builder.Owners = append(builder.Owners, createdBy)
	return builder
}

func (builder *ServiceBuilder) SetModifiedBy(modifiedBy string) *ServiceBuilder {
	builder.ModifiedBy = modifiedBy
	return builder
}

func (builder *ServiceBuilder) SetModifiedOn(modifiedOn string) *ServiceBuilder {
	builder.ModifiedOn = modifiedOn
	return builder
}

func (builder *ServiceBuilder) Build() *Service {
	return &Service{
		ID:         builder.ID,
		Name:       builder.Name,
		Owners:     builder.Owners,
		CompanyID:  builder.CompanyID,
		ServiceKey: builder.ServiceKey,
		CreatedBy:  builder.CreatedBy,
		CreatedOn:  builder.CreatedOn,
		ModifiedBy: builder.ModifiedBy,
		ModifiedOn: builder.ModifiedOn,
	}
}
