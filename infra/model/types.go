package model

type Role int32

const (
	RoleUndefined Role = 0
	RoleAdmin     Role = 1
	RoleUser      Role = 2
)
