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

type AppBuilder struct {
	ID         string
	Name       string
	Owners     []string
	CompanyID  int64
	AppKey     string
	CreatedBy  string
	CreatedOn  string
	ModifiedBy string
	ModifiedOn string
}

func NewAppBuilder() *AppBuilder {
	appKey, _ := generateAppKey()
	year, month, date := time.Now().Date()

	return &AppBuilder{
		ID:         uuid.New().String(),
		Name:       "",
		Owners:     []string{},
		CompanyID:  0,
		AppKey:     appKey,
		CreatedBy:  "",
		CreatedOn:  fmt.Sprintf("%02d-%02d-%d", date, month, year),
		ModifiedBy: "",
		ModifiedOn: "",
	}
}

func (builder *AppBuilder) SetName(name string) *AppBuilder {
	builder.Name = name
	return builder
}

func (builder *AppBuilder) SetCompanyId(companyId int64) *AppBuilder {
	builder.CompanyID = companyId
	return builder
}

func (builder *AppBuilder) SetCreatedBy(createdBy string) *AppBuilder {
	builder.CreatedBy = createdBy
	builder.Owners = append(builder.Owners, createdBy)
	return builder
}

func (builder *AppBuilder) SetModifiedBy(modifiedBy string) *AppBuilder {
	builder.ModifiedBy = modifiedBy
	return builder
}

func (builder *AppBuilder) SetModifiedOn(modifiedOn string) *AppBuilder {
	builder.ModifiedOn = modifiedOn
	return builder
}

func (builder *AppBuilder) Build() *App {
	return &App{
		ID:         builder.ID,
		Name:       builder.Name,
		Owners:     builder.Owners,
		CompanyID:  builder.CompanyID,
		AppKey:     builder.AppKey,
		CreatedBy:  builder.CreatedBy,
		CreatedOn:  builder.CreatedOn,
		ModifiedBy: builder.ModifiedBy,
		ModifiedOn: builder.ModifiedOn,
	}
}

type ConfigBuilder struct {
	Name        string
	Description string
	Type        Type
	Value       any
	ServiceID   string
	CreatedBy   string
	CreatedOn   string
}

func NewConfigBuilder() *ConfigBuilder {
	year, month, date := time.Now().Date()
	return &ConfigBuilder{
		Name:        "",
		Description: "",
		Type:        STRING,
		Value:       nil,
		ServiceID:   "",
		CreatedBy:   "",
		CreatedOn:   fmt.Sprintf("%02d-%02d-%d", date, month, year),
	}
}

func (builder *ConfigBuilder) SetName(name string) *ConfigBuilder {
	builder.Name = name
	return builder
}

func (builder *ConfigBuilder) SetDescription(description string) *ConfigBuilder {
	builder.Description = description
	return builder
}

func (builder *ConfigBuilder) SetType(typeValue Type) *ConfigBuilder {
	builder.Type = typeValue
	return builder
}

func (builder *ConfigBuilder) SetValue(value any) *ConfigBuilder {
	builder.Value = value
	return builder
}

func (builder *ConfigBuilder) SetServiceID(serviceID string) *ConfigBuilder {
	builder.ServiceID = serviceID
	return builder
}

func (builder *ConfigBuilder) SetCreatedBy(createdBy string) *ConfigBuilder {
	builder.CreatedBy = createdBy
	return builder
}

func (builder *ConfigBuilder) Build() *Config {
	return &Config{
		ID:          uuid.NewString(),
		Name:        builder.Name,
		Description: builder.Description,
		Type:        builder.Type,
		Value:       builder.Value,
		ServiceID:   builder.ServiceID,
		CreatedBy:   builder.CreatedBy,
		CreatedOn:   builder.CreatedOn,
		ModifiedBy:  "",
		ModifiedOn:  "",
	}
}
