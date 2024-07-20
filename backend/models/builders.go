package models

import (
	"github.com/google/uuid"
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
	return &CompanyBuilder{
		Name:         "",
		Email:        "",
		AdminId:      "",
		RegisteredOn: "",
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

func (builder *CompanyBuilder) SetRegisteredOn(registeredOn string) *CompanyBuilder {
	builder.RegisteredOn = registeredOn
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
	return &UserBuilder{
		ID:           uuid.New().String(),
		Email:        "",
		Name:         "",
		Password:     "",
		Role:         USER,
		CompanyID:    0,
		Status:       ACTIVE,
		RegisteredOn: "",
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

func (builder *UserBuilder) SetRegisteredOn(registeredOn string) *UserBuilder {
	builder.RegisteredOn = registeredOn
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
