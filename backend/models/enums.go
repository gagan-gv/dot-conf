package models

type Role int
type UserStatus int
type RequestStatus int
type Type int

const (
	USER Role = iota
	ADMIN
	SUPER_ADMIN
)

func (r Role) String() string {
	return [...]string{"USER", "ADMIN", "SUPER_ADMIN"}[r]
}

const (
	ACTIVE UserStatus = iota
	INACTIVE
)

func (u UserStatus) String() string {
	return [...]string{"ACTIVATED", "DEACTIVATED"}[u]
}

const (
	APPROVED RequestStatus = iota
	OPEN
	REJECTED
)

func (r RequestStatus) String() string {
	return [...]string{"APPROVED", "OPEN", "REJECTED"}[r]
}

const (
	STRING Type = iota
	NUMBER
	OBJECT
	BOOLEAN
)

func (t Type) String() string {
	return [...]string{"STRING", "NUMBER", "OBJECT", "BOOLEAN"}[t]
}
